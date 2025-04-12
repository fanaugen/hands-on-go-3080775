// flow-control/loop-infinite/main.go
package main

import (
	"fmt"
	"time"
)

func main() {

	i := 0
	for { // a loop without any condition is an endless loop
		// this will keep repeating, using 100% CPU until interrupted
		i++
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
