package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	names := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "", names)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg, names)
			f.Close()
		}
	}

	for line, num := range counts {
		if num > 1 {
			fmt.Printf("%d\t%s", num, line)
			for name, inFile := range names[line] {
				if inFile {
					fmt.Printf("\t%s", name)
				}
			}
			fmt.Print("\n")
		}
	}
}

func countLines(f *os.File, counts map[string]int, arg string, names map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if names[input.Text()] == nil {
			names[input.Text()] = make(map[string]bool)
		}
		if arg != "" {
			names[input.Text()][arg] = true
		}
		}
}

