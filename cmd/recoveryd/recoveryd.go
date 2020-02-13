// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// Package recoveryd implements the client side of the goes-recovery
// service.
package recoveryd

import (
	"time"

	"github.com/platinasystems/goes"
	"github.com/platinasystems/goes/cmd"
	"github.com/platinasystems/goes/lang"
)

type Command struct {
	g    *goes.Goes
	done chan struct{}
}

const pollInterval = time.Second * 120

func (*Command) String() string { return "recoveryd" }

func (*Command) Usage() string { return "recoveryd" }

func (*Command) Apropos() lang.Alt {
	return lang.Alt{
		lang.EnUS: "goes-recovery client daemon",
	}
}

func (c *Command) Close() error {
	close(c.done)
	return nil
}

func (c *Command) Goes(g *goes.Goes) { c.g = g }

func (*Command) Kind() cmd.Kind { return cmd.Daemon }

func (c *Command) Main(args ...string) error {
	c.done = make(chan struct{})

	for {
		if !func() bool {
			t := time.NewTicker(pollInterval)
			defer t.Stop()

			select {
			case <-c.done:
				return false
			case <-t.C:
				return true
			}
		}() {
			return nil
		}
	}

	return nil
}
