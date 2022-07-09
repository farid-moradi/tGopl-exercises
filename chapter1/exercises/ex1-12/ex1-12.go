package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.Black,
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x0f, 0xff},
	color.RGBA{0xff, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0xff, 0xff}}

const (
	blackIndex = 0
	greenIndex = 1
	redIndex   = 2
	rbIndex    = 3
	rgIndex    = 4
	gbIndex    = 5
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		var nframes, delay int
		var res, cycles, size float64
		for k, v := range r.Form {
			switch k {
			case "cycles":
				cycles, _ = strconv.ParseFloat(v[0], 4)
				fmt.Println(cycles)
			case "res":
				res, _ = strconv.ParseFloat(v[0], 4)
			case "size":
				size, _ = strconv.ParseFloat(v[0], 4)
			case "nframes":
				nframes, _ = strconv.Atoi(v[0])
			case "delay":
				delay, _ = strconv.Atoi(v[0])
			}
		}
		lissajous(w, cycles, res, size, nframes, delay)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles float64, res float64, size float64, nframes int, delay int) {
	var colorIndex uint8 = 0
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*int(size)+1, 2*int(size)+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(int(size+x*size+0.5), int(size+y*size+0.5), colorIndex)
		}
		phase += 0.1
		colorIndex++
		if colorIndex == 6 {
			colorIndex = 0
		}
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
