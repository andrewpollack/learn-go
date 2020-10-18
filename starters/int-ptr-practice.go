package main

import "fmt"

func main() {
	var p *int = new(int)
	b := new(int)
	*p = 3
	*b = 2

	fmt.Printf("%v %v\n", p, b)
	fmt.Printf("%v %v\n", *p, *b)
}
