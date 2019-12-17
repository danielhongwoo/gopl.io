// Modify dup2 to print the names of all files in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Word struct {
	count     int
	filenames []string
}

func main() {
	counts := make(map[string]*Word)
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
	for line, n := range counts {
		if n.count > 1 {
			fmt.Printf("%d\t%s\t%v\n", n.count, line, n.filenames)
		}
	}
}

func countLines(f *os.File, counts map[string]*Word) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		k := input.Text()
		_, found := counts[k]
		if found {
			counts[k].count++
			names := counts[k].filenames

			found = false
			for _, v := range names {
				if v == f.Name() {
					found = true
					break
				}
			}
			if found == false {
				counts[k].filenames = append(counts[k].filenames, f.Name())
			}
		} else {
			counts[k] = new(Word)
			counts[k].count = 1
			counts[k].filenames = []string{f.Name()}
		}
	}
}
