package main

import "fmt"

func getNumbers(start int, c chan<- int) {
	for n := start; n < 10; n += 2 {
		c <- n
	}
	close(c)
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go getNumbers(1, c1) // odd numbers
	go getNumbers(2, c2) // even numbers

	n := 0
	moreEvens, moreOdds := true, true

	for {
		select {
		case n, moreOdds = <-c1:
			if moreOdds {
				fmt.Println(n, moreOdds)
			}
		case n, moreEvens = <-c2:
			if moreEvens {
				fmt.Println(n, moreEvens)
			}
		}
		if !moreEvens && !moreOdds {
			break
		}
	}
}
