package main

import "fmt"

func main() {
	i, j := 42, 2701
	fmt.Println(&i, &j)

	p := &i
	fmt.Println(*p) // * before the variable while using is called deferencing .
	// This means that instead of returing an address of the pointer it will give you the value
	*p = 21
	// This will change the value of i
	fmt.Println(i)

	p = &j
	*p = *p / 37 // This means we are changing value of j
}
