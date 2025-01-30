// types/variables/begin/main.go
package main

import "fmt"

// declare package-level variables of type int
var d, e, f int

// declare package-level variables of type bool and override the default values (also known as "zero")
var x, y, z bool = false, true, false

func main() {
	// print ints
	fmt.Println("d, e, and f have the values: ", d, e, f)

	// override the default value of a package-level variable
	d = 1_000_000
	fmt.Printf("d is now: %d\n", d)

	// declare and initialize variables of similar names as the booleans but of local scope
	x, y, z := 0, 1.25, "hello"
	fmt.Println("x, y, and z are:", x, y, z)

	// print the package-level variables x, y, and z
	printVars()
}

func printVars() {
	fmt.Println("package-level x, y z are:", x, y, z)
}