package main

import (
	"fmt"
	"math/rand"
	"time"
)

func task(numChan chan int) {

	for num := range numChan {
		fmt.Println("Performing task: ", num)
		time.Sleep(time.Second)
	}
}

func process(done chan bool) {
	defer func() { done <- true }()

	fmt.Println("... processing complete")
}

// Adding send and receive only channel types to restrict
// how the channel is used in this function
// Syntax: <-chan (receive only channel), chan<- (send only channel)
func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()

	for email := range emailChan {
		fmt.Println("Sending email to: ", email)
	}

}

// Channels are mostly used for the communication between goroutines
// Note: When revising, start reading this function from below
func main() {

	// ------ Example 6 ------
	chan1 := make(chan string)
	chan2 := make(chan int)

	go func() {
		chan1 <- "hello"
	}()

	go func() {
		chan2 <- 7
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Value := <-chan1:
			fmt.Println("Data received in chan1", chan1Value)
		case chan2Value := <-chan2:
			fmt.Println("Data received in chan2", chan2Value)
		}
	}

	// ------ Example 5 ------
	// A naive email sending queue implementation
	emailChan := make(chan string, 100) // This is a buffered channel
	done := make(chan bool)

	go emailSender(emailChan, done)

	for i := 1; i < 5; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}

	// Note: It's VERY IMPORTANT to close this channel or else
	// the code will be stuck in a deadlock, since the function
	// emailSender receives data in an infinite loop, the defer
	// is never called
	// A for range loop over a channel will read values sequentially
	// and will only stop when the channel is explicitly closed.
	// If you do not close it, the loop will read all buffered values
	// and then block indefinitely waiting for more data, resulting in
	// a goroutine leak or a deadlock panic.
	close(emailChan)
	<-done
	fmt.Println("All emails sent...")

	// ------ Example 4 ------
	// Note: Sending and receiving data to or from a buffered channel
	// is not blocking, therefore this code doesn't result in a deadlock

	emailReceiveChan := make(chan string, 5) // This is a buffered channel
	emailReceiveChan <- "johndoe@example.com"
	emailReceiveChan <- "new@example.com"

	fmt.Println(<-emailReceiveChan)
	fmt.Println(<-emailReceiveChan)

	// ------ Example 3 ------
	// Using channels for the synchronization of goroutines
	// (did it earlier using waitgroups: 19_waitgroups/waitgroups.go)
	// Question: Which do we use for synchronization: Waitgroups or Channels ?
	// Answer: If we have only a single goroutine, we can use channels, otherwise use waitgroups

	result := make(chan bool)

	go process(result)

	<-result // blocking code

	// ------ Example 2 ------
	// Note: this is an unbuffered channel
	numChan := make(chan int)

	go task(numChan)

	for {
		// This is also blocking code, so the main goroutine
		// will wait till data has been sent in the channel
		numChan <- rand.Intn(100)
	}

	// ------ Example 1 ------
	// messageChan := make(chan string)

	// // This code results in a deadlock
	// messageChan <- "hello there !" // Since this line is blocking

	// message := <-messageChan

	// fmt.Println(message)
}
