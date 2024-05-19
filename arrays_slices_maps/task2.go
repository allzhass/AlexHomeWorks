package main

import "fmt"

/*
2. Create a function in Go that takes a slice of integers as input
and returns a new slice containing only the even numbers from the original slice.
*/
func main() {
	arr := make([]int, 10)
	arr[0] = 44
	arr[1] = 2
	arr[2] = 4
	arr[3] = 5
	arr[4] = 55
	arr[5] = 66
	arr[6] = 78
	arr[7] = 87
	arr[8] = 90
	arr[9] = 100

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
