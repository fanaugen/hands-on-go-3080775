// types/structs/fields/begin/main.go
package main

import "fmt"

// define a struct type for author
type author struct {
	first, last string
}

func main() {
	// intialize author
	king := author{first: "Stephen", last: "King"}

	// print the author
	fmt.Printf("%v\n", king)
	// print the internal representation of the struct
	fmt.Printf("%#v\n", king) // %#v is like `inspect` in ruby
	fmt.Printf("The type of local variable `king` is: %T\n", king)
}
