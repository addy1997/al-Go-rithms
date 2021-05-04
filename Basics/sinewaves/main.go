package main

// rootVIII
// Draw sine waves iteratively and recursively.

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"math"
)

func toRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}

// drawSine iterativelly draws a sinewave with two
// complete cycles/alternations (360 * 2 == 720)
func drawSine(img draw.Image) {
	cyan := color.RGBA{100, 200, 200, 0xff}
	var x, y float64
	for i := 0; i <= 720; i++ {
		x = float64(i)
		// must convert degrees to radians:
		y = math.Sin(toRadians(x)) * 100
		y *= -1
		fmt.Printf("%f, %f\n", x, y)
		img.Set(int(x), int(y), cyan)
	}
}

// drawRecursiveSine recursively draws a sine wave.
func drawRecursiveSine(img draw.Image, x float64) {
	if x <= 720 {
		y := math.Sin(toRadians(x)) * 50
		y *= -1
		fmt.Printf("%f, %f\n", x, y)
		img.Set(int(x), int(y), color.Black)
		drawRecursiveSine(img, x+1)
	}
}

func writePNG(img draw.Image, path string) {
	buf := &bytes.Buffer{}
	err := png.Encode(buf, img)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		err = ioutil.WriteFile(path, buf.Bytes(), 0600)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}

func main() {
	img := image.NewRGBA(image.Rect(-50, -200, 800, 200))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
	// draw two sine waves and superimpose them into a single image:
	drawSine(img)
	drawRecursiveSine(img, 0)
	writePNG(img, "composition.png")
	println("Created composition.png")
}
