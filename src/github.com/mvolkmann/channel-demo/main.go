package main

import "fmt"

func getNumbers(min, max, delta int, c chan<- int) {
	for n := min; n < max; n += delta {
		c <- n
	}
	close(c)
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go getNumbers(1, 10, 2, c1) // odd numbers
	go getNumbers(2, 10, 2, c2) // even numbers

	n := 0
	moreEvens, moreOdds := true, true // TODO: Is there any reason to initialize these?

	for {
		select {
		case n, moreOdds = <-c1:
			if moreOdds {
				fmt.Println(n)
			}
		case n, moreEvens = <-c2:
			if moreEvens {
				fmt.Println(n)
			}
		}
		if !moreEvens && !moreOdds {
			break
		}
	}
}
