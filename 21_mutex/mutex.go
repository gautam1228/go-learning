package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex // We add this for safe access to shared data
}

// Note: mu.RWMutex is used for ready heavy operations
// when we use it, there can be multiple concurrent readers
// using mu.RLock
// but there can only be one writer at a time
// if a goroutine calles mu.Lock
// it blocks all readers and writers until it's done

// Synchronous method
// func (p *post) inc() {
// 	p.views++
// }

// Concurrent method
// Note: Upon execution, this code might get
// into race conditions and end up incorrectly
// updating the value of views, since it is being
// concurrently accessed by multiple goroutines.
// func (p *post) inc(wg *sync.WaitGroup) {
// 	defer func() {
// 		wg.Done()
// 	}()

// 	p.views++
// }

// To fix that issue, we use mutex locks
// Note: always lock the line in which
// the shared state is being udpated
// not on the whole logic otherwise
// concurrency proves to be of no help
func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()

		// We add this here as a good practice
		// because there might me an error while
		// updating / reading the data in the lock
		// then unlock will never be called
		p.mu.Unlock()
	}()

	p.mu.Lock()
	p.views++
}

func main() {

	newPost := post{}

	// This is synchronous code
	// it will always work correctly
	// for range 100 {
	// 	newPost.inc()
	// }

	// This is concurrent code
	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)
		go newPost.inc(&wg)
	}

	wg.Wait()
	fmt.Println("Updated views:", newPost.views)
}
