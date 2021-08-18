package main

import (
	"fmt"
)

/**
- The '{}' positions are fixed and cannot be changed, like many other languages.
*/

func msg(str string) {
	fmt.Println(str)
}

// If the parameters' data type are the, they can be grouped
func msg2(str1, str2 string) {
	fmt.Println(str1, str2)
}

/** Variadic functions
RULES:
1. You can only have one.
2. It should be the last parameter
*/
func sum(values ...int) {
	fmt.Println(values)
	res := 0
	for _, v := range values {
		res += v
	}
	fmt.Println(res)
}

func returnNum() int {
	return 10
}

// Return pointer
// Go will automatically promote variable from Stack to Heap(Shared) memory.
func returnPointer() *int {
	a := 12
	return &a
}

// Named return
// define the variable in the "return area" and use that for computation, then return.
// Notice the lack of any variables after the 'return' keyword.
func namedReturn(values ...int) (result int) {
	for _, v := range values {
		result += v
	}
	return
}

// Return multiple values
func multipleValues() (int, int) {
	return 1, 2
}

type greeter struct {
	greeting string
	name     string
}

/** This is a method. '(g greeter)' tells the compiler that the function 'greet',
belongs to the struct 'greeter' and 'g' is just the name assigned to 'greeter' struct, that is,
'greeter' will known as 'g' within this method.

- Also, this is en example of pass by value, but this can be converted to a pass-by-reference by using '*' operator.
*/
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

func main() {
	msg("Hello!")

	msg2("Hello", "World")

	sum(1, 2, 3, 4, 5)

	fmt.Println(returnNum())

	fmt.Println(*returnPointer())

	fmt.Println(namedReturn(1, 2, 3, 4, 5))

	num1, num2 := multipleValues()
	fmt.Println(num1, num2)

	// ANONYMOUS FUNCTION

	// Simple Anonymous function
	func() {
		fmt.Println("Hello!")
	}()

	// Anonymous function with parameters
	func(j int) {
		fmt.Println(j)
	}(2)

	// Assign a function to a variable
	var f func(int) = func(j int) { fmt.Println("Inside a var: ", j) }
	f(3)

	// Anonymous function with return values
	var g func(int, int) (int, int) = func(i, j int) (int, int) {
		return i, j
	}
	fmt.Println(g(4, 2))

	sampleStruct := greeter{
		greeting: "Hello",
		name:     "Methods!",
	}
	sampleStruct.greet() // This is a method.
	// A function is a method when it has a context
}
