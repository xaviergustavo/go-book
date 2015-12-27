// Exercise 1.1
// Modify the echo program to also print os.Args[0],
// the name of the command that invoked it.
package ch1

import (
	"fmt"
	"os"
	"strings"
)

func Echo1() {
	fmt.Println(strings.Join(os.Args, " "))
}
