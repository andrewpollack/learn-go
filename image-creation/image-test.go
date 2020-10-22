package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	width := 200
	height := 100

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	green := color.RGBA{50, 255, 50, 0xff}
	red := color.RGBA{255, 50, 50, 0xff}
	purple := color.RGBA{150, 0, 255, 0xff}
	yellow := color.RGBA{255, 255, 100, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				img.Set(x, y, red)
			case x >= width/2 && y >= height/2: // lower right quadrant
				img.Set(x, y, yellow)
			case x >= width/2 && y < height/2: // upper right quadrant
				img.Set(x, y, purple)
			default:
				img.Set(x, y, green)
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
