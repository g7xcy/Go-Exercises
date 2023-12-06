package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = iota
	blackIndex = iota
)

type Params struct {
	cycles  int
	res     float64
	size    int
	nframes int
	delay   int
}

func lissajous(out io.Writer, params Params) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: params.nframes}
	phase := 0.0
	for i := 0; i < params.nframes; i++ {
		rect := image.Rect(0, 0, 2*params.size+1, 2*params.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(params.cycles)*2*math.Pi; t += params.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(params.size+int(x*float64(params.size)+0.5), params.size+int(y*float64(params.size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, params.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
