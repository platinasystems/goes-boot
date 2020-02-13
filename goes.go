// Copyright © 2015-2020 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// +build !bootrom

package main

import (
	"github.com/platinasystems/goes"
	"github.com/platinasystems/goes/cmd"
	"github.com/platinasystems/goes/cmd/bang"
	"github.com/platinasystems/goes/cmd/buildid"
	"github.com/platinasystems/goes/cmd/buildinfo"
	"github.com/platinasystems/goes/cmd/cat"
	"github.com/platinasystems/goes/cmd/cd"
	"github.com/platinasystems/goes/cmd/chmod"
	"github.com/platinasystems/goes/cmd/cli"
	"github.com/platinasystems/goes/cmd/cmdline"
	"github.com/platinasystems/goes/cmd/cp"
	"github.com/platinasystems/goes/cmd/daemons"
	"github.com/platinasystems/goes/cmd/dhcpcd"
	"github.com/platinasystems/goes/cmd/dmesg"
	"github.com/platinasystems/goes/cmd/echo"
	"github.com/platinasystems/goes/cmd/elsecmd"
	"github.com/platinasystems/goes/cmd/env"
	"github.com/platinasystems/goes/cmd/exec"
	"github.com/platinasystems/goes/cmd/exit"
	"github.com/platinasystems/goes/cmd/export"
	"github.com/platinasystems/goes/cmd/falsecmd"
	"github.com/platinasystems/goes/cmd/femtocom"
	"github.com/platinasystems/goes/cmd/ficmd"
	"github.com/platinasystems/goes/cmd/function"
	"github.com/platinasystems/goes/cmd/grub"
	"github.com/platinasystems/goes/cmd/grubd"
	"github.com/platinasystems/goes/cmd/ifcmd"
	"github.com/platinasystems/goes/cmd/iminfo"
	"github.com/platinasystems/goes/cmd/insmod"
	"github.com/platinasystems/goes/cmd/iocmd"
	"github.com/platinasystems/goes/cmd/ip"
	"github.com/platinasystems/goes/cmd/kexec"
	"github.com/platinasystems/goes/cmd/keys"
	"github.com/platinasystems/goes/cmd/kill"
	"github.com/platinasystems/goes/cmd/ln"
	"github.com/platinasystems/goes/cmd/log"
	"github.com/platinasystems/goes/cmd/ls"
	"github.com/platinasystems/goes/cmd/lsmod"
	"github.com/platinasystems/goes/cmd/lsof"
	"github.com/platinasystems/goes/cmd/mkdir"
	"github.com/platinasystems/goes/cmd/mknod"
	"github.com/platinasystems/goes/cmd/mount"
	"github.com/platinasystems/goes/cmd/mountd"
	"github.com/platinasystems/goes/cmd/ping"
	"github.com/platinasystems/goes/cmd/platina/mk1/bootc"
	"github.com/platinasystems/goes/cmd/ps"
	"github.com/platinasystems/goes/cmd/pwd"
	"github.com/platinasystems/goes/cmd/reboot"
	"github.com/platinasystems/goes/cmd/reload"
	"github.com/platinasystems/goes/cmd/restart"
	"github.com/platinasystems/goes/cmd/rm"
	"github.com/platinasystems/goes/cmd/rmmod"
	"github.com/platinasystems/goes/cmd/slashinit"
	"github.com/platinasystems/goes/cmd/sleep"
	"github.com/platinasystems/goes/cmd/source"
	"github.com/platinasystems/goes/cmd/start"
	"github.com/platinasystems/goes/cmd/stop"
	"github.com/platinasystems/goes/cmd/stty"
	"github.com/platinasystems/goes/cmd/sync"
	"github.com/platinasystems/goes/cmd/testcmd"
	"github.com/platinasystems/goes/cmd/thencmd"
	"github.com/platinasystems/goes/cmd/truecmd"
	"github.com/platinasystems/goes/cmd/umount"
	"github.com/platinasystems/goes/cmd/version"
	"github.com/platinasystems/goes/cmd/wget"
	"github.com/platinasystems/goes/lang"
)

var Goes = &goes.Goes{
	NAME:    name,
	APROPOS: Apropos,
	ByName: map[string]cmd.Cmd{
		"!":        bang.Command{},
		"cli":      &cli.Command{},
		"bootc":    &bootc.Command{},
		"cat":      cat.Command{},
		"cd":       &cd.Command{},
		"chmod":    chmod.Command{},
		"cp":       cp.Command{},
		"daemons":  daemons.Admin,
		"dhcpcd":   &dhcpcd.Command{},
		"dmesg":    dmesg.Command{},
		"echo":     echo.Command{},
		"else":     &elsecmd.Command{},
		"env":      &env.Command{},
		"exec":     exec.Command{},
		"exit":     exit.Command{},
		"export":   export.Command{},
		"false":    falsecmd.Command{},
		"femtocom": femtocom.Command{},
		"fi":       &ficmd.Command{},
		"function": &function.Command{},
		"goes-daemons": &daemons.Server{
			Init: [][]string{
				[]string{"dhcpcd"},
				[]string{"mountd"},
				[]string{"grubd"},
			},
		},
		"grub":   &grub.Command{},
		"grubd":  &grubd.Command{},
		"if":     &ifcmd.Command{},
		"insmod": insmod.Command{},
		"io":     iocmd.Command{},
		"ip":     ip.Goes,
		"kexec":  &kexec.Command{},
		"keys":   keys.Command{},
		"kill":   kill.Command{},
		"ln":     ln.Command{},
		"log":    log.Command{},
		"ls":     ls.Command{},
		"lsmod":  lsmod.Command{},
		"lsof":   lsof.Command{},
		"mkdir":  mkdir.Command{},
		"mknod":  mknod.Command{},
		"mount":  mount.Command{},
		"mountd": mountd.Command(make(chan struct{})),
		"ping":   ping.Command{},
		"ps":     ps.Command{},
		"pwd":    pwd.Command{},
		"reboot": &reboot.Command{},

		"reload":  reload.Command{},
		"restart": &restart.Command{},
		"rm":      rm.Command{},
		"rmmod":   rmmod.Command{},

		"show": &goes.Goes{
			NAME:  "show",
			USAGE: "show OBJECT",
			APROPOS: lang.Alt{
				lang.EnUS: "print stuff",
			},
			ByName: map[string]cmd.Cmd{
				"buildid":   buildid.Command{},
				"buildinfo": buildinfo.Command{},
				"cmdline":   cmdline.Command{},
				"iminfo":    iminfo.Command{},
				"machine":   Machine,
				"version":   version.Command{},
			},
		},

		"/init":  &slashinit.Command{Hook: disableBootdog},
		"sleep":  sleep.Command{},
		"source": &source.Command{},
		"start":  &start.Command{Gettys: consoles},
		"stop":   &stop.Command{},
		"stty":   stty.Command{},
		"sync":   sync.Command{},
		"[":      testcmd.Command{},
		"then":   &thencmd.Command{},
		"true":   truecmd.Command{},
		"umount": umount.Command{},
		"wget":   wget.Command{},
	},
}
