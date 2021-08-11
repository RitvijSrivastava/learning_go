package main

import (
	"fmt"
	"reflect" // Required for tags
)

/**
- Structs are passed by VALUE.
*/

// THIS struct is available globally but the fields are not available (lowercase first letter)
type Student struct {
	name           string `max:"100"` // This is tag
	registrationID int
	subjects       []string
}

// This is available GLOBALLY, including the fields (uppercase first letter)
type GlobalStudent struct {
	Name           string
	RegistrationID int
	Subjects       []string
}

func main() {
	student := Student{
		name:           "John Doe",
		registrationID: 100,
		subjects: []string{
			"ACD",
			"DSA",
		},
	}
	fmt.Println("Student Struct:", student)
	fmt.Println("Student name:", student.name)
	fmt.Println("Student registration ID:", student.registrationID)
	fmt.Println("Student subjects:", student.subjects)

	// Alternative way to declare struct. (Using posiional argument assignemnt)
	// NOT RECOMMENDED
	student2 := Student{
		"Jane Doe",
		200,
		[]string{
			"ACD",
			"Computer Graphics",
		},
	}
	fmt.Println("Student2:", student2)

	// Anonymous struct
	aDoctor := struct{ name string }{name: "Doctor doct"}
	fmt.Println("Anon Struct:", aDoctor)

	// All tags do is provide a string. What to do with it is the job of the validation framework
	t := reflect.TypeOf(Student{})
	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag)
}
