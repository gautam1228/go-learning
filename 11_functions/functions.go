package main

import "fmt"

// function as an argument
func callFunc(callable func(int) int) int {
	return callable(10)
}

// normal function
func doubleNumber(num int) int {
	return num * 2
}

// variadic function: variable number of args
func sum(nums ...int) int {
	// here nums is a slice: []int
	s := 0
	for _, val := range nums {
		s += val
	}

	// alternatively:
	// for i := 0; i < len(nums); i++ {
	// 	s += nums[i]
	// }
	return s
}

// named returns
func sum2(nums ...int) (s int, s2 int) {
	// here nums is a slice: []int
	for _, val := range nums {
		s += val
		s2 += 2 * val
	}

	// alternatively:
	// for i := 0; i < len(nums); i++ {
	// 	s += nums[i]
	// }
	return s, s2
}

func main() {
	x := callFunc(doubleNumber)
	fmt.Println(x)

	// anonymous function
	y := callFunc(func(num int) int {
		return num + 1
	})
	fmt.Println(y)

	// assigning function to a var
	z := func(coord int) int {
		return coord * 3
	}
	fmt.Printf("%T\n", z)
	fmt.Println(z(4))

	fmt.Println("Sum:", sum(1, 2, 43, 2, 11))
	fmt.Println("Sum:", sum([]int{1, 2, 4, 5}...)) // this is similar to the sread operator in JS
	// Note: this ... syntax works only with variadic functions
}
