package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str string

	fmt.Println("Write a word, please:")
	fmt.Scan(&str)

	runes := []rune(str)
	fmt.Printf("String length: %d. First character by rune: %c\n", len(str), runes[0])

	firstRune, _ := utf8.DecodeRuneInString(str)
	fmt.Printf("String length: %d. First character by utf8: %c\n", len(str), firstRune)
}
