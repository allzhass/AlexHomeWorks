package main

import "fmt"

/*
2. Create a function in Go that takes a slice of integers as input
and returns a new slice containing only the even numbers from the original slice.
*/
func main() {
	arr := []int{55, 1, 33, 34, 4, 5, 6, 77, 88, 100}

	fmt.Printf("Even Numbers: %v", getEvenNumberSlice(arr))
}

func getEvenNumberSlice(inArr []int) []int {
	evenArr := make([]int, 0)
	for _, v := range inArr {
		if v%2 == 0 {
			evenArr = append(evenArr, v)
		}
	}
	return evenArr
}
