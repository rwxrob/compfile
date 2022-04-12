// Copyright 2022 Robert S. Muhlestein.
// SPDX-License-Identifier: Apache-2.0

/*
Package compfile is a completion driver for Bonzai command trees and
fulfills the bonzai.Completer. See Complete method for details.
*/
package compfile

import (
	"path/filepath"
	"strings"

	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/fn/filt"
	"github.com/rwxrob/fs"
	"github.com/rwxrob/fs/dir"
)

// New returns a pointer to a struct that fulfills bonzai.Completer.
// This can be called from within Z.Cmd assignment:
//
//     var Cmd = &Z.Cmd{
//       Name: `some`,
//       Comp: compfile.New(),
//     }
//
func New() *comp { return new(comp) }

// comp fulfills the bonzai.Completer interface.
type comp struct{}

// Complete returns all file names for the directory and file prefix passed.
// If nothing is passed assumes the current working directory.  This
// completer is roughly based on the behavior of the bash shell with
// forward slashes as separators and escaped spaces. By using this
// completer (instead of the shell) the command line interface remains
// consistent across all runtimes. Note that unlike bash completion no
// indication of the type of file is provided.
func (comp) Complete(x bonzai.Command, args ...string) []string {

	if len(args) > 1 {
		return []string{}
	}

	if args == nil || (len(args) > 0 && args[0] == "") {
		return dir.EntriesWithSlash(".")
	}

	// catch edge cases
	if len(args) == 0 {
		if x != nil {
			return []string{x.GetName()} // will add tailing space
		}
		return dir.EntriesWithSlash("")
	}

	first := strings.TrimRight(args[0], string(filepath.Separator))
	d, pre := filepath.Split(first)

	if d == "" {
		list := filt.HasPrefix(dir.Entries("."), pre)
		if len(list) == 1 && fs.IsDir(list[0]) {
			return dir.EntriesWithSlash(list[0])
		}
		return dir.AddSlash(list)
	}

	for {
		list := filt.BaseHasPrefix(dir.Entries(d), pre)
		if len(list) > 1 {
			return dir.AddSlash(list)
		}
		if fs.IsDir(list[0]) {
			d = list[0]
			continue
		}
		return dir.AddSlash(list)
	}

	return []string{}
}
