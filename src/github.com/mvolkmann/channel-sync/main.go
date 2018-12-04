package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Sends random integers from 1 to 10 to the "numbers" channel
// and sends true to the "done" channel when zero is generated.
func myAsync(numbers chan<- int, done chan<- bool) {
	for {
		n := rand.Intn(11) // 0 to 10
		if n == 0 {
			done <- true
			break
		} else {
			numbers <- n
		}
	}
}

func main() {
	// Seed random number generation based on the current time.
	rand.Seed(int64(time.Now().Nanosecond()))

	// Create all the channels to be used.
	numbers1 := make(chan int)
	numbers2 := make(chan int)
	done := make(chan bool)

	// Start two goroutines to generate random numbers.
	go myAsync(numbers1, done)
	go myAsync(numbers2, done)

	// Print the random numbers as they arrive.
loop:
	for {
		select {
		case n := <-numbers1:
			fmt.Println("from first,", n)
		case n := <-numbers2:
			fmt.Println("from second,", n)
		case <-done:
			break loop // without label, just breaks from select
		}
	}
	fmt.Println("done")
}
