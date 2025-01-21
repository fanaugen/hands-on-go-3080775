// packages/visibility/main.go
package main

import "fmt"

func main() {
	fmt.Println("Trying to access a private function...")
	fmt.newPrinter()
}
