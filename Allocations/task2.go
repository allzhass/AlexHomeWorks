package main

import "fmt"

/*
2. Create a function in Go that takes the size of an array as an argument,
allocates memory for this array using new(), fills it with numbers from 1 to the size of the array,
and returns a pointer to the array.
*/
func main() {
	arrSize := 15
	arr := makeArray(arrSize)
	fmt.Printf("Pointer to arr: %v\n\n", &arr)

	fmt.Printf("Value: %v", *arr)
}

func makeArray(size int) *[]int {
	arr := make([]int, size)

	for i := 0; i < size; i++ {
		arr[i] = i
	}
	return &arr
}
