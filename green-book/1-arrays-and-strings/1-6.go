/**
For an input string (characters a-z), return a compressed version
aaaabbccccccdddd -> a4b2c6d4
*/
package main

import (
	"fmt"
	"strconv"
)

type TestCase struct {
	exec   func(string) string
	input  string
	output string
}

func getCompressed(word string) string {
	var compressedWord string
	prevChar := ""
	prevCount := 0
	for _, char := range word {
		currChar := string(char) // conv from rune
		if prevChar != currChar {
			if prevCount >= 1 {
				compressedWord += prevChar
			}
			if prevCount > 1 {
				compressedWord += strconv.Itoa(prevCount)
			}
			prevCount = 1
		} else {
			prevCount++
		}
		prevChar = currChar
	}
	if prevCount >= 1 {
		compressedWord += prevChar
	}
	if prevCount > 1 {
		compressedWord += strconv.Itoa(prevCount)
	}

	return compressedWord
}

func main() {
	var tests []TestCase
	tests = append(tests, TestCase{getCompressed, "aaaabbccccccdddd", "a4b2c6d4"})
	tests = append(tests, TestCase{getCompressed, "dog", "dog"})
	tests = append(tests, TestCase{getCompressed, "thequickbrownfoxjumpedoverthelazydog", "thequickbrownfoxjumpedoverthelazydog"})
	tests = append(tests, TestCase{getCompressed, "caat", "ca2t"})
	tests = append(tests, TestCase{getCompressed, "", ""})

	for _, test := range tests {
		desired := test.output
		actual := test.exec(test.input)
		if desired != actual {
			panic(fmt.Sprintf("Test case failed. Desired: %v, Actual: %v", desired, actual))
		}
	}
	fmt.Println("All tests passed.")
}
