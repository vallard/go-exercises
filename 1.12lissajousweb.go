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
	"strings"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0xf0, 0xaa, 0xf0, 0xff},
	color.RGBA{0xff, 0x22, 0x44, 0xff},
}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/debug", printDebug)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var cycles float64 = 5
	var size int = 200
	// parse the form.
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	if r.Form["cycles"] != nil {
		c := strings.Join(r.Form["cycles"], "")
		cy, err := strconv.ParseFloat(c, 64)
		if err != nil {
			fmt.Fprintf(w, "Invalid cycle value")
			return
		} else {
			cycles = cy
		}
	}
	if r.Form["size"] != nil {
		c := strings.Join(r.Form["size"], "")
		size, _ = strconv.Atoi(c)
	}
	lissajous(w, cycles, float64(size))
}

func printDebug(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func lissajous(out io.Writer, cycles float64, size float64) {
	const (
		//cycles  = 5
		res = 0.001
		// size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0 // relative frequency of oscilator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*int(size)+1, 2*int(size)+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(int(size)+int(x*size+0.5), int(size)+int(y*size+0.5), uint8(rand.Intn(len(palette))))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
