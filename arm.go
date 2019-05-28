// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// +build arm

package main

import "github.com/platinasystems/goes/lang"

type iocmd struct{}

func (iocmd) String() string { return "io" }
func (iocmd) Usage() string  { return "io" }

func (iocmd) Apropos() lang.Alt {
	return lang.Alt{
		lang.EnUS: "not available for arm",
	}
}

func (iocmd) Main(...string) error { return nil }

var io iocmd
var consoles = []string{"/dev/ttymxc0"}
