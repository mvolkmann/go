package main

import "fmt"

// Only writes to channel.
func writer(c chan<- int) {
	c <- 2
	c <- 7
}

// Only reads from channel.
func reader(c <-chan int) {
	n := <-c
	fmt.Println("got", n)
	n = <-c
	fmt.Println("got", n)
}

func main() {
	c := make(chan int) // works buffered or not
	go writer(c)        // runs in a new goroutine
	reader(c)           // runs in current goroutine
}
