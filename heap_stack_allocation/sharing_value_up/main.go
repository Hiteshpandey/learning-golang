package main

import "fmt"

// Pushed to stack frame 1
// function main and n in stack
func main() {

	// answer function is called and is placed in stack frame 2
	n := answer()
	// Back to stack frame 1
	// Reuse and override existing stack frame 2 to push print and n value
	// The value is n is accessible here even after overriding the stack frame 2
	// dereference value of n and divide by 2
	// perform the print operation
	fmt.Println(*n / 2)
}

// Stack frame 2
// Now *int as before means the value would be a reference varaible. a pointer
// But this time the pointer reference is sent upward
// stack fame2 is sending return value as a reference to it's varaible but to stack frame 1
// This is in revese from before

/*
This is a big problem since if the reference value is stored in stack frame 2 of the function answer
This means that after returning reference of x back to the main function the stack frame 2 will become invalid
This makes the frame eligible to be overwritten by next operation since invalid stacks are overriden and reused by golang
But as we can clearly see we are dereferencing the value of x by using *n and doing operations on it
That operation now will be invalid operation in theory but that doesn't happen since the value of x is not stored in the stack frame 2

To prevent this case
In case of bottom to top stack passing of reference varaible
or we can say passing backward of reference varaible by current fucntion to the previous calling function

Golang saves the refere varaible into a HEAP MEMORY
now when the reference is sent, it is basically a reference to the varaible value that is stored in HEAP memory
Rather than the stack memory frame 2
*/
func answer() *int {
	// Assign a value to x
	x := 42
	// Return the reference of variable x of which value is in a heap memory
	return &x
}
