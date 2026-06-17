package main

import "fmt"

// This will throw an error as short hand can only be used inside functions
// lastName := "Singh"
var lastName string = "Singh"

func main() {
	const name string = "Gautam"

	// name = "Singh" // will throw an error as we can't assign again to const
	fmt.Println(name)

	fmt.Println(lastName)

	// const grouping
	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}
