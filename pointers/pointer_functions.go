package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	a := 4
	squareVal(a)
	squareValMod(&a)
	fmt.Printf("main-->%p\n", initPerson())
}

func initPerson() *person { // we are goint to return address pointer of person type
	m := person{name: "hitesh", age: 29}
	fmt.Printf("initPerson--> %p\n", &m)
	return &m // in this case the address is copied into the heap not the  execution stack since the stack will be empty once the function ends
}

func squareVal(v int) { // Value of a is copied and then returned
	v *= v
	fmt.Println(&v, v)
}

func squareValMod(v *int) { // In this case * is not a dereferencing operator, it means an address value of type pointer will be sent
	*v *= *v
	fmt.Println(v, *v)
}
