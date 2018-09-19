package main

import (
	"fmt"
	"log"
	"time"
)

func startTimer(name string) func() {
	t := time.Now()
	log.Println(name, "started")
	return func() {
		delta := time.Now().Sub(t)
		log.Println(name, "took", delta)
	}
}

func doWork() int {
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum += i
	}
	return sum
}

func main() {
	defer startTimer("doWork")()
	fmt.Println(doWork())
}
