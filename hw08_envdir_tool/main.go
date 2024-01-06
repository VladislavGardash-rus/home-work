package main

import (
	"fmt"
	"os"
)

const errorReturnCode = 111 //nolint:all
const successReturnCode = 0 //nolint:all

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "go-envdir error: not enough arguments")
		os.Exit(errorReturnCode)
	}

	environments, err := ReadDir(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "go-envdir error: %w", err)
		os.Exit(errorReturnCode)
	}

	code := RunCmd(args[1:], environments)
	if code == errorReturnCode {
		fmt.Fprintln(os.Stderr, "go-envdir error: not execute %w", args[1])
		os.Exit(errorReturnCode)
	}

	os.Exit(code)
}
