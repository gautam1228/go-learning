package main

import "fmt"

func main() {
	var name string = "Gautam"

	fmt.Println(name)

	// Auto type inferencing by go compiler
	var name1 = "Singh"

	fmt.Println(name1)

	// Declaration & Asignment
	name2 := "name"

	fmt.Println(name2)

	isAdult := true
	fmt.Println(isAdult)

	// Separate declaration and assignment

	var price float32 = 34.67
	fmt.Println(price)
}
