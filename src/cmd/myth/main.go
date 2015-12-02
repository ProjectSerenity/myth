// Copyright 2015 Myth Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

type Command interface {
	Short() string
	Long() string
	Run(args []string) error
}

var commands = make(map[string]Command)

func Register(name string, command Command) {
	if _, ok := commands[name]; ok {
		panic("command " + name + " has already been registered")
	}

	commands[name] = command
}

func usage() {
	fmt.Fprintln(os.Stderr, `Myth is a tool for managing Myth source code.

Usage:

	myth command [arguments]

The commands are:
`)

	names := make([]string, 0, len(commands))
	for name := range commands {
		names = append(names, name)
	}

	sort.Strings(names)

	w := tabwriter.NewWriter(os.Stderr, 0, 8, 0, '\t', 0)
	for _, name := range names {
		cmd := commands[name]
		fmt.Fprintf(w, "\xff\t\xff%s\t%s\n", name, cmd.Short())
	}
	w.Flush()

	os.Exit(2)
}

func main() {
	if len(os.Args) == 1 {
		usage()
	}
}
