package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Doing task: ", id)
}

func display(message string) {
	for range 3 {
		fmt.Println(message)
	}
}

func main() {
	// This is Blocking-code, runs one after the completion of the other
	// for i := 1; i <= 10; i++ {
	// 	task(i)
	// }

	// This code spawns 10 parallel lightweight threads
	// Non-Blocking Code
	for i := 1; i <= 10; i++ {
		go task(i)
	}

	time.Sleep(time.Second * 1) // naive way to handle goroutines and concurrency

	go display("Hi, Goroutine !")

	time.Sleep(time.Second) // naive way

	display("Hi, main !")
}
