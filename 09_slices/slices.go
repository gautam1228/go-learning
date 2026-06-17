package main

import (
	"fmt"
	"slices"
)

// slices -> dynamic
// most used construct in go
// + useful methods
func main() {

	// Initializng a slice
	nums := []int{}
	fmt.Println(nums)
	fmt.Println("len: ", len(nums)) // by default length is 0
	fmt.Println("cap: ", cap(nums)) // by default capacity is 0

	// add an element to this slice
	nums = append(nums, 1)

	fmt.Println(nums)
	fmt.Println("len: ", len(nums)) // by default length is 0
	fmt.Println("cap: ", cap(nums)) // by default capacity is 0

	// Note: capacity doubles each time an element is added
	// beyond the original capacity of the slice

	// using make
	vals := make([]int, 3) // slice of length 3 and capacity 3
	// vals := make([]int, 3, 5) // slice of length 3 with a capacity 5
	fmt.Println(len(vals))
	fmt.Println(vals == nil)

	// comparing two slices
	nums1 := []int{1, 2, 3}
	nums2 := []int{1, 2, 4}

	fmt.Println(slices.Equal(nums1, nums2))

	// copying one slice into another
	cp1 := make([]int, 3)
	cp2 := make([]int, 5)

	cp1[0] = 1
	cp1[1] = 2
	cp1 = append(cp1, 3)
	copy(cp2, cp1)

	fmt.Println("cp1", cp1)
	fmt.Println("cp2", cp2)

	// iterating over a slice of strings using the range keyword
	fruits := []string{"apple", "banana", "watermelon"}

	for idx, val := range fruits {
		fmt.Printf("Index: %d, Value: %s \n", idx, val)
	}

	// transforming an array
	prices := []int{1, 2, 3}

	for idx := range prices {
		prices[idx] = prices[idx] * 10
	}

	// Note: this doesn't work since val is a copy of the value
	// and not the actual value itself
	// for _, val := range prices {
	// 	val = val * 10
	// }

	fmt.Println("Final prices: ", prices)

}
