package main

import "fmt"

// This function can only print a slice of ints,
// but we may as well need to print a slice of strings
func printIntSlice(items []int) {
	for _, item := range items {
		fmt.Printf("%v ", item)
	}
}

// We'll have to write a function like this
// We end up violating the DRY principle :/
// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Printf("%v ", item)
// 	}
// }

// Generics come to our rescue !

func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Printf("%v ", item)
	}
}

// We can also restric the generic type
// to a few types
func printIntOrStringSlice[T int | string](items []T) {
	for _, item := range items {
		fmt.Printf("%v ", item)
	}
}

// A naive Generic struct
type Stack[T any] struct {
	elements []T
}

// Keeping all method names lowercase as they are intended
// to be used in this file only

// Note: The more idiomatic Go way would be to return
// (value, bool), but here we are just doing a naive implementation
func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() T {
	lastIdx := len(s.elements) - 1
	topElement := s.elements[lastIdx]
	s.elements = s.elements[:lastIdx]

	return topElement
}

func (s *Stack[T]) top() T {
	return s.elements[len(s.elements)-1]
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s Stack[T]) printStack() {
	if len(s.elements) == 0 {
		fmt.Println(s.elements)
		return
	}
	for i := len(s.elements) - 1; i >= 0; i-- {
		fmt.Println(s.elements[i])
	}
}

func main() {
	// printIntSlice([]int{1, 2, 3, 4, 5, 6})
	// printStringSlice([]string{"apple", "banana", "cat", "dog"})

	printSlice([]int{1, 2, 3, 4, 5, 6})
	printSlice([]string{"apple", "banana", "cat", "dog"})

	newStack := Stack[int]{
		elements: []int{1, 2, 3, 4, 5},
	}

	fmt.Println("Original Stack (Top To Bottom):")
	newStack.printStack()

	newStack.push(7)
	fmt.Println("New Stack after pushing 7:")
	newStack.printStack()

	fmt.Println("Stack Top:", newStack.top())
	fmt.Println("Stack is empty:", newStack.isEmpty())

	fmt.Println("Removing the top element: ", newStack.pop())
	fmt.Println("New Stack after pop:")
	newStack.printStack()

	newStack.elements = []int{}
	fmt.Println("New Stack after removing all elements:")
	newStack.printStack()
	fmt.Println("Stack is empty:", newStack.isEmpty())

}
