package main

import "fmt"

/*
Implement a function in Go that takes an array of integers
and a pointer to a function for transforming each element of the array.
The function should modify each element of the array according to the logic defined by the passed function.
*/
func main() {
	arr := []int{1, 1, 1, 1}

	intFunc := func(x int) int {
		return x + 1
	}

	fmt.Printf("arr: %v\n", arr)
	trans(arr, intFunc)
	fmt.Printf("arr: %v\n", arr)
}

func trans(arr []int, f func(int) int) {
	for i := range arr {
		arr[i] = f(arr[i])
	}
}
