package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	f, err := os.Create("hehe.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lissajous(f)
}

func lissajous(out io.Writer) {
	palette := []color.Color{color.Black}
	const (
		cycles  = 20
		res     = 0.0005
		size    = 250
		nframes = 100
		delay   = 5
	)

	for i := 0; i < nframes; i++ {
		palette = append(palette, color.RGBA{34, 238, 239, 1})
	}

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			xcount := size + int(x*size+0.5)
			ycount := size + int(y*size+0.5)
			img.SetColorIndex(xcount, ycount, blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
