package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mutex sync.Mutex
var slice = make([]string, 0)
var wg sync.WaitGroup

func addString(s string) {
	mutex.Lock() // prevent concurrent access to the slice

	// Generate a random duration from zero to 500 milliseconds.
	duration := time.Duration(rand.Int63n(500)) * time.Millisecond
	fmt.Println("duration =", duration)
	time.Sleep(duration)

	slice = append(slice, s)
	fmt.Println("appended", s)

	mutex.Unlock() // finished using slice
}

// This adds a string to the slice a given number of times.
func addStrings(s string, count int) {
	for n := 0; n < count; n++ {
		addString(s)
	}
	wg.Done() // mark this goroutine as done
}

func main() {
	wg.Add(2) // we will create two new goroutines
	go addStrings("X", 5) // starts first goroutine
	go addStrings("O", 3) // starts second goroutine
	wg.Wait() // wait for the two goroutines to be done
	fmt.Printf("%v\n", slice)
}
