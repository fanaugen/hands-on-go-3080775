// flow-control/defer-panic-recover/begin/main.go
package main

import "fmt"

func cleanup(msg string) { fmt.Println("INFO: cleaning up " + msg) }

func main() {
	// defer function calls. These will be called at the end, but last-in-first-out
	defer cleanup("first deferred")
	defer cleanup("second deferred")

	fmt.Println("INFO: doing work in main function...")

	// defer recovery. recover() is a Go built-in function
	defer func() { // IIFE (Immediately Invoked Function Expression), aka (anonymous) inline function
		if err := recover(); err != nil {
			fmt.Println("WARN: recovered from panic", err)
		}
	}()

	// panic
	panic("ERROR: uh-oh")
}
