// types/conversion/begin/main.go
package main

import "fmt"

func main() {
  f := 42.00001   // float
	i := uint(f)    // unsigned int ❗Precision lost❗

	fmt.Printf("%v is a %T, but %v is a %T", f, f, i, i)
}
