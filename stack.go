package main

import (
	"fmt"
)

// Stack
func main1() {
	n := 4
	n2 := square(n)
	fmt.Println(n2)
}

//go:noinline
func square(x int) int {
	return x * x
}
