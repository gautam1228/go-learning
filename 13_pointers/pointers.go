package main

import "fmt"

func changeNum(num int) {
	num += 1

	fmt.Println("Num value inside changeNum: ", num)
}

// Passing the var by reference
func changeNumWithPtr(num *int) {
	*num += 1

	fmt.Println("Num value inside changeNumWithPtr: ", *num)
}

func main() {
	num := 1

	// changeNum(num)
	fmt.Println("Memory address of num: ", &num) // Keeps changing with every run
	changeNumWithPtr(&num)

	// fmt.Println("Num value after calling changeNum func: ", num) // stays the same as declared
	fmt.Println("Num value after calling changeNumWithPtr func: ", num) // changed because we used the ptr to the location of the var in mem
}

// NOTE: By default Go passes vars by reference
