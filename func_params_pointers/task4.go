package main

import (
	"fmt"
	"math/big"
)

/*
Write Fibonacchi function that would work for big number and don't have stackoverflow
*/
func main() {
	// не запускай, тормозит
	//fib := new(int64)
	//*fib = 100
	//fmt.Printf("Fibbonachi %d equals to %d", *fib, fibbonachi(fib))

	fib1 := new(uint64)
	*fib1 = 500
	fmt.Printf("Fibbonachi2 %d equals to %d\n", *fib1, fibbonachi2(fib1))

	fib2 := new(uint64)
	*fib2 = 500
	fmt.Printf("Fibbonachi3 %d equals to %v\n", *fib2, fibbonachi3(fib2))
}

func fibbonachi3(num *uint64) *big.Int {
	res0 := big.NewInt(0)
	res1 := big.NewInt(1)
	res := big.NewInt(1)

	if *num == 1 {
		res.Set(res0)
	} else if *num == 2 {
		res.Set(res1)
	} else {
		var i uint64
		for i = 3; i <= *num; i++ {
			res = res.Add(res0, res1)
			res0.Set(res1)
			res1.Set(res)
		}
	}
	return res
}

func fibbonachi2(num *uint64) uint64 {
	res0 := new(uint64)
	res1 := new(uint64)
	res := new(uint64)
	*res0 = 0
	*res1 = 1
	if *num == 1 {
		res = res0
	} else if *num == 2 {
		res = res1
	} else {
		var i uint64
		for i = 3; i <= *num; i++ {
			*res = *res0 + *res1
			*res0 = *res1
			*res1 = *res
		}
	}
	return *res
}

func fibbonachi(fib *int64) int64 {
	if *fib == 1 {
		return 0
	} else if *fib == 2 {
		return 1
	} else {
		lastFib1 := new(int64)
		*lastFib1 = *fib - 1

		lastFib2 := new(int64)
		*lastFib2 = *fib - 2

		return fibbonachi(lastFib1) + fibbonachi(lastFib2)
	}
}
