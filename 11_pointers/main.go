package main

import "fmt"

/**
- No Pointer Arithmetics.
- Zero Value of a pointer is 'nil'.
*/

type myStruct struct {
	a int
}

func main() {
	var a int = 42
	var b *int = &a

	fmt.Println(a, *b)
	*b = 30
	fmt.Println(a, *b)

	var tmpStruct *myStruct
	tmpStruct = new(myStruct)
	(*tmpStruct).a = 30
	fmt.Println((*tmpStruct).a)

	// The above two operations can also be done as: (Go will interpret correctly)
	tmpStruct.a = 42
	fmt.Println(tmpStruct.a)

}
