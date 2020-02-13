// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// +build bootrom

// First stage bootstrap. Looks for second stage or recovers the system
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
)

func copyFile(src, dst string, perm os.FileMode) (err error) {
	b, err := ioutil.ReadFile(src)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("Error reading %s: %s", src, err)
	}
	err = ioutil.WriteFile(dst, b, perm)
	if err != nil {
		return fmt.Errorf("Error writing %s: %s", dst, err)
	}
	return nil
}

const name = "goes-bootrom"

func setupGoesBoot(dev string) (err error) {
	err = syscall.Mount("none", "/dev", "devtmpfs", 0, "")
	if err != nil {
		fmt.Printf("syscall.Mount(/dev) failed: %s\n", err)
		return
	}

	defer func() {
		err = syscall.Unmount("/dev", 0)
		if err != nil {
			fmt.Printf("syscall.Unmount(/dev) failed: %s\n", err)
		}
	}()

	err = syscall.Mount(dev, "/boot", "ext4", syscall.MS_RDONLY,
		"")
	if err != nil {
		fmt.Printf("syscall.Mount(/boot) failed: %s\n", err)
		return
	}

	defer func() {
		err = syscall.Unmount("/boot", 0)
		if err != nil {
			fmt.Printf("syscall.Unmount(/boot) failed: %s\n", err)
			return
		}
	}()

	for _, c := range []struct {
		src  string
		dst  string
		perm os.FileMode
	}{
		{"/boot/boot/goes/goes-boot", "/sbin/goes-boot", 0755},
		{"/boot/boot/goes/init", "/etc/goes/init", 0644},
		{"/boot/boot/goes/start", "/etc/goes/start", 0644},
		{"/boot/boot/goes/stop", "/etc/goes/stop", 0644},
		{"/boot/boot/goes/authorized_keys",
			"/etc/goes/sshd/authorized_keys", 0600},
		{"/boot/boot/goes/resolv.conf", "/etc/resolv.conf", 0644},
	} {
		err = copyFile(c.src, c.dst, c.perm)
		if err != nil {
			fmt.Printf("Error in copyFile: %s\n", err)
			return
		}
	}
	return nil
}

func execGoesBoot(dev string) (err error) {
	err = setupGoesBoot(dev)
	if err != nil {
		return
	}
	fmt.Printf("trying goes-bootrom (%s): args %v\n", dev, os.Args)
	err = syscall.Exec("/sbin/goes-boot", os.Args, os.Environ())
	fmt.Printf("syscall.Exec (%s) failed: %s\n", dev, err)
	return
}

func main() {
	fmt.Printf("in goes-bootrom: args %v\n", os.Args)
	if os.Args[0] == "/init" {
		_ = execGoesBoot("/dev/sdb1")
		_ = execGoesBoot("/dev/sda1")
	}

	fmt.Printf("Invoking goes - args: %v\n", os.Args)

	args := os.Args
	if filepath.Base(args[0]) == name {
		args[0] = name
	}

	if err := Goes.Main(os.Args...); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	os.Exit(1) // Will panic and reboot

}
