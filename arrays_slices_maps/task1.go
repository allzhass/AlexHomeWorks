package main

import "fmt"

/*
1. Write a Go program that creates an array of integers with a length of 5,
fills it with numbers from 1 to 5, and then prints this array in reverse order.
*/
func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Printf("array[%d]: %d\n", i, arr[i])
	}
}
