package main

import "fmt"

/*
3. Write a program that creates a pointer to a Person structure in Go,
allocates memory for this structure using new(), fills in the structure's fields (name and age),
and prints information about the person.
*/
func main() {
	person := new(Person)
	person.name = "Olzhas"
	person.age = 34

	fmt.Printf("Person's pointer: %v\n", &person)
	fmt.Printf("Person's name is %s \n", person.name)
	fmt.Printf("Person's age is %d \n", person.age)
}

type Person struct {
	name string
	age  int
}
