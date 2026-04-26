package main

import (
	"os"
	"os/exec"
	"strings"
)

// RunScript executes a shell script string and returns combined output.
func RunScript(script string) (string, error) {
	cmd := exec.Command("sh", "-c", script)
	out, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(out)), err
}

// ReadFileOrEmpty reads a file; returns empty string on error.
func ReadFileOrEmpty(path string) string {
	b, _ := os.ReadFile(path)
	return string(b)
}

// ExpandHome replaces leading ~/ with the user's home directory.
func ExpandHome(p string) string {
	if len(p) >= 2 && p[:2] == "~/" {
		home, _ := os.UserHomeDir()
		return home + "/" + p[2:]
	}
	return p
}
