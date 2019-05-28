// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// +build amd64

package main

import "github.com/platinasystems/goes/cmd/iocmd"

var consoles = []string{"/dev/ttyS0"}
var io = iocmd.Command{}
