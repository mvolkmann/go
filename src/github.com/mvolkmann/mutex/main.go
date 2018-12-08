package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// We will need to prevent concurrent access to this slice.
var slice = make([]string, 0)

var mutex sync.Mutex
var wg sync.WaitGroup

// addString adds a given string to the slice
// after a random number of milliseconds.
func addString(s string) {
	// Generate a random duration from zero to 500 milliseconds.
	duration := time.Duration(rand.Intn(500)) * time.Millisecond

	fmt.Println("duration =", duration)
	time.Sleep(duration)

	mutex.Lock() // prevent concurrent access to slice
	slice = append(slice, s)
	fmt.Println("appended", s)
	mutex.Unlock() // finished using slice

	wg.Done() // mark this goroutine as done
}

// addStrings adds a string to the slice a given number of times.
func addStrings(s string, count int) {
	wg.Add(count) // we will create "count" new goroutines
	for n := 0; n < count; n++ {
		go addString(s)
	}
}

func main() {
	// Seed random number generation based on current time.
	rand.Seed(time.Now().UnixNano())

	addStrings("X", 5) // starts first goroutine
	addStrings("O", 3) // starts second goroutine

	wg.Wait() // wait for all the goroutines to be done

	fmt.Printf("%v\n", slice)
}
