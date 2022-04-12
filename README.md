# 🌳 Go Bonzai™ Completer Template

*Create a new GitHub project using this template and change this
README.md to match your project. Make all your template changes before
making your first commit.*

[![GoDoc](https://godoc.org/github.com/rwxrob/compfoo?status.svg)](https://godoc.org/github.com/rwxrob/compfoo)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

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

## Reminders

* Change `compfoo` every place to your project name (`git grep compfoo`)
* Remove anything you don't need
* Change `.github/FUNDING.yaml` to your own information
* Update `.gitignore` to your liking
* Will need to `go get -u` to update dependencies
