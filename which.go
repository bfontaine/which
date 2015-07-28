package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bfontaine/which/which"
)

var (
	listAll bool
	silent  bool
)

func usage() {
	fmt.Fprint(os.Stderr, "usage: which [-as] program ...\n")
	os.Exit(1)
}

func printPath(path string) {
	if !silent && path != "" {
		fmt.Println(path)
	}
}

func main() {
	var failure bool

	flag.BoolVar(&listAll, "a", false, "List all instances of executables found")
	flag.BoolVar(&silent, "s", false, "No output, return 0 if any executable is found")

	flag.Parse()

	pathenv := os.Getenv("PATH")
	if pathenv == "" {
		return
	}

	if flag.NArg() == 0 {
		usage()
	}

	for _, cmd := range flag.Args() {
		if listAll {
			paths := which.All(cmd)
			for _, path := range paths {
				printPath(path)
			}
			if len(paths) == 0 {
				failure = true
			}
		} else {
			path := which.One(cmd)
			printPath(path)
			if path == "" {
				failure = true
			}
		}
	}

	if failure {
		os.Exit(1)
	}
}
