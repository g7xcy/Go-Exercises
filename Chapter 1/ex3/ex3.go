package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(Echo1(os.Args))
	fmt.Println(Echo2(os.Args))
}
