// Package home provides utilities for dealing with the user's home directory.
package home

import (
	"log/slog"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

var homedir, homedirErr = os.UserHomeDir()

func init() {
	if homedirErr != nil {
		slog.Error("Failed to get user home directory", "error", homedirErr)
	}
}

// Dir returns the user home directory.
func Dir() string {
	return homedir
}

// Short replaces the actual home path from [Dir] with `~`.
func Short(p string) string {
	if homedir == "" || !strings.HasPrefix(p, homedir) {
		return p
	}
	return filepath.Join("~", strings.TrimPrefix(p, homedir))
}

// Long expands the leading `~` (including `~user`) to a user's home directory
// and replaces environment variables like $HOME.
func Long(p string) string {
	if p == "" {
		return p
	}

	p = expandTilde(p)
	return os.ExpandEnv(p)
}

func expandTilde(p string) string {
	if !strings.HasPrefix(p, "~") {
		return p
	}

	afterTilde := p[1:]
	sep := strings.IndexAny(afterTilde, "/\\")
	var username, remainder string
	if sep == -1 {
		username = afterTilde
		remainder = ""
	} else {
		username = afterTilde[:sep]
		remainder = afterTilde[sep+1:]
	}

	homePath := userHome(username)
	if homePath == "" {
		return p
	}

	if remainder == "" {
		return homePath
	}

	return filepath.Join(homePath, remainder)
}

func userHome(username string) string {
	if username == "" {
		return homedir
	}

	u, err := user.Lookup(username)
	if err != nil || u.HomeDir == "" {
		return ""
	}

	return u.HomeDir
}
