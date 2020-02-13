// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// +build bootrom

package main

import (
	"fmt"

	"github.com/platinasystems/goes"
	"github.com/platinasystems/goes-recovery/cmd/recoveryd"
	"github.com/platinasystems/goes/cmd"
	"github.com/platinasystems/goes/cmd/bang"
	"github.com/platinasystems/goes/cmd/cli"
	"github.com/platinasystems/goes/cmd/cp"
	"github.com/platinasystems/goes/cmd/daemons"
	"github.com/platinasystems/goes/cmd/dhcpcd"
	"github.com/platinasystems/goes/cmd/ip"
	"github.com/platinasystems/goes/cmd/kexec"
	"github.com/platinasystems/goes/cmd/ls"
	"github.com/platinasystems/goes/cmd/mount"
	"github.com/platinasystems/goes/cmd/reboot"
	"github.com/platinasystems/goes/cmd/slashinit"
	"github.com/platinasystems/goes/cmd/source"
	"github.com/platinasystems/goes/cmd/start"
	"github.com/platinasystems/goes/cmd/stop"
	"github.com/platinasystems/goes/cmd/umount"
	"github.com/platinasystems/goes/cmd/wget"
	"github.com/platinasystems/goes/lang"

	"github.com/platinasystems/ioport"
)

var Goes = &goes.Goes{
	NAME: name,
	APROPOS: lang.Alt{
		lang.EnUS: "the coreboot goes machine",
	},
	ByName: map[string]cmd.Cmd{
		"!":       bang.Command{},
		"cli":     &cli.Command{},
		"cp":      &cp.Command{},
		"daemons": daemons.Admin,
		"dhcpcd":  &dhcpcd.Command{},
		"goes-daemons": &daemons.Server{
			Init: [][]string{
				[]string{"dhcpcd"},
				[]string{"recoveryd"},
			},
		},
		"ip":        ip.Goes,
		"kexec":     &kexec.Command{},
		"ls":        ls.Command{},
		"mount":     mount.Command{},
		"reboot":    &reboot.Command{},
		"recoveryd": &recoveryd.Command{},
		"/init":     &slashinit.Command{Hook: disableBootdog},
		"source":    &source.Command{},
		"start":     &start.Command{Gettys: consoles},
		"stop":      &stop.Command{},
		"umount":    umount.Command{},
		"wget":      wget.Command{},
	},
}

func disableBootdog() (err error) {
	b, err := ioport.Inb(0x604)
	if err != nil {
		return fmt.Errorf("Error in Inb(0x604): %s", err)
	}
	b = b & 0xfd
	err = ioport.Outb(0x604, b)
	if err != nil {
		return fmt.Errorf("Error in Outb(0x604, %x): %s", b, err)
	}
	return
}
