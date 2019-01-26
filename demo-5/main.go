package main

import "fmt"

// [gocyclo] 4 main main main.go:5:1
func main() {
	x := 1
	y := 2
	z := 2

	if x == 1 {
		if x > y {
			y = x
		} else {
			if z > y {
				z = x
			}
		}
	} else {
		y = z
	}

	fmt.Printf("x: %d, y: %d, z: %d", x, y, z)
}
