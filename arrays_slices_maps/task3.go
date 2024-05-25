package main

import (
	"errors"
	"fmt"
)

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

	i := 0
	for i < studentNumber {
		name, err := getName(&i)
		if err != nil {
			fmt.Println(err)
			continue
		}
		age, err := getAge(&i)
		if err != nil {
			fmt.Println(err)
			continue
		}

		students[name] = age
		i++
	}

	fmt.Println("Students:")
	for name, age := range students {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}

func getName(i *int) (string, error) {
	var name string
	fmt.Printf("[%d] Enter student name: ", *i+1)
	fmt.Scan(&name)
	if (len(name) < 2) || (len(name) > 256) {
		return "", errors.New("Name must be string and its' length must be at least 2 chars and less than 256 chars")
	}
	return name, nil
}

func getAge(i *int) (int, error) {
	var age int
	fmt.Printf("[%d] Enter student age: ", *i+1)
	fmt.Scan(&age)
	if (age <= 16) || (age >= 90) {
		return 0, errors.New("Age must be integer which more than 16 and less than 90")
	}
	return age, nil
}
