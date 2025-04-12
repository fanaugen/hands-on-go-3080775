// flow-control/loop-range/begin/main.go
package main

import "fmt"

func main() {
	// initialize a slice of ints
	nums := []int{123, 456, 789}

	// use for-range to iterate over the slice and print each value
	for i, n := range nums {
		fmt.Println(i, n) // first line to print should be 0 123
	}

	// declare a map of strings to ints
	var dict = map[string]int{"one": 1983, "two": 1988, "three": 1998}

	// use for-range to iterate over the map and print each key/value pair
	for key, value := range dict {
		fmt.Println(key, value)
	}
}
