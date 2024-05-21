package main

import "fmt"

/*
Create a function in Go that takes two int arguments and swaps their values using pointers.
*/
func main() {
	a := 10
	b := 20
	fmt.Printf("a = %d, b = %d\n", a, b)

	swap(&a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}
