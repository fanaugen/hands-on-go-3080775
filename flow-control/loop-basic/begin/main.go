// flow-control/loop-basic/begin/main.go
package main

import "fmt"

func main() {
	// declare a string to iterate over
	str := "apple"

	// iterate over the string with basic for loop
	for i := 0; i < len(str); i++ {
		fmt.Printf("Char %d: %s\n", i, string(str[i]))
	}
}
