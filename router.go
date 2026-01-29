// Copyright (c) 2025 Arc Engineering
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"syscall"
)

// dispatchSubcommand finds and executes an arc-* binary.
func dispatchSubcommand(name string, args []string) {
	binary := findSubcommand(name)
	if binary == "" {
		fmt.Fprintf(os.Stderr, "arc: '%s' is not an arc command. See 'arc help'.\n", name)
		os.Exit(1)
	}

	// Replace current process with subcommand
	argv := append([]string{filepath.Base(binary)}, args...)
	env := os.Environ()

	if err := syscall.Exec(binary, argv, env); err != nil {
		fmt.Fprintf(os.Stderr, "arc: failed to exec %s: %v\n", name, err)
		os.Exit(1)
	}
}

// findSubcommand searches for an arc-* binary.
func findSubcommand(name string) string {
	binaryName := "arc-" + name

	// 1. Same directory as arc binary
	if self, err := os.Executable(); err == nil {
		selfDir := filepath.Dir(self)
		candidate := filepath.Join(selfDir, binaryName)
		if isExecutable(candidate) {
			return candidate
		}
	}

	// 2. PATH lookup
	if path, err := exec.LookPath(binaryName); err == nil {
		return path
	}

	return ""
}

// isExecutable checks if a file exists and is executable.
func isExecutable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Mode()&0111 != 0
}

// listSubcommands discovers available arc-* binaries.
func listSubcommands() []string {
	var cmds []string
	seen := make(map[string]bool)

	// Check same directory as arc
	if self, err := os.Executable(); err == nil {
		selfDir := filepath.Dir(self)
		matches, _ := filepath.Glob(filepath.Join(selfDir, "arc-*"))
		for _, m := range matches {
			name := filepath.Base(m)[4:] // strip "arc-"
			if !seen[name] && isExecutable(m) {
				cmds = append(cmds, name)
				seen[name] = true
			}
		}
	}

	// Check PATH
	pathDirs := filepath.SplitList(os.Getenv("PATH"))
	for _, dir := range pathDirs {
		matches, _ := filepath.Glob(filepath.Join(dir, "arc-*"))
		for _, m := range matches {
			name := filepath.Base(m)[4:]
			if !seen[name] && isExecutable(m) {
				cmds = append(cmds, name)
				seen[name] = true
			}
		}
	}

	sort.Strings(cmds)
	return cmds
}
