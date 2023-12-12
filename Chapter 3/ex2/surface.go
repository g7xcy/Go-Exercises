package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

const (
	width, height = 1200, 640
	cells 		  = 100
	xyrange 	  = 30.0
	xyscale 	  = width/ 2/ xyrange
	zscale 	 	  = height* 0.4
	angle 		  = math.Pi/ 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main () {
	str := generateSVG(saddle)
	if err := ioutil.WriteFile("saddle.html", []byte(str), 0666); err != nil {
		fmt.Fprintf(os.Stderr, "surface: %v\n", err)
	}
	str = generateSVG(eggBox)
	if err := ioutil.WriteFile("eggBox.html", []byte(str), 0666); err != nil {
		fmt.Fprintf(os.Stderr, "surface: %v\n", err)
	}
	str = generateSVG(moguls)
	if err := ioutil.WriteFile("moguls.html", []byte(str), 0666); err != nil {
		fmt.Fprintf(os.Stderr, "surface: %v\n", err)
	}
}

func generateSVG(f func(float64, float64) float64) string {
	str := fmt.Sprintf(
			"<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0;j < cells; j++ {
			ax, ay, ok := corner(i+1, j, f)
			if !ok {
				continue
			}
			bx, by, ok := corner(i, j, f)
			if !ok {
				continue
			}
			cx, cy, ok := corner(i, j+1, f)
			if !ok {
				continue
			}
			dx, dy, ok := corner(i+1, j+1, f)
			if !ok {
				continue
			}
			str += fmt.Sprintf(
				"<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	return str
}

func corner(i, j int, f func(float64, float64) float64) (float64, float64, bool) {
	x := xyrange* (float64(i)/ cells- 0.5)
	y := xyrange* (float64(j)/ cells- 0.5)

	z := f(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		fmt.Println("z is Inf or NaN")
		return 0, 0, false
	}
	
	sx := width/ 2+ (x- y)*cos30 * xyscale
	sy := height/ 2+ (x+ y)*sin30 * xyscale- z* zscale
	return sx, sy, true
}

func defaultF(x, y float64) float64 {
	r := math.Hypot(x, y)

	return math.Sin(r)/ r
}

func eggBox(x, y float64) float64 {
	// z = 1/5* (sin(x)+ sin(y))
	r := (math.Sin(x)+ math.Sin(y))/ 5

	return r
}

func moguls(x, y float64) float64 {
	r := math.Pow(math.Sin(x/ xyscale* 3* math.Pi), 2)* math.Cos(y/ xyrange* 3* math.Pi)

	return r
}

func saddle(x, y float64) float64 {
	// z = x^2/ 4 - y^2/ 9
	r := math.Pow(x* 2/ xyrange, 2)- math.Pow(y* 3/ xyrange, 2)

	return r
}
