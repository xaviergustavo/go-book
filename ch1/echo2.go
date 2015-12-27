// Exercise 1.2
// Modify the echo program to print the index and value
// of each of its arguments, one per line.
package ch1

import (
	"fmt"
	"os"
)

func Echo2() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}
