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
	
	// Passing value of n by reference
	square(&n)
}

// Stack frame 2
// At execution time stored in stack
// This will be in Stack Frame 2
func square(x int) int {
	// we have variable x = 4 in stack frame 2
	// return computed value
	return x * x
}

// (*) int means that the value passed to the function inc would be passed by reference, i.e a pointer
// This also means that at varaible x, we would have a memory loaction pointer rather than a value
// If we print x directly, we would get a hexadecimal vaue that will be a reference to the location of the value , in our case n
// This makes it possible to update the passed value that is passed from the main function residing the main function execution stack memory frame
// The value of x currently not scoped and it points to the stack frame 1/ main funciton instance of variable "n"
// So... currenlty the pointer value in x is pointing to the stack memory for this operation i.e variable "n"
// Concluding : if we change x then n should reflect the same change, since n and x are pointing to same meory location, as x is pointing to the memory address of value n
func inc(x *int) {
	// (*) is dereferencing the value of pointer x\
	// which is getting the value from memory location pointer by pointer x
	// and increment the value at pointer location by one
	*x++
}
