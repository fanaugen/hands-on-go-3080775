// functions/begin/main.go
package main

import (
	"errors"
	"fmt"
	"strconv"
)

// simple greet function
func greet() string {
	return "Hello"
}

// greetWithName returns a greeting with the name
func greetWithName(name string) string {
	return "Hello, " + name + "!"
}

// greetWithNameAndAge returns a greeting with the name and age of the person
// NOTE: `greeting` is a named return parameter, and line 23 is a so called naked return
func greetWithNameAndAge(name string, age int) (greeting string) {
	greeting = "Good evening, " + name + "! I’m happy to report that you are " + strconv.Itoa(age) + " years old"
	return
}

// divide divides two numbers and returns the result
// if the second number is zero, it returns  error
func divide(dividend, divisor int) (quotient int, err error) {
	if divisor == 0 {
		return 0, errors.New("Cannot divide by zero!")
	}
	return dividend / divisor, nil
}

func main() {
	// invoke the functions defined above
	fmt.Println(greet())
  fmt.Println(greetWithName("Alyx"))
  fmt.Println(greetWithNameAndAge("Robert", 45))

	// invoke divide function
	fmt.Println(divide(156, 12))

	// invoke divide function with zero denominator to get an error
	fmt.Println(divide(9000, 0))
}
