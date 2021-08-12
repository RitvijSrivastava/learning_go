package main

import "fmt"

func main() {

	// Simple Loop (initializer and updation statements are optional)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Using Multiple variables
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	// While loops
	i := 0
	for i < 5 {
		fmt.Printf("%v ", i)
		i++
	}
	fmt.Println("")

	// Break keyword
	i = 0
	for i < 5 {
		if i&1 == 1 {
			break
		}
		fmt.Printf("%v ", i)
		i++
	}
	fmt.Println("")

	// Continue keyword
	i = 0
	for i < 5 {
		if i&1 == 1 {
			i++
			continue
		}
		fmt.Printf("%v ", i)
		i++
	}
	fmt.Println("")

	// Nested for loops
	for j := 1; j <= 3; j++ {
		for k := 1; k <= 3; k++ {
			fmt.Printf("%v ", k)
		}
		fmt.Println("")
	}

	// Nested for loops with labels
Outer:
	for j := 1; j <= 3; j++ {
		for k := 1; k <= 3; k++ {
			fmt.Printf("%v ", k)
			if k == 2 {
				break Outer
			}
		}
		fmt.Println("")
	}
	fmt.Println("")

	// Looping over a Slice. (Can be user for Maps, Strings, Channels, basically anything that can be accessed using index.)
	// Except for maps, key gives index of the element.
	s := []int{1, 2, 3}
	for key, value := range s {
		fmt.Println(key, ":", value)
	}

	// ONLY KEYS
	for key := range s {
		fmt.Printf("%v ", key)
	}
	fmt.Println()

	//ONLY VALUES (Use write only variable '_')
	for _, value := range s {
		fmt.Printf("%v ", value)
	}

}
