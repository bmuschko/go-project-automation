package main

import "fmt"

func main() {
	// [go] ./main.go:6:2: age declared and not used
	age := 14
	// [go-vet] ./main.go:7: Printf format %s reads arg #1, but call has 0 args
	fmt.Printf("Ben's age is %s")
}
