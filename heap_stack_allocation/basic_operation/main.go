package main

import "fmt"

// Stack Frame 1
// step 1. main func is pushed to stack for execution
func main() {
	// assign value to n
	// stored in stack memory n=4
	n := 4
	// store in stack memory
	// Stack frame 2 is started for square
	// 4 is returned from stack 2 square func -> stack 2
	// n2 = 16 is stored in stack frame 1
	n2 := square(n)
	/*Note:-
	Here stack frame 2 is not deleted and still exists,
	Currently it is still considered invalid to use.
	this will be used in next step*/

	// For println function a new stack frame is needed
	// But since we aready have Stack Frame 2
	// It will override that stack frame 2
	// Reusing it to store values for println function execution

	// Stack frame 2 reused and flushed
	// println and n2 are pushed to stack frame 2
	fmt.Println(n2)

}

// Stack frame 2
// At execution time stored in stack
// This will be in Stack Frame 2
func square(x int) int {
	// we have variable x = 4 in stack frame 2
	// return computed value
	return x * x
}
