// Copyright © 2020 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// Package recoveryd implements the client side of the goes-recovery
// service.
package recoveryd

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/platinasystems/goes"
	"github.com/platinasystems/goes/cmd"
	"github.com/platinasystems/goes/lang"

	"github.com/platinasystems/url"

	"github.com/cavaliercoder/go-cpio"
	"github.com/jpillora/backoff"
	"github.com/ulikunitz/xz"
	//"golang.org/x/sys/unix"
)

type Command struct {
	g   *goes.Goes
	Url string
}

func (*Command) String() string { return "recoveryd" }

func (*Command) Usage() string { return "recoveryd" }

func (*Command) Apropos() lang.Alt {
	return lang.Alt{
		lang.EnUS: "goes-recovery client daemon",
	}
}

func (c *Command) Goes(g *goes.Goes) { c.g = g }

func (*Command) Kind() cmd.Kind { return cmd.Daemon }

func (c *Command) Main(args ...string) error {
	b := &backoff.Backoff{
		Min:    10 * time.Second,
		Max:    240 * time.Second,
		Factor: 2,
		Jitter: false,
	}
	for {
		err := c.tryArchive()
		if err == nil {
			fmt.Printf("Success!\n")
			err := syscall.Kill(1, syscall.SIGTERM)
			if err != nil {
				fmt.Printf("Error sending SIGTERM to init: %s\n",
					err)
			}
			select {
			case <-goes.Stop:
				return nil
			}
		} else {
			fmt.Fprintf(os.Stderr, "Error in tryArchive: %s\n",
				err)
		}
		for {
			if !func() bool {
				t := time.NewTicker(b.Duration())
				defer t.Stop()

				select {
				case <-goes.Stop:
					return false
				case <-t.C:
					return true
				}
			}() {
				return nil
			}
			break
		}
	}

	return nil
}

func (c *Command) tryArchive() (err error) {
	txz, err := url.Open(c.Url)
	if err != nil {
		return fmt.Errorf("Error opening %s: %w", c.Url, err)
	}
	defer txz.Close()

	tr, err := xz.NewReader(txz)
	if err != nil {
		return fmt.Errorf("Error in xz.NewReader: %w", err)
	}

	archive := cpio.NewReader(tr)
	target := "/var/run/goes/recoveryd-" + strconv.Itoa(os.Getppid())
	err = os.RemoveAll(target)
	if err != nil {
		return fmt.Errorf("Unable to remove %s: %w", target, err)
	}
	err = os.MkdirAll(target, 0755)
	if err != nil {
		return fmt.Errorf("Unable to MkdirAll %s: %w", target, err)
	}
	err = syscall.Mount("", target, "tmpfs", 0, "")
	if err != nil {
		return fmt.Errorf("Unable to create tmpfs: %w", err)
	}
	defer func() {
		syscall.Unmount(target, 0)
	}()
	for {
		file, err := archive.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return fmt.Errorf("Error reading archive: %v", err)
		}
		path := filepath.Join(target, file.Name)
		if (file.FileInfo().Mode() & os.ModeSymlink) != 0 {
			err := os.Symlink(file.Linkname, path)
			if err != nil {
				fmt.Printf("Error creating symlink %s->%s:%s\n",
					path, file.Linkname, err)
			}
			continue
		}
		if file.FileInfo().IsDir() {
			err := os.MkdirAll(path, os.FileMode(file.Mode))
			if err != nil {
				fmt.Printf("Error creating directory %s: %s\n",
					path, err)
			}
			continue
		}
		//if (file.FileInfo().Mode() & os.ModeDevice) != 0 {
		//			err := syscall.Mknod(path, uint32(file.FileInfo().Mode()),
		//		int(unix.Mkdev(uint32(file.Devmajor), uint32(file.Devminor))))
		//	if err != nil {
		//		fmt.Printf("Error creating special file %s: %s",
		//			path, err)
		//	}
		//	continue
		//}
		if !file.FileInfo().Mode().IsRegular() {
			fmt.Printf("Can't create %x file %s\n",
				file.FileInfo().Mode(), path)
			continue
		}
		t, err := os.OpenFile(
			path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(file.Mode))
		if err != nil {
			return err
		}
		defer t.Close()
		if _, err := io.Copy(t, archive); err != nil {
			return err
		}

	}
	fmt.Printf("Target directory is %s\n", target)
	err = copyRecurse(target, "/", true)
	if err != nil {
		fmt.Printf("copyRecurse error %s\n", err)
		return err
	}
	return nil
}

// copyRecurse copies the tree of files at src to dst
func copyRecurse(src, dst string, overwrite bool) (err error) {
	stat, err := os.Lstat(src)
	if err != nil {
		return
	}
	if stat.Mode().IsDir() {
		files, err := ioutil.ReadDir(src)
		if err != nil {
			return err
		}
		if _, err := os.Stat(dst); os.IsNotExist(err) {
			err = os.Mkdir(dst, stat.Mode())
			if err != nil {
				return err
			}
			fmt.Printf("mkdir -m %o %s\n",
				stat.Mode()&os.ModePerm, dst)
		}
		for _, file := range files {
			base := path.Base(file.Name())
			err = copyRecurse(filepath.Join(src, base),
				filepath.Join(dst, base), overwrite)
			if err != nil {
				return err
			}
		}
		return nil
	}
	if (stat.Mode() & os.ModeSymlink) != 0 {
		link, err := os.Readlink(src)
		if err == nil {
			os.Remove(dst)
			err := os.Symlink(link, dst)
			if err != nil {
				fmt.Printf("Error creating link %s: %s\n",
					dst, err)
			}
		} else {
			fmt.Printf("Error reading link %s: %s\n",
				src, err)
		}
		return nil
	}
	if !stat.Mode().IsRegular() {
		fmt.Printf("Skipping non-regular file %s\n", src)
		return nil // Skip special files
	}
	fmt.Printf("copying %s=>%s mode %o\n", src, dst, stat.Mode())
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	if _, err := os.Stat(dst); (overwrite && err == nil) ||
		os.IsNotExist(err) {

		os.Remove(dst) // In case its a symlink
		destination, err := os.OpenFile(
			dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			stat.Mode())
		if err != nil {
			return err
		}
		defer destination.Close()

		_, err = io.Copy(destination, source)
		if err != nil {
			return err
		}
		fmt.Printf("copy %s=>%s\n", src, dst)
	}
	return nil
}
