// types/structs/methods/begin/main.go
package main

import "fmt"

type author struct {
	first, last string
}

// specify a receiver in round parentheses after the `func` keyword to
// attach the function as a method to the struct, to be called in dot notation
func (a author) fullName() string {
	return a.first + " " + a.last
}

func main() {
	a := author{
		first: "Marcus",
		last:  "Aurelius",
	}

	// call the `fullName` method defined above
	fmt.Println(a.fullName())
}
