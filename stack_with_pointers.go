package main

import (
	"fmt"
)

// Stack with pointers
// Sharing down typically stays on the stack

func main2() {
	n := 4
	inc(&n)
	fmt.Println(n)
}

//go:noinline
func inc(x *int) {
	*x++
}
