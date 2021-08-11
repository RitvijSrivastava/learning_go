package main

import "fmt"

func main() {

	// Notice that parans are not used. ALso {} are REQUIRED
	if true {
		fmt.Println("True")
	}

	// Initializer statement. (Perform some operation before checking a condition)
	if i := 10; i > 0 && i < 5 {
		fmt.Println("IF CLAUSE")
	} else if i > 100 {
		fmt.Println("ELSE-IF CLAUSE ")
	} else {
		fmt.Println("ELSE")
	}

	// Cases must be Unique.
	// Initializer is optional
	// This is the TAG Syntax
	switch tmp := 2; tmp {
	case 1, 5, 6:
		fmt.Println("Case 1")
	case 2:
		fmt.Println("Case 2")
	default:
		fmt.Println("Default")
	}

	// TAGLESS SYNTAX
	i := 10
	switch {
	case i <= 10:
		fmt.Println("CASE 1")
	case i <= 20:
		fmt.Println("CASE 2")
	default:
		fmt.Println("DEFAULT")
	}

	// Fallthrough (Will fall through even if subsequent cases do not fulfill the condition)
	// Fallthrough only moves to the NEXT statement.
	i = 20
	switch {
	case i <= 20:
		fmt.Println("LESS THAN 20")
		fallthrough
	case i > 25:
		fmt.Println("Greater than 20")
	default:
		fmt.Println("DEFAULT")
	}

	// TYPE SWITCHES
	// - Requires an EMPTY  'interface'.
	var j interface{} = "ABC"
	switch j.(type) {
	case int:
		fmt.Println("int")
		break // Used to break out of execution
		fmt.Println("This will not run!")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("Not Applicable")
	}

}
