package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

// origin from
// http://www.golangprograms.com/how-to-add-watermark-or-merge-two-image.html

type changeable interface {
	Set(x, y int, c color.Color)
	image.Image
}

func main() {
	image1, err := os.Open("qr.png")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	first, err := png.Decode(image1)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image1.Close()

	image2, err := os.Open("gopher.png")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	second, err := png.Decode(image2)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image2.Close()

	//resize
	n_second := resize.Resize(50, 50, second, resize.NearestNeighbor)
	//resize

	offset := image.Pt(75, 75)
	b := first.Bounds()
	image3 := image.NewRGBA(b)
	draw.Draw(image3, b, first, image.ZP, draw.Src)
	draw.Draw(image3, n_second.Bounds().Add(offset), n_second, image.ZP, draw.Over)

	third, err := os.Create("result.png")
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	png.Encode(third, image3)
	defer third.Close()
}
