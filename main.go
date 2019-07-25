// garbaged - clean up manually modified hosts, quick
// main.go: CLI initialization
//
// Copyright 2018 Threat Stack, Inc.
// Licensed under the BSD 3-clause license; see LICENSE.md for more information.
// Author: Patrick T. Cable II <pat.cable@threatstack.com>

package main

import "os"

func main() {
	os.Exit(Run(os.Args[1:]))
}
