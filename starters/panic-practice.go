package main

import (
	"fmt"
	"log"
)

func main() {
	/**
	Order of execution:
		func, deferred statements, panics, handle return value
	*/
	fmt.Println("Start")

	panicker()

	fmt.Println("End")
}

func panicker() {
	/**
	This function will stop executing at panic, but main will
	continue
	*/
	fmt.Println("About to panic")

	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()

	panic("Something bad happened")

	fmt.Println("Done panicking")
}
