// types/structs/pointers/begin/main.go
package main

import "fmt"

type author struct {
	first, last string
}

// fullName returns the full name of the author
func (a author) fullName() string {
	return a.first + " " + a.last
}

// method to update the author’s name
// if the receiver type in the first pair of brackets is given without an asterisk,
// golang will call the method on a copy of the struct
// To modify the actual struct in memory, we need to use a pointer receiver
func (a *author) updateName(first, last string) {
	a.first = first
	a.last = last
}

func main() {
	a := author{
		first: "Mar1",
		last:  "Twain",
	}

	fmt.Println(a.fullName())

	// update the author’s name
	a.updateName("Mark", "Twain")
	fmt.Println(a.fullName())
}
