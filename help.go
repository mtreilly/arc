// Copyright (c) 2025 Arc Engineering
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"strings"
)

func showHelp() {
	fmt.Println("arc - Agent-aware development workflows")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  arc <command> [args...]")
	fmt.Println("  arc <url>                  Auto-detect URL type and dispatch")
	fmt.Println()
	fmt.Println("Core Commands:")
	fmt.Println("  commit    Create commits with AI-generated messages")
	fmt.Println("  repo      Manage external repositories")
	fmt.Println("  arxiv     Fetch and manage arXiv papers")
	fmt.Println("  sessions  Manage agent sessions")
	fmt.Println("  tmux      Tmux integration")
	fmt.Println("  ai        AI-powered tools")
	fmt.Println("  ask       Ask AI questions")
	fmt.Println()
	fmt.Println("Management Commands:")
	fmt.Println("  config    View and edit configuration")
	fmt.Println("  db        Database utilities")
	fmt.Println("  env       Environment snapshots")
	fmt.Println("  prompt    Manage prompt templates")
	fmt.Println("  workflow  Workflow orchestration")
	fmt.Println()
	fmt.Println("Integration Commands:")
	fmt.Println("  discord   Discord integration")
	fmt.Println("  plugin    Plugin system")
	fmt.Println()

	// Show installed subcommands
	cmds := listSubcommands()
	if len(cmds) > 0 {
		fmt.Println("Installed Subcommands:")
		fmt.Printf("  %s\n", strings.Join(cmds, ", "))
		fmt.Println()
	}

	fmt.Println("URL Shortcuts:")
	fmt.Println("  arc https://github.com/owner/repo    → arc repo fetch <url>")
	fmt.Println("  arc https://arxiv.org/abs/2301.00001 → arc arxiv fetch <url>")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -h, --help      Show this help message")
	fmt.Println("  -v, --version   Show version information")
	fmt.Println()
	fmt.Println("Run 'arc <command> --help' for more information on a command.")
}
