/**
Determine if a string contains only unique characters (a-z)

*/
package main

import "fmt"

type TestCase struct {
	exec   func(string) bool
	input  string
	output bool
}

func isUnique(word string) bool {
	var intArr ['z' - 'a' + 1]int
	for _, char := range word {
		if intArr['z'-char] != 0 {
			return false
		}
		intArr['z'-char]++
	}

	return true
}

func main() {
	TestCaseOne := TestCase{isUnique, "cat", true}
	TestCaseTwo := TestCase{isUnique, "caat", false}
	TestCaseThree := TestCase{isUnique, "", true}

	fmt.Printf("Desired: %v, Output: %v\n", TestCaseOne.output, TestCaseOne.exec(TestCaseOne.input))
	fmt.Printf("Desired: %v, Output: %v\n", TestCaseTwo.output, TestCaseTwo.exec(TestCaseTwo.input))
	fmt.Printf("Desired: %v, Output: %v\n", TestCaseThree.output, TestCaseThree.exec(TestCaseThree.input))

}
