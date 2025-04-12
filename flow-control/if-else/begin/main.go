// flow-contorl/if-else/begin/main.go
package main

import (
	"fmt"
)

// parseOddsEvens returns two slices, one with the odd numbers and one with the even numbers
func parseOddsEvens(ints []int) (odds []int, evens []int) {
	// use a for-range loop to iterate over the incoming slice
	for _, num := range ints {
		if odd := num%2 != 0; odd {
			odds = append(odds, num)
		} else {
			evens = append(evens, num)
		}
	}
	// use the modulo operator to check if the number is odd or even and add it to the appropriate slice
	return odds, evens // implicit return also works: it’ll return the two variables declared in the function type
}

func main() {
	// invoke and print the results of parseOddsEvens
	odds, evens := parseOddsEvens([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(odds)
	fmt.Println(evens)
}
