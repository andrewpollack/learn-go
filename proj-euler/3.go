package main

import (
	"fmt"
	"math"
)

// GetAllPrimesUpTo ...
func GetAllPrimesUpTo(maxSize int) []int {
	var primes []int
	primes = append(primes, 2)

	// Only need to check up to size sqrt(maxSize)
	for i := 3; i < int(math.Sqrt(float64(maxSize))); i += 2 {
		passCheck := true
		for j := 1; j < len(primes); j++ {
			if i%primes[j] == 0 {
				passCheck = false
				continue
			}
		}
		if passCheck {
			primes = append(primes, i)
		}
	}
	return primes
}

// GetLargestPrimeFactor ...
func GetLargestPrimeFactor(maxSize int, primes []int) int {
	largestPrime := 1
	for i := 0; i < len(primes); i++ {
		if maxSize%primes[i] == 0 {
			largestPrime = primes[i]
		}
	}
	return largestPrime
}

func main() {
	var targetVal = 600851475143
	var primes []int = GetAllPrimesUpTo(targetVal)
	var largestPrimeFactor = GetLargestPrimeFactor(targetVal, primes)
	fmt.Printf("%v's largest prime factor is: %v\n", targetVal, largestPrimeFactor)
}
