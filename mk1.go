// Copyright © 2020 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// +build mk1

package main

import (
	"github.com/platinasystems/goes/external/log"
	"github.com/platinasystems/ioport"
)

const packageName = "goes-boot-platina-mk1"
const recoveryUrl = "https://platina.io/goes/goes-boot-platina-mk1.cpio.xz"

func disableBootdog() (err error) {
	c1, err := ioport.Inb(0x604)
	if err != nil {
		log.Printf("err", "Error reading CPLD Ctrl-1: %s", err)
		c1 = 0 // Assume register read as zero
	}

	s1, err := ioport.Inb(0x602)
	if err != nil {
		log.Printf("err", "Error reading CPLD Status-1: %s", err)
		s1 = 0 // Assume register read as zero
	}

	qspiBoot := 0
	qspiSel := 0
	if s1&0x80 != 0 {
		qspiBoot = 1
	}
	if c1&0x80 != 0 {
		qspiSel = 1
	}
	if qspiBoot == qspiSel {
		log.Printf("alert", "Booted from QSPI%d", qspiBoot)
	} else {
		log.Printf("alert", "Booted from QSPI%d (QSPI%d was selected)",
			qspiBoot, qspiSel)
	}

	if c1&0x3 != 0 {
		en := ""
		if c1&0x1 != 0 {
			en = "WDT_RCVRY_EN"
		}
		if c1&0x2 != 0 {
			if en != "" {
				en = en + " and "
			}
			en = en + "WDT_DOG_EN"
		}
		log.Printf("alert", "CPLD watchdog status: %s", en)

		c1 &= 0x7c
		if qspiBoot != 0 {
			c1 |= 0x80
		}
		err = ioport.Outb(0x604, c1)
		if err != nil {
			log.Printf("alert",
				"Error disabling WDT in CPLD Ctrl-1: %s",
				err)
		}
	}
	return
}
