// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// First stage bootstrap. Looks for second stage or recovers the system
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func main() {
	args := os.Args
	if filepath.Base(args[0]) == "goes-recovery" {
		args[0] = "goes-boot"
	}
	err := syscall.Exec("../goes-boot/goes-boot", args, os.Environ())
	fmt.Printf("syscall.Exec failed: %s\n", err)
}
