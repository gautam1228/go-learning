package main

import "fmt"

func main() {
	// can declare and assign the variable within the conditional
	if age := 22; age >= 18 {
		fmt.Println("Person is an adult.")
	} else if age >= 13 {
		fmt.Println("Person is a teenager.")
	} else {
		fmt.Println("Person is a child.")
	}

	// Note: there is no ternary operator in go
}
