// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to make a package-level variable "val" and assign it a value of "global"
var val = "global"


func main() {
	val := 99 // local variable "val" 

	// print the value and the type of local variable "val"
	fmtStr := "%v is a %T"
	fmt.Printf("\nfunction-locally, " + fmtStr, val, val)

	// print the value and type of the package-level variable "val"
	printPkgVal(fmtStr)

	// update the package-level variable "val" and print it again
  updatePkgVal("NEW VALUE")
	printPkgVal(fmtStr)

	// print the pointer address of the local variable "val"
	ptr := &val
	fmt.Printf("\nPointer: %v (%T)\n", ptr, ptr)
  
	// assign a value directly to the pointer address
	// of the local variable "val"
	// and print the value of the local variable "val"
	*ptr = 1001
  fmt.Println("local val is now", val)
}

func printPkgVal(fs string) {
  fmt.Printf("\npagkage-globally, " + fs, val, val)
}

func updatePkgVal(new_value string) bool {
	val = new_value
	return true
}
