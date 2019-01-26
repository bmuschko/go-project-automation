package main

import (
	"os"
)

func main() {
	f, _ := os.Open("foo.txt")
	// [errcheck] main.go:9:9:	f.Write([]byte("Hello, world."))
	f.Write([]byte("Hello, world."))
	// [errcheck] main.go:10:9:	f.Close()
	f.Close()
}
