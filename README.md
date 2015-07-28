# which

**which** is a Go implementation of the UNIX command of the same name.

## Usage

The executable works exactly like the `which` command:

    which [-as] program ...

## Install

    go get github.com/bfontaine/which

## Library

**which** is also usable as a Go library:

```go
// import "github.com/bfontaine/which/which"

// get the first executable in $PATH
executable := which.One("vim")

// get all executables in $PATH
executables := which.All("vim")
```
