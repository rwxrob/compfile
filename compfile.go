// Copyright 2022 Robert S. Muhlestein.
// SPDX-License-Identifier: Apache-2.0

/*
Package compfile is a completion driver for Bonzai command trees and
fulfills the bonzai.Completer. See Complete method for details.
*/
package compfile

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/fn/filt"
)

// New returns a pointer to a struct that fulfills bonzai.Completer.
// This can be called from within Z.Cmd assignment:
//
//	var Cmd = &Z.Cmd{
//	  Name: `some`,
//	  Comp: compfile.New(),
//	}
func New() *comp { return new(comp) }

// comp fulfills the bonzai.Completer interface.
type comp struct{}

// Complete returns all file names for the directory and file prefix
// passed. If nothing is passed assumes the current working directory.
// This completer is roughly based on the behavior of the bash shell
// with forward slashes as separators and escaped spaces. By using this
// completer (instead of the shell) the command line interface remains
// consistent across all runtimes. Note that unlike bash completion no
// indication of the type of file is provided.
func (comp) Complete(_ bonzai.Command, args ...string) []string {
	if len(args) == 0 || args[len(args)-1] == "" {
		return entriesWithSlash(".")
	}

	arg := args[len(args)-1]
	arg = expandHome(arg)

	path, sub := filepath.Split(arg)
	if path == "" {
		path = "."
	}

	return recurse(path, sub)
}

var seperator string

func init() {
	// "/" on unix, "\" on windows
	seperator = string(filepath.Separator)
}

func entriesWithSlash(dir string) (entries []string) {
	items, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	for _, i := range items {
		name := filepath.Join(dir, i.Name())
		if i.IsDir() {
			name += seperator
		}

		entries = append(entries, name)
	}

	return
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func expandHome(path string) string {
	if !strings.HasPrefix(path, "~") {
		return path
	}

	path = strings.TrimPrefix(path, "~")
	home, err := os.UserHomeDir()
	if err != nil {
		// panicking or logging errors in a completion driver will ruin
		// UX
		return path
	}

	return home + path
}

func recurse(path, file string) []string {
	entries := entriesWithSlash(path)

	// complete relative paths
	if file == "." || file == ".." {
		sep := seperator
		entries = append(entries, file+sep)
	}

	list := filt.BaseHasPrefix(entries, file)
	if len(list) == 1 && isDir(list[0]) {
		return recurse(list[0], "")
	}

	return list
}
