package main

import "fmt"

/*
3. Write a program that uses a map to store information about students (name and age).
The program should prompt the user for the number of students,
then their names and ages, save this information in a map, and print it out.
*/
func main() {
	students := make(map[string]int)

	var studentNumber int
	fmt.Printf("Number of students: ")
	fmt.Scan(&studentNumber)

	for i := 0; i < studentNumber; i++ {
		var name string
		var age int

		fmt.Printf("[%d] Enter student name:", i+1)
		fmt.Scan(&name)

		fmt.Printf("[%d] Enter student age: ", i+1)
		fmt.Scan(&age)

		students[name] = age
	}

	fmt.Println("Students:")
	for name, age := range students {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}
