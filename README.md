# 🌳 Go Bonzai™ File Completer

[![GoDoc](https://godoc.org/github.com/rwxrob/compfile?status.svg)](https://godoc.org/github.com/rwxrob/compfile)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

Provides the regular shell completion most expect for files and
directories with the trailing slash added to directories for quick
visualization. Nothing special beyond that (colors, emojis, etc.)

## Style Guidelines

* Everything through `go fmt` or equiv, no exceptions
* In Vim `set textwidth=72` (not 80 to line numbers fit)
* Use `/* */` for package documentation comment, `//` elsewhere
* Smallest possible names for given scope while still clear
* Favor additional packages (possibly in `internal`) over long names
* Package globals that will be used a lot can be single capital
* Must be good reason to use more than 4 character pkg name
* Avoid unnecessary comments
* Use "deciduous tree" emoji 🌳 to mark Bonzai stuff

## Legal

Copyright 2022 Robert S. Muhlestein (<mailto:rob@rwx.gg>)  
SPDX-License-Identifier: Apache-2.0

"Bonzai" and "bonzai" are legal trademarks of Robert S. Muhlestein but
can be used freely to refer to the Bonzai™ project
<https://github.com/rwxrob/bonzai> without limitation. To avoid
potential developer confusion, intentionally using these trademarks to
refer to other projects --- free or proprietary --- is prohibited.
