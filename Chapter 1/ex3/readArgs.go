package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadArgs() []string {
	f, err := os.Open("args.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ReadArgs: %v\n", err)
		os.Exit(1)
	}
	input := bufio.NewScanner(f)
	var args []string
	if input.Scan() {
		args = strings.Split(input.Text(), " ")
	}
	//shuffleArray(args)
	return args
}

/*
func shuffleArray(arr []string) {
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
*/
