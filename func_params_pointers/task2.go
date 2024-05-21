package main

import "fmt"

/*
Write a program that creates a function to calculate the factorial of a number in Go.
The function should take an integer as input and return its factorial using recursion.
*/
func main() {
	a := 4
	fmt.Printf("%d! = %d", a, factorial(a))
}

func factorial(a int) int {
	if a > 1 {
		return a * factorial(a-1)
	} else {
		return a
	}
}
