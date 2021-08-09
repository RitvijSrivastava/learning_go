package main

import "fmt"

/**
- Zero Value of an Array is 0 at every index of the array
- Zero Value of a Slice if [nil]
*/

func main() {

	// Array declaration [size]<data_type>{values}
	grades := [3]int{10, 20, 30}
	fmt.Printf("Grades: %v\n", grades)

	// Alternative way to declare arrays (Size extrapolated from the values)
	moreGrades := [...]int{40, 50, 60}
	fmt.Printf("More Grades: %v\n", moreGrades)

	// Only declaring the array
	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "John"
	students[1] = "Holly"
	students[2] = "Jane"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #2: %v\n", students[1])
	fmt.Printf("Size: %v\n", len(students)) // Size of the array

	copyGrades := grades // ! COPIES THE ARRAY
	copyGrades[1] = 100
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Copied Array: %v\n", copyGrades)

	// 2D Array
	var mat [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}} // can also be done as "mat := ..."
	fmt.Printf("Matrix: %v\n", mat)

	// SLICES
	// Dynamic Array
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Slice: %v\n", mySlice)
	fmt.Printf("Len: %v\n", len(mySlice))
	fmt.Printf("Capacity: %v\n", cap(mySlice))

	copySlice := mySlice // COPY BY REFERENCE. (Any value changes affected in both the slices)
	copySlice[1] = 100
	fmt.Printf("Slice: %v\n", mySlice)
	fmt.Printf("Copied Slice: %v\n", copySlice)

	// Slicing operations (Can also work with arrays)
	a := mySlice[:]   // Slice of all elements
	b := mySlice[3:]  // Slice from index 3 to end
	c := mySlice[:6]  // Slice from start to index 5
	d := mySlice[3:6] // Slice from index 3 to index 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	newSlice := make([]int, 3, 100) // (<data_type>, <length>, <capacity(optional)>)
	fmt.Printf("New Slice using make(): %v\n", newSlice)

	emptySlice := []int{}
	fmt.Printf("Empty Slice size: %v\n", len(emptySlice))
	fmt.Printf("Empty Slice capacity: %v\n", cap(emptySlice))

	// Append elements to a slice
	emptySlice = append(emptySlice, 1)
	emptySlice = append(emptySlice, 2, 3, 4, 5)
	fmt.Printf("After append: %v\n", emptySlice)

	// Spread operator expands the slice.
	emptySlice = append(emptySlice, mySlice...)
	fmt.Printf("Concatinating slices: %v\n", emptySlice)

}
