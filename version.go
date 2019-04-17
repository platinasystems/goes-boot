package main

import (
	"fmt"

	"github.com/platinasystems/goes/lang"
)

type version string

// Version format :: v<MAJOR>.<MINOR>.<PATCH>[-rc<CANDIDATE>]
var Version version = "v1.3.0-test"

func (*version) String() string { return "version" }
func (*version) Usage() string  { return "[show ]version" }

func (*version) Apropos() lang.Alt {
	return lang.Alt{
		lang.EnUS: "print machine version",
	}
}

func (p *version) Main(...string) error {
	fmt.Println(string(*p))
	return nil
}
