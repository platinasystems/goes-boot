// Copyright © 2020 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package main

import (
	"fmt"

	"github.com/platinasystems/goes/lang"
)

type machine string

func (machine) String() string { return "machine" }
func (machine) Usage() string  { return "show machine" }

func (machine) Apropos() lang.Alt {
	return lang.Alt{
		lang.EnUS: "print machine name",
	}
}

func (s machine) Main(...string) error {
	fmt.Println(string(s))
	return nil
}
