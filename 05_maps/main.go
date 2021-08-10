package main

import "fmt"

/**
- Maps in Go are similar to UNORDERED_MAPS in C++, that is, the order of keys is arbitrary.
- ONLY the data Types that can be checked for EQUALITY can be used for maps
- Underlying data is PASSED BY REFERENCE
*/

func main() {

	people := map[string]int{
		"John": 30,
		"Doe":  40,
		"Jane": 20, // This last comma is REQUIRED
	}
	fmt.Println(people)

	// Add a key to map
	people["Doe"] = 20
	fmt.Println(people["Doe"])

	// ALternative way of declaring a map
	alterPeople := make(map[string]int)
	alterPeople = map[string]int{
		"John2": 20,
		"Jane2": 20,
	}
	fmt.Println(alterPeople)

	// Length of a map
	fmt.Println("Length:", len(people))

	// Delete a key from a map
	delete(people, "Jane")
	fmt.Println(people)
	fmt.Println(len(people))

	// What happens if a non-existent key is called ?
	fmt.Println("Deleted key:", people["Jane"]) // The zero value of 'value' is printed

	// To Actually know if the value exists or not, we can use the ", ok" syntax
	// The [ok] stores whether the value if present or not (true/false)
	person, ok := people["Jane"]
	fmt.Println("Deleted Value (comma, ok syntax):", person, ok)

}
