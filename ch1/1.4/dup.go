// Exercise 1.4
// Modify dup2 to print the names of all files in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, names := range counts {
		if len(names) > 1 {
			fmt.Printf("%d\t%s\t%s\n", len(names), line, names)
		}
	}
}

func countLines(f *os.File, counts map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		files := counts[text]
		if files == nil {
			counts[text] = []string{}
		}
		counts[text] = append(counts[text], f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}
