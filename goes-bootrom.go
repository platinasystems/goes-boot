// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// +build bootrom

package main

import (
	"github.com/platinasystems/goes"
	"github.com/platinasystems/goes-boot/cmd/recoveryd"
	"github.com/platinasystems/goes/cmd"
	"github.com/platinasystems/goes/cmd/bang"
	"github.com/platinasystems/goes/cmd/cli"
	"github.com/platinasystems/goes/cmd/cp"
	"github.com/platinasystems/goes/cmd/daemons"
	"github.com/platinasystems/goes/cmd/dhcpcd"
	"github.com/platinasystems/goes/cmd/echo"
	"github.com/platinasystems/goes/cmd/elsecmd"
	"github.com/platinasystems/goes/cmd/falsecmd"
	"github.com/platinasystems/goes/cmd/ficmd"
	"github.com/platinasystems/goes/cmd/function"
	"github.com/platinasystems/goes/cmd/grub"
	"github.com/platinasystems/goes/cmd/grubd"
	"github.com/platinasystems/goes/cmd/ifcmd"
	"github.com/platinasystems/goes/cmd/ip"
	"github.com/platinasystems/goes/cmd/kexec"
	"github.com/platinasystems/goes/cmd/ls"
	"github.com/platinasystems/goes/cmd/mkdir"
	"github.com/platinasystems/goes/cmd/mount"
	"github.com/platinasystems/goes/cmd/mountd"
	//	"github.com/platinasystems/goes/cmd/ping"
	"github.com/platinasystems/goes/cmd/ps"
	"github.com/platinasystems/goes/cmd/reboot"
	"github.com/platinasystems/goes/cmd/slashinit"
	"github.com/platinasystems/goes/cmd/source"
	"github.com/platinasystems/goes/cmd/start"
	"github.com/platinasystems/goes/cmd/stop"
	"github.com/platinasystems/goes/cmd/stty"
	"github.com/platinasystems/goes/cmd/testcmd"
	"github.com/platinasystems/goes/cmd/thencmd"
	"github.com/platinasystems/goes/cmd/truecmd"
	"github.com/platinasystems/goes/cmd/umount"
	"github.com/platinasystems/goes/cmd/wget"
	"github.com/platinasystems/goes/lang"
)

const Machine machine = "goes-bootrom"

var Apropos = lang.Alt{
	lang.EnUS: "the goes-bootrom goes machine",
}

var Goes = &goes.Goes{
	NAME:    name,
	APROPOS: Apropos,
	ByName: map[string]cmd.Cmd{
		"!":        bang.Command{},
		"cli":      &cli.Command{},
		"cp":       &cp.Command{},
		"daemons":  daemons.Admin,
		"echo":     echo.Command{},
		"else":     &elsecmd.Command{},
		"dhcpcd":   &dhcpcd.Command{},
		"false":    falsecmd.Command{},
		"fi":       &ficmd.Command{},
		"function": &function.Command{},
		"goes-daemons": &daemons.Server{
			Init: [][]string{
				[]string{"dhcpcd"},
				[]string{"mountd"},
				[]string{"grubd"},
				[]string{"recoveryd"},
			},
		},
		"grub":   &grub.Command{},
		"grubd":  &grubd.Command{},
		"if":     &ifcmd.Command{},
		"ip":     ip.Goes,
		"kexec":  &kexec.Command{},
		"ls":     ls.Command{},
		"mkdir":  mkdir.Command{},
		"mount":  mount.Command{},
		"mountd": &mountd.Command{},
		//"ping":      ping.Command{},
		"ps":        ps.Command{},
		"reboot":    &reboot.Command{},
		"recoveryd": &recoveryd.Command{Url: recoveryUrl},
		"/init":     &slashinit.Command{Hook: disableBootdog},
		"source":    &source.Command{},
		"start":     &start.Command{Gettys: consoles},
		"stop":      &stop.Command{},
		"stty":      stty.Command{},
		"[":         testcmd.Command{},
		"then":      &thencmd.Command{},
		"true":      truecmd.Command{},
		"umount":    umount.Command{},
		"wget":      wget.Command{},
	},
}
