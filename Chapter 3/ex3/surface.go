package main

import (
	"archive/zip"
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
			ax, ay, ok := corner(i+1, j)
			if !ok {
				fmt.Fprintln(os.Stderr, "surface: z is NaN or Inf.")
				zVal[i* cells+ j] = math.NaN()
				continue
			}
			bx, by, ok := corner(i, j)
			if !ok {
				fmt.Fprintln(os.Stderr, "surface: z is NaN or Inf.")
				zVal[i* cells+ j] = math.NaN()
				continue
			}
			cx, cy, ok := corner(i, j+1)
			if !ok {
				fmt.Fprintln(os.Stderr, "surface: z is NaN or Inf.")
				zVal[i* cells+ j] = math.NaN()
				continue
			}
			dx, dy, ok := corner(i+1, j+1)
			if !ok {
				fmt.Fprintln(os.Stderr, "surface: z is NaN or Inf.")
				zVal[i* cells+ j] = math.NaN()
				continue
            }
            x := xyrange* ((float64(i)+ 0.5)/ cells- 0.5)
            y := xyrange* ((float64(j)+ 0.5)/ cells- 0.5)
            z := f(x, y)
			zVal[i* cells+ j] = z
            // NEVER DO THIS IN REAL PRODUCTION ENVIRONMENTS
            // THIS MAY CALL INJECTION ATTACK
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

func isNaNOrInf(x float64) bool {
    return math.IsNaN(x) || math.IsInf(x, 0)
}

func fillColor(str string, z [cells* cells]float64) string {
	// k = (colorMax- colorMin)/ ((zMax- zMin))
	// b = colorMaz- zMax* k
	k := (0xff- 0)/ (zMax- zMin)
	b := 0xff- zMax* k
	var color [] interface{}
	for _, val := range zVal {
		if isNaNOrInf(val) {
			continue
		}
        c := int(k* val+ b)
        if c > 0xff {
            c = 0xff
        }
        if c < 0 {
            c = 0
        }
        color = append(color, fmt.Sprintf("%02x00",c)+ fmt.Sprintf("%02x", 0xff- c))
    }
    str = fmt.Sprintf(str, color...)
	return str
}

func corner(i, j int) (float64, float64, bool) {
	x := xyrange* (float64(i)/ cells- 0.5)
	y := xyrange* (float64(j)/ cells- 0.5)
	z := f(x, y)
    if isNaNOrInf(z) {
        return 0, 0, false
    }
    if z > zMax {
		zMax = z
	}
	if z < zMin {
		zMin = z
	}
	
	sx := width/ 2+ (x- y)* cos30* xyscale
	sy := height/ 2+ (x+ y)* sin30* xyscale- z* zscale
	return sx, sy, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)

	return math.Sin(r)/ r
}

