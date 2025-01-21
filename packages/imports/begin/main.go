// packages/imports/begin/main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	// use the fmt package to print the string "Hello Gopher!"
	fmt.Println("Hello Gopher world!")

	// use the time package to print the current weekday
	fmt.Printf("It happened on a %s", time.Now().Weekday())
}
