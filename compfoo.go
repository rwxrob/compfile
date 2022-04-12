/*

Package compfoo is a completion driver for Bonzai command trees and
fulfills the bonzai.Completer package interface by returning "foo" and
"bar". It is conventional to begin Completer packages with "comp" and
always included a New function that returns a pointer to private struct
that fulfills the bonzai.Completer interface. This is because multiple completions for different branch commands might use the same completion package but require different state. The Z.Comp, however, is the default used for any command that does not assign its own Cmd.Completer.

*/
package compfoo

import "github.com/rwxrob/bonzai"

// New returns a pointer to a struct that fulfills bonzai.Completer.
// This can be called from within Z.Cmd assignment:
//
//     var Cmd = &Z.Cmd{
//       Name: `some`,
//       Comp: compfoo.New(),
//     }
//
func New() *comp { return new(comp) }

// comp fulfills the bonzai.Completer interface.
type comp struct{}

// Complete completes everything with the word "foo" and "bar"
func (comp) Complete(x bonzai.Command, args ...string) []string {
	return []string{"foo", "bar"}
}
