// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// +build arm

package main

import (
	"github.com/platinasystems/goes/cmd/start"
)

const consoles = []start.TtyCon{
	{Tty: "/dev/ttymxc0", Baud: 115200},
}
