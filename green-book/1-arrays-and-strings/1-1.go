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
	var tests []TestCase
	tests = append(tests, TestCase{isUnique, "cat", true})
	tests = append(tests, TestCase{isUnique, "dog", true})
	tests = append(tests, TestCase{isUnique, "thequickbrownfoxjumpedoverthelazydog", false})
	tests = append(tests, TestCase{isUnique, "caat", false})
	tests = append(tests, TestCase{isUnique, "", true})

	for _, test := range tests {
		desired := test.output
		actual := test.exec(test.input)
		if desired != actual {
			panic(fmt.Sprintf("Test case failed. Desired: %v, Actual: %v", desired, actual))
		}
	}
	fmt.Println("All tests passed.")
}
