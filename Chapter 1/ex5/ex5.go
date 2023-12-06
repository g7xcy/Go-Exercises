package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{ color.White, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	indexWhite = iota
	indexGreen = iota
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "lissajous: Wrong number of filename. 1 is needed but %d is given.\n", len(os.Args)-1)
		return
	}
	rand.Seed(time.Now().UTC().UnixNano())
	buffer := bytes.Buffer{}
	lissajous(&buffer)
	err := ioutil.WriteFile(os.Args[1], buffer.Bytes(), 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "lissajous: %v\n", err)
	}
}

func lissajous(out io.Writer) {
	const (
		cycle   = 5
		res     = 0.01
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycle*2*math.Pi; t += res {
			x := math.Cos(t)
			y := math.Sin(freq*t + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), indexGreen)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
