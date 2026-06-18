package main

import "fmt"

// map -> key, val store
func main() {

	// classic syntax: very verbose
	var mp map[int]string = map[int]string{1: "a"}
	fmt.Println(mp)

	// implicit assignment
	mp1 := map[int]string{2: "b"}
	fmt.Println(mp1)

	// using make function
	mp2 := make(map[int]string)

	// adding a key, val to the map
	mp2[3] = "c"
	mp2[4] = "d"
	fmt.Println(mp2)

	// deleting a key from the map
	delete(mp2, 3)
	fmt.Println("After deleting: ", mp2)

	// checking if a key exists in a map
	val, ok := mp2[3]
	if ok {
		fmt.Println("key exists with value: ", val)
	} else {
		fmt.Println("key doesn't exist")
	}
}
