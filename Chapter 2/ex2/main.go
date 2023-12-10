package main

import (
	"bufio"
	"ex2/unitconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
    if len(os.Args) > 1 {
        for _, arg := range os.Args[1:] {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "unitonv: %q\n", err)
                os.Exit(1)
            }
            err = Unitconv(f)
            if err != nil {
                fmt.Fprintf(os.Stderr, "unitconv: %q\n", err)
                os.Exit(1)
            }
            f.Close()
        }        
    } else {
        err := Unitconv(os.Stdin)
        if err != nil {
            fmt.Fprintf(os.Stderr, "unitconv: %q\n", err)
        }
    }

}

func Unitconv(f *os.File) error {
    input := bufio.NewScanner(f)
    for input.Scan() {
        num, err := strconv.ParseFloat(input.Text(), 64)
        if err != nil {
            return err
        }
        c := unitconv.Celsius(num)
        f := unitconv.Fahrenheit(num)
        fmt.Printf("%s = %s, %s = %s\n", c, unitconv.CTof(c), f, unitconv.FToc(f))
        
        ft := unitconv.Feet(num)
        m := unitconv.Meter(num)
        fmt.Printf("%s = %s, %s = %s\n", ft, unitconv.FtTom(ft), m, unitconv.MToft(m))

        lb := unitconv.Pound(num)
        kg := unitconv.Kilogram(num)
        fmt.Printf("%s = %s, %s = %s\n", lb, unitconv.LbTokg(lb), kg, unitconv.KgToLb(kg))
    }
    return input.Err()
}

