package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Enter your name and age (ex. Max 33), please: ")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("Hello, %s! You are %d years old.", name, age)
}
