package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	num = 70.4
	fmt.Printf("Round of %v is %v", num, int(math.Round(num)))
}
