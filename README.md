# which

**which** is a Go implementation of the UNIX command of the same name.

Its main (and only) advantage over the standard command is its `-l` flag that makes it resolve symbolic links. This can
be especially useful when using [Homebrew](https://brew.sh/).

## Usage

The executable works exactly like the `which` command:

    which [-as] program ...

In addition, it supports a `-l` command which makes `which` resolve symbolic
links before printing the paths.

However, note that this implementation doesn’t support combined flags
(e.g. `-al` won’t work, you’ll have to use `-a -l`).

You’ll have to ensure that `$GOPATH/bin` is at the beginning of your `PATH`
environnment variable if you want to use this implementation instead of the
original one.

## Install

    go install github.com/bfontaine/which@1.0.1

## Example

```
$ which vim
/usr/local/bin/vim

$ which -l vim
/usr/local/Cellar/vim/7.4.712_1/bin/vim

$ which -a vim
/usr/local/bin/vim
/usr/bin/vim
```

## Library

**which** is also usable as a Go library:

```go
package main

import "github.com/bfontaine/which/which"

// get the first executable in $PATH
executable := which.One("vim")

// get all executables in $PATH
executables := which.All("vim")
```

## Why?

I know it doesn’t really make sense to re-write a simple tool like `which`, but
I needed the `-l` option so I wrote this. I use [Homebrew](http://brew.sh/) on
macOS and it installs binaries in a directory then symlinks them into
`/usr/local/bin/`, which means it’s not possible to get the original path by
using the original `which` command alone.
