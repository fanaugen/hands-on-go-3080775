// types/constants/begin/main.go
package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// declare a constant representing pi
const pi = 3.141593 // recognized by compiler as a float

// declare a rune constant for the letter 'a'
const a rune = 'A'

func main() {
	// print the value and types of pi and a
	// In the format string, %T is the data type
	fmt.Printf("pi: %v is a %T\n", pi, pi)
	fmt.Printf("a:  %c is a %T\n", a, a)

	// use unicode package to confirm that the rune is a letter
	fmt.Println(strconv.QuoteRune(a) + " is a letter? ", unicode.IsLetter(a))
}
