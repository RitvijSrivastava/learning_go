package main

import "fmt"

func main() {
	/**
	- DEFER moves the function passed to it to run JUST BEFORE the function ends
	- DEFER statement takes in a "function call"
	- Multiple defer statements work in LIFO (Last in, first out) order.
	- Defer statement takes up the value which was defined before the function called deffered, even if the value changes afterwards.
	*/
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")

	a := "start"
	defer fmt.Println(a)
	a = "end"

	/**
	- PANIC breaks the execution of the program and immediately exits the program (Runtime Error)
	- Deferred function calls still run before the program exists.
	- Panic happens AFTER the deferred function calls.
	*/
	fmt.Println("I will print")
	panic("I panicked!")
	fmt.Println("I will NOT print")

}
