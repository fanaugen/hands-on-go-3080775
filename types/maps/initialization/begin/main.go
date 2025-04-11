// types/maps/initialization/begin/main.go
package main

import "fmt"

type author struct {
	name string
}

func main() {
	// declare a map of string keys and author values
	var authors map[string]author
	// initialize the map with make
	authors = make(map[string]author)

	// add authors to the map
	authors["sk"] = author{name: "Stephen King"}
	authors["jk"] = author{name: "J.K. Rowling"}

	// print the map with fmt.Printf
	fmt.Printf("%v\n", authors)

	// read a value from the map with a known key
	fmt.Printf("Author with key 'sk': %v\n", authors["sk"])

	// now, declaration and initialization in one, using a composite literal
	reviewers := map[string]author{
		"sk": {name: "Stephen King"}, // note that due to the type of this map, placing 'author' before the {...} would be redundant
		"jk": {name: "J.K. Rowling"},
	}
	fmt.Println(reviewers)

	// if the key is not found, the zero value of the value type is returned ⚠️
	fmt.Printf("AZ: %v\n", reviewers["az"]) // prints AZ: {} ⚠️

	// To avoid this: use the 'comma ok' idiom:
	reviewer, ok := reviewers["az"]
	if ok {
		fmt.Printf("AZ: %v\n", reviewer)
	} else {
		fmt.Println("reviewer with initials AZ not in the dictionary")
	}
}
