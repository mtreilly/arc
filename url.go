// Copyright (c) 2025 Arc Engineering
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"os"
	"strings"
)

// isURL checks if a string looks like a URL.
func isURL(s string) bool {
	return strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://")
}

// dispatchURL handles URL arguments by dispatching to the appropriate subcommand.
func dispatchURL(url string, args []string) {
	switch {
	case strings.Contains(url, "arxiv.org"):
		dispatchSubcommand("arxiv", append([]string{"fetch", url}, args...))
	case strings.Contains(url, "github.com"):
		dispatchSubcommand("repo", append([]string{"fetch", url}, args...))
	case strings.Contains(url, "gitlab.com"):
		dispatchSubcommand("repo", append([]string{"fetch", url}, args...))
	case strings.Contains(url, "bitbucket.org"):
		dispatchSubcommand("repo", append([]string{"fetch", url}, args...))
	default:
		fmt.Fprintf(os.Stderr, "arc: unknown URL type: %s\n", url)
		os.Exit(1)
	}
}
