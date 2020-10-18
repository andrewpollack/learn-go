/**
By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
*/
package main

import "fmt"

func main() {
	var aTerm int = 1
	var bTerm int = 2

	var tempTerm int = 0
	var totalSum int = 0

	for bTerm <= 4000000 {
		if bTerm&0b1 == 0b0 {
			totalSum += bTerm
		}
		tempTerm = bTerm
		bTerm += aTerm
		aTerm = tempTerm
	}

	fmt.Printf("The sum of the even-valued Fibonacci sequence terms is: %v\n", totalSum)
}
