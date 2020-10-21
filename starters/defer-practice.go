package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func funTest() {
	fmt.Println("First")
	defer fmt.Println(*returnThree())
	defer fmt.Println("Fourth")
	fmt.Println("Second")
	defer fmt.Println("Third")
}

func returnThree() *int {
	var i *int = new(int)
	defer func() {
		*i++ // 3
	}()
	*i++ // 1
	*i++ // 2
	return i
}

func canonicalWebGet() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() // Closes at end, clean to do right after open

	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func main() {
	/* Defer will execute at function exit, but before
	   the function returns.
	   Defer will also function in LFFO (Last function
	   first out) similar to a stack
	*/

	funTest()
	canonicalWebGet()
}
