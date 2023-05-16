package main

import (
	"fmt"
)

// Head with pointers
// Sharing up typically escapes to the heap

func main() {
	n := answer()
	fmt.Println(*n / 2)
}

//go:noinline
func answer() *int {
	// the compiler knows that x will be used outside of the stack frame, the compiler assigned it to the heap at compile time
	x := 42
	return &x
}
