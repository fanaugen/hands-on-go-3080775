// types/arrays/begin/main.go
package main

import "fmt"

func main() {
	var a [3]int // array of size 3, data type integer

	fmt.Println(a)

	a[0] = 1

	fmt.Println(a)
}

// ⚠️ NOTE: an Array in Go cannot be resized.
//   For convenience, in practice we mostly work with slices.
//   A slice is like a sliding window onto an underlying array,
//   and it hides sizing concerns, as it’ll automatically resize.
