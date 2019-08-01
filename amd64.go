// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// +build amd64

package main

import (
	"github.com/platinasystems/goes/cmd/start"
)

var consoles = []start.TtyCon{
	{Tty: "/dev/ttyS0", Baud: 115200},
	{Tty: "/dev/ttyS1", Baud: 57600},
}
