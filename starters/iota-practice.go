package main

import "fmt"

/**
Encoding information using iota
*/
const (
	errorFlag = 1 << iota

	dogOwner
	catOwner
	birdOwner
	lizardOwner
)

func main() {
	var person byte = dogOwner | catOwner | lizardOwner

	fmt.Printf("Person is formatted as: 0b%b\n", person)
	fmt.Printf("Does person own a dog? %v\n", person&dogOwner == dogOwner)
	fmt.Printf("Does person own a cat? %v\n", person&catOwner == catOwner)
	fmt.Printf("Does person own a bird? %v\n", person&birdOwner == birdOwner)
	fmt.Printf("Does person own a lizard? %v\n", person&lizardOwner == lizardOwner)
}
