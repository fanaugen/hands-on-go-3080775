// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to make a package-level variable "val" and assign it a value of "global"
var val = "global"


func main() {
	val := 99 // local variable "val" is an int 

	// print the value and the type of local variable "val"
	fmtStr := "val is %v <%T>"
	fmt.Printf("\nfunction-locally, " + fmtStr, val, val)

	// print the value and type of the package-level variable "val"
	printGlobal(fmtStr)

	// update the package-level variable "val" and print it again
  updateGlobal("NEW VALUE")
	printGlobal(fmtStr)

	// print the pointer address of the local variable "val"
	ptr := &val
	fmt.Printf("\nPointer: %v <%T>\n", ptr, ptr)
  
	// assign a value directly to the pointer address
	// of the local variable "val"
	// and print the value of the local variable "val"
	*ptr = 1001
  fmt.Println("\nlocal val is now", val)
}

func printGlobal(fs string) {
  fmt.Printf("\npagkage-globally, " + fs, val, val)
}

func updateGlobal(new_value string) bool {
	val = new_value
	return true
}
