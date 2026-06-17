package main

import "fmt"

// for loop is the only loop in go
func main() {

	// Creating a while loop
	i := 0
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// Classic for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// infinite loop
	// for {
	// 	fmt.Println(1)
	// }

	// break and continue work in the same way as usual
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	for i := 1; i <= 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	// range keyword
	// as the range keyword goes (range n) : indexes 0 -> n - 1
	for i := range 10 {
		fmt.Println(i)
	}
}
