package main

import (
	"ex1/tempconv"
	"fmt"
)

func main() {
	var c tempconv.Celsius = 0
	fmt.Println(c, tempconv.CTof(c), tempconv.CTok(c))
	c = 100
	fmt.Println(c, tempconv.CTof(c), tempconv.CTok(c))
}
