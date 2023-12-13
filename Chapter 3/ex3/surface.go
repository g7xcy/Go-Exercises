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
	colorMax      = 0xff0000
	colorMin      = 0x0000ff
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

var zMax float64 = -height
var zMin float64 = height
var zVal [cells* cells]float64

func main () {
	str := fmt.Sprintf(
			"<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: gray; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0;j < cells; j++ {
			var h float64 = 0.
			ax, ay, z := corner(i+1, j)
			if isNotOK(z) {
				fmt.Fprintln(os.Stderr, "surface: z is NaN or Inf.")
				zVal[i* cells+ j] = math.NaN()
				continue
			}
			h += z
			bx, by, z := corner(i, j)
			if isNotOK(z) {
				fmt.Fprintln(os.Stderr, "surface: z is NaN or Inf.")
				zVal[i* cells+ j] = math.NaN()
				continue
			}
			h += z
			cx, cy, z := corner(i, j+1)
			if isNotOK(z) {
				fmt.Fprintln(os.Stderr, "surface: z is NaN or Inf.")
				zVal[i* cells+ j] = math.NaN()
				continue
			}
			h += z
			dx, dy, z := corner(i+1, j+1)
			if isNotOK(z) {
				fmt.Fprintln(os.Stderr, "surface: z is NaN or Inf.")
				zVal[i* cells+ j] = math.NaN()
				continue
			}
			h += z
			zVal[i* cells+ j] = h
			str += fmt.Sprintf(
				"<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style='fill:#%s' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, "%s")
		}
	}
	str = fillColor(str, zVal)

	if err := ioutil.WriteFile("svg.html", []byte(str), 0666); err != nil {
		fmt.Fprintf(os.Stderr, "surface: %v\n", err)
	}
}

func fillColor(str string, z [cells* cells]float64) string {
	// k = (colorMax- colorMin)/ ((zMax- zMin))
	// b = colorMaz- zMax* k
	k := (colorMax- colorMin)/ (zMax- zMin)
	b := colorMax- zMax* k
	var color [] interface{}
	for _, val := range zVal {
		if isNotOK(val) {
			continue
		}
		color= append(color, fmt.Sprintf("%08x", int(k* val+ b))[2:])
	}
	str = fmt.Sprintf(str, color...)
	return str
}

func isNotOK(z float64) bool {
	return math.IsInf(z, 0) || math.IsNaN(z)
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange* (float64(i)/ cells- 0.5)
	y := xyrange* (float64(j)/ cells- 0.5)
	z := f(x, y)
	if !isNotOK(z) && z > zMax {
		zMax = z
	}
	if !isNotOK(z) && z < zMin {
		zMin = z
	}
	
	sx := width/ 2+ (x- y)* cos30* xyscale
	sy := height/ 2+ (x+ y)* sin30* xyscale- z* zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)

	return math.Sin(r)/ r
}

