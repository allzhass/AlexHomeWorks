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
		name := getName(&i)
		age := getAge(&i)

		students[name] = age
	}

	fmt.Println("Students:")
	for name, age := range students {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}

func getName(i *int) string {
	var name string
	fmt.Printf("[%d] Enter student name: ", *i+1)
	fmt.Scan(&name)
	if (len(name) < 2) || (len(name) > 256) {
		fmt.Println("Name must be string and its' length must be at least 2 chars and less than 256 chars")
		return getName(i)
	}
	return name
}

func getAge(i *int) int {
	var age int
	fmt.Printf("[%d] Enter student age: ", *i+1)
	fmt.Scan(&age)
	if (age <= 16) || (age >= 90) {
		fmt.Println("Age must be integer which more than 16 and less than 90")
		return getAge(i)
	}
	return age
}
