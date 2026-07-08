package main

import (
	"fmt"
	"sync"
)

// Note: need to pass the WaitGroup pointer to the function so that we can call Done() on it
func task(id int, wg *sync.WaitGroup) {
	// This line runs after this function has finished executing
	// we basically decrement by 1 like we added when calling the goroutine
	defer wg.Done()
	fmt.Println("Doing task: ", id)
}

func main() {
	/*
	 * If we just let this code run without adding anything
	 * this code would exit, as it would be done
	 * before the goroutines have completed
	 * and we would not see any output.
	 */
	// for i := 1; i <= 10; i++ {
	// 	go task(i)
	// }

	// This code waits for all go routines to be finished
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()

}
