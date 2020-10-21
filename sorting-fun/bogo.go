package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Always set the seed (semi) randomly!!
	maxArrSize := 10
	numTries := 1000

	for arrSize := 1; arrSize < maxArrSize; arrSize++ {
		var attempts float64 = 0

		for tries := 0; tries < numTries; tries++ {
			// Going for an average over many tries to avoid outlier results
			a := MakeRange(1, arrSize)
			ShuffleArray(a)

			for !IsInOrder(a) {
				ShuffleArray(a)
				attempts++
			}
		}
		fmt.Printf("Array size %v took on average %.0f attempts to sort\n", arrSize, attempts/float64(numTries))
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
	for i := 0; i < len(a); i++ {
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
