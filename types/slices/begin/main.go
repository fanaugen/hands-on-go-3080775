// types/slices/begin/main.go
package main

import "fmt"

func main() {
	// declare a slice of strings using the make built-in function
	names := make([]string, 0) // initial length is zero, not the same as capacity

	// append 3 names to the slice
	names = append(names, "Alice")
	names = append(names, "Bob")
	names = append(names, "Claire")

	fmt.Println(names)
	fmt.Printf("%#v\n", names)

	// slice the slice using syntax slice[low:high]
	fmt.Println(names[1:3]) // [Bob Claire]
	fmt.Println(names[1:])  // [Bob Claire]
	fmt.Println(names[:1])  // [Alice]
	fmt.Println(names[:])   // [Alice Bob Claire]
}
