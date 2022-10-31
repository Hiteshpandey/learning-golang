package main

import "fmt"

// Stack frame 1
func main() {
	// assign value to n
	// stored in stack memory n=4
	n := 4
	// Stack frame 2 is started for increment function
	// Passing value of n by reference
	// This will now create stack frame 2
	// this will update n with new value as 5
	inc(&n)
	// back to stack frame 1
	// Reuse existing stack frame 2 to push function println
	// push n to stack frame 2 and print value of n
	fmt.Println(n)

	/*Note:-
	Point to be noted
	Passing pointer down usually stays in the stack only
	if we pass a reference to the function to be executed next, the variable stays in the stack
	and will reference the value in the previous function state
	in our case x in inc function points to n=4 in main function
	*/

}

// (*) int means that the value passed to the function inc would be passed by reference, i.e a pointer

// This also means that at varaible x, we would have a memory loaction pointer rather than a value

// If we print x directly, we would get a hexadecimal vaue that will be a reference to the location of the value , in our case n

// This makes it possible to update the passed value that is passed from the main function residing in the main function execution stack memory frame

// The value of x currently not scoped and it points to the stack frame 1/ main funciton instance of variable "n"

// Concluding : if we change x then n should reflect the same change, since n and x are pointing to same meory location, as x is pointing to the memory address of value n

func inc(x *int) {
	// (*) is dereferencing the value of pointer x\
	// which means getting the value from memory location by pointer x
	// and increment the value at pointer location by one
	*x++
}
