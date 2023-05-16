package main

import (
	"fmt"
)

// Stack with pointers
// Sharing gown typically stays on the stack

func main2() {
	n := 4
	inc(&n)
	fmt.Println(n)
}

//go:noinline
func inc(x *int) {
	*x++
}
