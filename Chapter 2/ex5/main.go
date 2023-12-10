package main

import (
    "ex5/popcount"
    "fmt"
)

func main() {
    var x uint64 = 1027
    fmt.Println(popcount.PopCount1(x), popcount.PopCount2(x))
}
