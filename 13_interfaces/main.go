package main

import "fmt"

/**
- Interfaces describe behaviours, that is, they only contain function prototype.
- Zero value of an interface is [nil].

- Implementing with values vs pointers
	- Method set of values is all methods with value receivers.
	- Method set of pointer is all methods, regardless of receiver type.

- Naming Conventions:
	- Single method interface should be "<method_name>" + "er".

- Best Practices:
	- Use many, small interfaces.
	- Don't export interfaces for types that will be consumed.
	- Do export interfaces for types that will be used by package.
	- Design your functions and methods to receive interfaces whenever possible.
*/

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	myInt := IntCounter(0) // Cast int to IntCounter (due to type alias)
	var inc Incrementer = &myInt
	for i := 0; i < 5; i++ {
		fmt.Println(inc.Increment())
	}
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

// Implicitly implement the interface.
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

type Closer interface {
	Close() error
}

// Interface composed of other interfaces.
type WriterClose interface {
	Writer
	Closer
}
