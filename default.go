// Copyright © 2020 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// +build !bootrom

package main

import (
	"github.com/platinasystems/goes/lang"
)

const Machine machine = "goes-boot"

var Apropos = lang.Alt{
	lang.EnUS: "the goes-boot goes machine",
}
