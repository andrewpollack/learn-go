package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"math/rand"
	"os"
	"time"
)

const w, h int = 50, 50
const maxArray = w

var images []*image.Paletted
var delays []int
var img *image.Paletted
var a []int
var grayPalette = GetGrayPalette()

func main() {
	img = image.NewPaletted(image.Rect(0, 0, w, h), grayPalette)

	rand.Seed(time.Now().UnixNano()) // Always set the seed (semi) randomly!!
	a := MakeRange(1, maxArray)
	ShuffleArray(a)

	BubbleSort(a)

	FlushGif()
}

func FlushGif() {
	f, err := os.OpenFile("bubble.gif", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})

}

func CaptureImage(a []int) {
	for x := 0; x < w; x++ {
		colVal := uint8(a[x] * int(256/len(a)))
		col := color.RGBA{
			colVal,
			colVal,
			colVal,
			255,
		}
		for y := 0; y < h; y++ {
			img.Set(x, y, col)
		}
	}

	paletted := image.NewPaletted(img.Bounds(), grayPalette)
	draw.FloydSteinberg.Draw(paletted, img.Bounds(), img, image.ZP)

	images = append(images, paletted)
	delays = append(delays, 0)
}

func AlterImage(a []int, i int, j int) {
	aICol := uint8(a[i] * int(256/len(a)))
	colI := color.RGBA{
		aICol,
		aICol,
		aICol,
		255,
	}
	aJCol := uint8(a[j] * int(256/len(a)))
	colJ := color.RGBA{
		aJCol,
		aJCol,
		aJCol,
		255,
	}

	for y := 0; y < h; y++ {
		img.Set(i, y, colI)
		img.Set(j, y, colJ)
	}

	paletted := image.NewPaletted(img.Bounds(), grayPalette)
	draw.FloydSteinberg.Draw(paletted, img.Bounds(), img, image.ZP)

	images = append(images, paletted)
	delays = append(delays, 0)
}

func BubbleSort(a []int) {
	CaptureImage(a)

	for i := 0; i < len(a); i++ {
		fmt.Println(i)
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				AlterImage(a, j, j+1)
			}
		}
	}

	images = append(images, img)
	delays = append(delays, 1000)
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

func GetGrayPalette() []color.Color {
	var grayPalette []color.Color
	for i := 0; i < 256; i++ {
		grayPalette = append(grayPalette, color.RGBA{uint8(i), uint8(i), uint8(i), 0xff})
	}
	return grayPalette
}
