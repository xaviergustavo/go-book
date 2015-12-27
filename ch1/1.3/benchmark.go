// Exercise 1.3
// Experiment to measure the difference in running time between our potentially
// inefficient versions and the one that uses strings.Join (Section 1.6
// illustrates part of the time package and Section 11.4 shows how to write
// benchmark tests for systematic performance evaluation.)
package main

import (
	"fmt"
	"strings"
	"time"
)

var args = []string{"testing", "echo", "golang", "unix", "plan9", "java", "ruby", "python"}

func echo1(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func echo2(args []string) {
	for i, arg := range args {
		fmt.Println(i, arg)
	}
}

func benchmark(echo func(args []string)) {
	start := time.Now()
	echo(args)
	fmt.Println("Total time:", time.Since(start).String())
}

func main() {
	fmt.Println("Testing echo1:")
	benchmark(echo1)

	fmt.Println("Testing echo2:")
	benchmark(echo2)
}
