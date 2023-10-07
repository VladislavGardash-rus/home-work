package main

import (
	"golang.org/x/example/hello/reverse"
)

func main() {
	str := "Hello, OTUS!"
	revstring := reverse.String(str)
	print(revstring)
}
