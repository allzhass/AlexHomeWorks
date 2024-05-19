package main

import "fmt"

/*
1. Write a Go program that creates a pointer to an integer variable
using the new() function, assigns the value 10 to this variable,
and then prints the variable's value and the address to which the pointer points.
*/
func main() {
	intPointer := new(int)
	*intPointer = 10

	fmt.Printf("Value: %v\n", *intPointer)
	fmt.Printf("Pointer: %v\n", intPointer)
}
