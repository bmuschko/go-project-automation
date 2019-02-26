package main

import "fmt"

func main() {
	fmt.Println(foo(1))
}

func foo(bar int) int {
	if bar > 0 {
		return 123
	}
	return 456
}
