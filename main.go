// Copyright (c) 2025 Arc Engineering
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"os"
)

var (
	Version   = "dev"
	Commit    = "none"
	BuildDate = "unknown"
)

func main() {
	if len(os.Args) < 2 {
		runDashboard()
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "-h", "--help", "help":
		showHelp()
	case "-v", "--version", "version":
		showVersion()
	default:
		if isURL(cmd) {
			dispatchURL(cmd, os.Args[2:])
			return
		}
		dispatchSubcommand(cmd, os.Args[2:])
	}
}

func showVersion() {
	fmt.Printf("arc %s (commit: %s, built: %s)\n", Version, Commit, BuildDate)
}
