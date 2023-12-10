package main

import (
    "ex3/popcount"
    "fmt"
)

func main() {
    var x uint64 = 1025
    fmt.Println(popcount.PopCount1(x), popcount.PopCount2(x))
}
