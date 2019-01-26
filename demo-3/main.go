package main

import "fmt"

func main() {
	fmt.Println(foo(1))
}

// [golint] main.go:12:9: if block ends with a return statement, so drop this else and outdent its block
func foo(bar int) int {
	if bar > 0 {
		return 123
	} else {
		return 456
	}
}
