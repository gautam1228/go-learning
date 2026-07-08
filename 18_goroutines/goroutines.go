package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Doing task: ", id)
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

	time.Sleep(time.Millisecond * 500) // naive way to handle goroutines and concurrency
}
