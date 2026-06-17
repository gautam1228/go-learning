package main

import "fmt"

func main() {
	var name string = "Gautam"

	fmt.Println(name)

	// Auto type inferencing by go compiler
	var name1 = "Singh"

	fmt.Println(name1)

	// Declaration & Implicit Asignment
	name2 := "name"

	fmt.Println(name2)

	isAdult := true
	fmt.Println(isAdult)

	// Separate declaration and assignment
	var price float32 = 34.67
	fmt.Println(price)

	// printing the type of variables
	x := 32.14
	y := -32

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	// Type casting
	x1 := uint(0)    // by default would've been int32
	x2 := float32(0) // by default would've been float64

	fmt.Printf("%d: %T \n", x1, x1)
	fmt.Printf("%f: %T \n", x2, x2)

}
