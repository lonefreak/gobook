// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
// Modified version prints the names of all files in which each
// duplicated line occur
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filenames)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filenames)
			f.Close()
		}
	}
	for line, n := range counts {
		var sep, filelist string
		if n > 1 {
			for fn, _ := range filenames[line] {
				filelist += sep + fn
				sep = " "
			}
			fmt.Printf("%d\t%s [%s]\n", n, line, filelist)
		}
	}
}

func countLines(f *os.File, counts map[string]int, filenames map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		if text == "" {
			break
		}
		if filenames[text] == nil {
			filenames[text] = make(map[string]int)
		}
		filenames[text][f.Name()] = 1
		counts[text]++
	}
}
