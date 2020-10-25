/**
Given two strings, determine if one is a palindrome of the other

*/

package main

import "fmt"

type TestCase struct {
	exec     func(string, string) bool
	inputOne string
	inputTwo string
	output   bool
}

func isPalindrome(wordOne string, wordTwo string) bool {
	if len(wordOne) != len(wordTwo) {
		return false
	}

	var intArr ['z' - 'a' + 1]int
	for _, char := range wordOne {
		intArr['z'-char]++
	}

	for _, char := range wordTwo {
		intArr['z'-char]--
	}

	for i := 0; i < len(intArr); i++ {
		if intArr[i] != 0 {
			return false
		}
	}

	return true
}

func main() {
	var tests []TestCase
	tests = append(tests, TestCase{isPalindrome, "cat", "act", true})
	tests = append(tests, TestCase{isPalindrome, "dog", "odg", true})
	tests = append(tests, TestCase{isPalindrome, "thequickbrownfoxjumpedoverthelazydog", "atheqickbrownfoxjumpedoverthelazydog", false})
	tests = append(tests, TestCase{isPalindrome, "caat", "q", false})
	tests = append(tests, TestCase{isPalindrome, "", "", true})

	for _, test := range tests {
		desired := test.output
		actual := test.exec(test.inputOne, test.inputTwo)
		if desired != actual {
			panic(fmt.Sprintf("Test case failed. Desired: %v, Actual: %v", desired, actual))
		}
	}
	fmt.Println("All tests passed.")
}
