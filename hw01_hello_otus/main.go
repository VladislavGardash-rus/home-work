package main

import (
	"fmt" //nolint:gofumpt,gci
	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Print(reverse.String("Hello, OTUS!"))
}
