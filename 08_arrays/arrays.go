package main

import "fmt"

func main() {

	// Note: Default intializations
	// int -> 0, bool -> false, string -> ""

	// Declaring an array of length 4
	// array of int intialized with 0's by default
	var nums [4]int

	// get the length of the array
	fmt.Println(len(nums))

	nums[0] = 22

	fmt.Println(nums)

	var vals [3]bool // by default intialized with false
	fmt.Println(vals)

	// Single line declaration
	nums1 := [4]int{4, 31, 2, 67}

	fmt.Println(nums1)

	// 2d arrays
	matr := [2][2]int{{3, 4}, {5, 6}}

	fmt.Println(matr)

	// we use arrays when the size is predictable
	// gives us memory optimization
	// constant time access

	// The idiomatic way in go is to use slices
}
