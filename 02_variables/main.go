package main

import (
	"fmt"
)

/**
SCOPING in GO:
1. Global Scope - Variables is declared at package level and is Uppercase first letter. (go automatically exports these variables)
2. Package Scope - Variables is declared at package level and is lowercase first letter.
3. Block Scope - Variables declared inside a function.
*/

/**
NAMING CONVENTIONS:
- Pascal  or camelCase
	- Capitalize Acronyms (HTTP, URL)
- As short as reasonable
	- longer names for longer lives
*/

// Variables can be declared at package level (Alternate method cannot be used here)
var k int = 40

// Multiple variables can be declared like this.
var (
	name string = "John Doe"
	age  int    = 30
)

func main() {

	/** DATA TYPES
	string
	bool (DEFUALT: false)
	int  int8  int16  int32  int64 (DEFAULT: 0)
	uint uint8 uint16 uint32 uint64 uintptr (DEFAULT: 0)
	byte - alias for uint8
	rune - alias for int32
	float32 float64
	complex64 complex128
	*/

	// ! LOCAL VARIABLES MUST ALWAYS MUST BE USED. (Otherwise UnusedVar compile time error)

	// Variable declaration is var <var_name> <data_type>
	var i int
	i = 32

	// Alternate way of declaring variables (Go will automatically detect data type)
	j := 20.0

	// For Declaring multiple variables of the same types, we can do - var <var1>, <var2>, ... <data_type>
	var a, b, c int
	a = 20
	b = 10
	c = 8
	fmt.Println(a,b,c)

	// Alternative way of declaring variables
	d, e := "hEllo", 30
	fmt.Println(d,e)

	// DEFAULT TYPES:
	// Integer - int
	// Floating  point - float64

	var k int = 50 // Shadowing the package level variable

	fmt.Println(i)
	fmt.Printf("%v, %T\n", i, i) // %v -> value, %T -> Data Type
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)


	var l int = 10
	var f float32 = float32(l) // Always explicity convert from one type to another (! NO IMPLICIT TYPE CONVERSION)
	// Use strconv package for strings

	fmt.Printf("%v, %T\n", f, f)

	var op1 int = 1
	var op2 int8 = 2
	// fmt.Println(op1 + op2) // Won't run, data types are differnt
	fmt.Println("Sum: ", op1 + int(op2))

	// ! Strings are immutable
	s := "This is a string"
	byteArray := []byte(s) // Converts string into array of bytes (uint8)
	fmt.Printf("%v, %T\n", byteArray, byteArray)


}
