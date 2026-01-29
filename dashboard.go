// Copyright (c) 2025 Arc Engineering
// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/yourorg/arc-sdk/config"
	"github.com/yourorg/arc-sdk/db"
	"github.com/yourorg/arc-sdk/store"
	"github.com/yourorg/arc-sdk/utils"
)

func runDashboard() {
	cfg, err := config.Load()
	if err != nil {
		// Config not found - show welcome
		showWelcome()
		return
	}
	_ = cfg // Used for future expansion

	database, err := db.Open(db.DefaultDBPath())
	if err != nil {
		showWelcome()
		return
	}
	defer database.Close()

	ctx := context.Background()

	fmt.Println("arc - Agent-aware development workflows")
	fmt.Println()

	// Quick actions
	fmt.Println("Quick Actions:")
	fmt.Println("  arc repo fetch <url>     Add a repository")
	fmt.Println("  arc arxiv fetch <url>    Fetch an arXiv paper")
	fmt.Println("  arc sessions find        Find agent sessions")
	fmt.Println()

	// Recent sessions
	sessStore := store.NewSessionsStore(database)
	sessions, err := sessStore.FindRecent(ctx, 3)
	if err == nil && len(sessions) > 0 {
		fmt.Println("Recent Sessions:")
		for _, s := range sessions {
			id := s.ID
			if len(id) > 8 {
				id = id[:8]
			}
			timeAgo := utils.HumanizeTime(s.ModTS)
			fmt.Printf("  %s  %-10s  %s  (%s)\n", id, s.Agent, truncatePath(s.CWD, 40), timeAgo)
		}
		fmt.Println()
	}

	// Repository stats
	repoStore := store.NewReposStore(database)
	totalCount, _ := repoStore.Count(ctx)
	clonedCount, _ := repoStore.CountCloned(ctx)
	if totalCount > 0 {
		fmt.Printf("Repositories: %d indexed", totalCount)
		if clonedCount > 0 {
			fmt.Printf(" (%d cloned)", clonedCount)
		}
		fmt.Println()
		fmt.Println()
	}

	// Available subcommands
	cmds := listSubcommands()
	if len(cmds) > 0 {
		fmt.Println("Available commands:")
		fmt.Printf("  %s\n", strings.Join(cmds, ", "))
		fmt.Println()
	}

	fmt.Println("Run 'arc help' for more commands.")
}

func showWelcome() {
	fmt.Println("arc - Agent-aware development workflows")
	fmt.Println()
	fmt.Println("Get started:")
	fmt.Println("  arc repo fetch https://github.com/user/repo")
	fmt.Println("  arc arxiv fetch https://arxiv.org/abs/2301.00001")
	fmt.Println()
	fmt.Println("Run 'arc help' for more commands.")
}

func truncatePath(path string, maxLen int) string {
	if len(path) <= maxLen {
		return path
	}
	// Try to show the end of the path
	return "..." + path[len(path)-maxLen+3:]
}
