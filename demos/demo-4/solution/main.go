package main

import (
	"os"
)

func main() {
	f, err := os.Open("foo.txt")
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte("Hello, world."))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
}
