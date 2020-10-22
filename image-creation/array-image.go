package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Always set the seed (semi) randomly!!
	maxArrSize := 255
	a := MakeRange(1, maxArrSize)
	ShuffleArray(a)
	arrToImage(a)
}

// MakeRange will return an array built of the range [min, max]
func MakeRange(min int, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// ShuffleArray will shuffle the array in place. O(n)
func ShuffleArray(a []int) {
	for i := range a {
		// Number between [i,len(a))
		swapInd := rand.Intn(len(a)-i) + i
		a[swapInd], a[i] = a[i], a[swapInd]
	}
}

func arrToImage(a []int) {
	width := 255
	height := 500

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for x := 0; x < width; x++ {
		fmt.Println(a[x])
		cColor := color.RGBA{uint8(a[x]), 0, 0, 0xff}
		for y := 0; y < height; y++ {
			img.Set(x, y, cColor)
		}
	}

	// Encode as PNG.
	f, _ := os.Create("imagea.png")
	png.Encode(f, img)
}
