package main

import (
	"fmt"

	"github.com/platinasystems/goes/lang"
)

type machine string

const Machine machine = "goes-boot"

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
