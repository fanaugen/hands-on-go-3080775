// types/structs/embed/begin/main.go
package main

import "fmt"

type person struct {
	first, last string
}

// `fullName` returns the full name of a person
func (p person) fullName() string {
	return p.first + " " + p.last
}

/*******************************************************************
  Go does not support inheritance, but it does support composition.
	This means that we can embed one struct inside another struct.
********************************************************************/
// embed person
type author struct {
	person         // ← embed a person type here, which is also a struct
	penName string // ← pen name is an "attribute" (field) specific to the author type
}

/*
 * (i) NOTE:
 * You can also embed an interface type inside a struct.
 * This will be explored in the chapter on interfaces.
 */
func main() {
	// initialize and print a person's full name
	p := person{first: "Stephen", last: "King"}
	fmt.Println(p.fullName())

	// initialize and print an author with an embedded person
	a := author{
		person:  p,
		penName: "Richard Bachman",
	}
	fmt.Println(a.fullName())
}

// override fullName method for author
func (a author) fullName() string {
	return fmt.Sprintf("%s (writing as %s)", a.person.fullName(), a.penName)
}
