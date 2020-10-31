package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"math/rand"
	"os"
	"time"
)

func main() {
	const w, h int = 300, 300

	var images []*image.Paletted
	var delays []int
	const maxArray = 6

	rand.Seed(time.Now().UnixNano()) // Always set the seed (semi) randomly!!
	a := MakeRange(1, maxArray)
	ShuffleArray(a)
	img := image.NewPaletted(image.Rect(0, 0, w, h), palette.WebSafe)

	for !IsInOrder(a) {
		img = image.NewPaletted(image.Rect(0, 0, w, h), palette.WebSafe)
		images = append(images, img)
		delays = append(delays, 0)

		for x := 0; x < w; x++ {
			ind := int(x / (w / maxArray))
			col := color.RGBA{
				uint8(a[ind] * 25),
				uint8(a[ind] * 25),
				uint8(a[ind] * 25),
				255,
			}
			for y := 0; y < h; y++ {
				img.Set(x, y, col)
			}
		}
		ShuffleArray(a)
	}

	images = append(images, img)
	delays = append(delays, 250)

	f, err := os.OpenFile("bogo.gif", os.O_WRONLY|os.O_CREATE, 0600)
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

// IsInOrder will return true if the passed array is in order. O(n)
func IsInOrder(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			return false
		}
	}
	return true
}
