package main

import "fmt"

/**
- GO doesn't support OOPS, instead it has COMPOSITION. (HAS-A relationship, as in, HAS properties of that struct)
*/

type Animal struct {
	name   string
	origin string
}

type Bird struct {
	Animal // ! This is embedding Animal struct into Bird struct (Composition).

	// "Animal Animal" <- This is not Embedding since this line means adding Animal struct to Bird. (Something like aggregation)

	speedKPH float32
	canFly   bool
}

func main() {
	bird := Bird{}
	bird.name = "Emu"         // This is set using bird struct, even though it is not defined in Bird
	bird.origin = "Australia" // This is set using bird struct, even though it is not defined in Bird
	bird.speedKPH = 69
	bird.canFly = false

	fmt.Println(bird)

	// Using Literal Syntax, we need to assign values to embedded structs like given shown below.
	anotherBird := Bird{
		Animal:   Animal{name: "Emu", origin: "Australia"},
		speedKPH: 69,
		canFly:   false,
	}
	fmt.Println(anotherBird)
}
