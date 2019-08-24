// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

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

func main() {
	args := os.Args
	fmt.Printf("in goes-recovery: args %v\n", args)
	if filepath.Base(args[0]) == "goes-recovery" {
		args[0] = "goes-boot"
	}

	err := syscall.Mount("none", "/dev", "devtmpfs", 0, "")
	if err != nil {
		fmt.Printf("syscall.Mount(/dev) failed: %s\n", err)
		os.Exit(1)
	}

	err = syscall.Mount("/dev/sda1", "/boot", "ext4", syscall.MS_RDONLY,
		"")
	if err != nil {
		fmt.Printf("syscall.Mount(/boot) failed: %s\n", err)
		os.Exit(1)
	}

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
			os.Exit(1)
		}
	}

	err = syscall.Unmount("/boot", 0)
	if err != nil {
		fmt.Printf("syscall.Unmount(/boot) failed: %s\n", err)
		os.Exit(1)
	}

	err = syscall.Unmount("/dev", 0)
	if err != nil {
		fmt.Printf("syscall.Unmount(/dev) failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("leaving goes-recovery: args %v\n", args)
	err = syscall.Exec("/sbin/goes-boot", args, os.Environ())
	fmt.Printf("syscall.Exec failed: %s\n", err)
}
