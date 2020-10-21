package main

import (
	"fmt"
)

func main() {
	// Pointer for struct myStruct
	var ms *myStruct
	ms = new(myStruct)
	ms.foo = 42
	fmt.Println(ms.foo)
	changeFooValuePtr(120, ms)
	fmt.Println(ms.foo)

	// Local scope myStruct, not pointer
	var os myStruct
	os.foo = 50
	fmt.Println(os.foo)
	changeFooValue(130, os)
	fmt.Println(os.foo)
	changeFooValuePtr(150, &os)
	fmt.Println(os.foo)
}

type myStruct struct {
	foo int
}

func changeFooValuePtr(newFoo int, ms *myStruct) {
	(*ms).foo = newFoo // Same as ms.foo
}

func changeFooValue(newFoo int, os myStruct) {
	os.foo = newFoo // Won't actually change in main
}
