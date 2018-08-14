package main

import (
	"fmt"
	"time"
)

func finite() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}() // calling the function
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}() // calling the function

	for i := 0; i < 2; i++ { // only expecting one message from each channel
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

func numbers(c chan int, start int, delta int, sleep int) {
	n := start
	for {
		time.Sleep(time.Duration(sleep) * time.Second)
		c <- n
		n += delta
	}
}

func infinite() {
	c1 := make(chan int)
	c2 := make(chan int)
	go numbers(c1, 1, 2, 1)
	go numbers(c2, 2, 2, 2)

loop:
	for { // expecting an infinite number of messages from each channel
		select {
		case n := <-c1:
			if n > 10 {
				break loop
			}
			fmt.Println("received", n)
		case n := <-c2:
			if n > 10 {
				break loop
			}
			fmt.Println("received", n)
		}
	}
}

func main() {
	finite()
	infinite() // doesn't start until finite finishes
}
