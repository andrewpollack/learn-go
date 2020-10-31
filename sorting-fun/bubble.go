package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Always set the seed (semi) randomly!!
	maxArrSize := 256

	a := MakeRange(0, maxArrSize)
	ShuffleArray(a)
	fmt.Println(a)
	BubbleSort(a)
	fmt.Println(a)
}

func BubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
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
