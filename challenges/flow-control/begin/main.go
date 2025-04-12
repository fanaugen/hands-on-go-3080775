// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("ERROR:", err)
		}
	}()

	// use os package to read the file specified as a command line argument
	file_path := os.Args[1]           // os.Args[0] is the program executable itself
	bs, err := os.ReadFile(file_path) // bs is conventional for 'byte slice'
	if err != nil {
		panic(fmt.Errorf("cannot read from file '%s'", file_path))
	}

	// convert the bytes to a string
	file_content := string(bs)

	// initialize a map to store the counts
	stats := make(map[string]int) // auto initializes values as zero

	// use the standard library utility package that can help us split the string into words
	// My naïve implementation was:
	// words := strings.Split(file_content, " ")
	words := strings.Fields(file_content) // strings.Fields splits by consecutive whitespace

	// capture the length of the words slice
	stats["words"] = len(words)

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	for _, w := range words { // each word
		for _, c := range w { // each character (Go type: rune)
			switch {
			case unicode.IsLetter(c):
				stats["letters"]++
			case unicode.IsDigit(c):
				stats["numbers"]++
			case unicode.IsPrint(c):
				stats["symbols"]++
			}
		}
	}

	// dump the map to the console using the spew package
	// spew.Dump(stats)

	// bonus: nicer printout of the stats
	fmt.Println("File '" + file_path + "' contains:")
	for kind, count := range stats {
		fmt.Printf("  %d %s\n", count, kind)
	}
}
