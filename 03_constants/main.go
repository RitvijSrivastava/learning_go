package main

import "fmt"

/**
- 'const' has SAME naming conventions as a normal variables.
- const variables must be initialized.
- It cannot be a runtime expression. (Value must be known at compile time)
*/

const a int16 = 17

// 'iota' is an enumerated constant that increases the value (follows the pattern) everytime it is called
const (
	aa = iota // 0
	bb        // 1
	cc        // 2
) // THis is how enums are implemented in GO.

// Custom values can be set
const (
	customA = 0
	customB = 2
	customC = 10
)

// '_' means read-only operator, this ignores the first value (0). This is done so as to not assign memory to the 0th value if not needed.
const (
	_    = iota
	newA // 1
	newB // 2
	newC // 3
)

// Basic Arithmetic can be used to initialize a 'const' (including bit shift operators)
const (
	_    = iota + 5
	addA // 6 (1 + 5)
	addB // 7 (2 + 5)
	addC // 8 (3 + 5)
)

const (
	bitShiftA = 1 << iota // 1 * 2^0
	bitShiftB             // 1 * 2^1
	bitShiftC             // 1 * 2^2

)

func main() {
	const myConst int = 42
	fmt.Println(myConst)

	// Alternative way to set a const. Notice the difference in the syntax.
	const myConst2 = 30
	fmt.Println(myConst2)

	const a = 16
	fmt.Println(a) // Const can be shadowed

	var b int16 = 20
	fmt.Println(a + b) // This works! Why?
	/**
	The above line works because when we don't set a data type explicitly, the compiler will
	automatically set the data type to the data type of the other operand.
	The compiler basically replaces const variables with its value, so the data type doesn't matter anymore.

	NOTICE:
	const tempa int = 20
	var tempb int16 = 20
	fmt.Println(tempa + tempb)  // This doesn;t work
	*/

	fmt.Printf("%b\n", bitShiftC) // %b -> binary representation

}
