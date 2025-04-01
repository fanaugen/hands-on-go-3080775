// challenges/types-primitive/begin/main.go
package main

import "fmt"

// 1. This variable is package-global. Its type (string) is inferred at compile time
var val = "global"

// NOTE: the shorthand x := "foo" is not available on the package level
// x := foo   // <-- would be a syntax error

func main() {
	val := 99 // 2. local variable "val" is an int

	// 3. print the value and the type of local variable "val"
	fmtStr := "val is %v <%T>"
	fmt.Printf("\nfunction-locally, "+fmtStr, val, val)

	// print the value and type of the package-level variable "val"
	printGlobal(fmtStr)

	// update the package-level variable "val" and print it again
	updateGlobal("NEW VALUE")
	printGlobal(fmtStr)

	// 6. &val is a pointer to the local variable "val"
	ptr := &val
	fmt.Printf("\nptr is a pointer %v of type <%T>\n", ptr, ptr)

	// 7. *ptr is a dereferenced pointer, which is effectively
	//    a memory address of the value stored in local variable "val"
	*ptr = 1001

	// 8. demonstrate that the value in memory was updated
	fmt.Println("\nLocal val is now: ", val)
}

// 4. prints the package-global variable
func printGlobal(fs string) {
	fmt.Printf("\npagkage-globally, "+fs, val, val)
}

// 5. updates the package-global variable
func updateGlobal(new_value string) bool {
	val = new_value
	return true
}
