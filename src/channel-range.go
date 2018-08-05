˜package main

import "fmt"

func getWords(c chan string) {
	c <- "alpha"
	c <- "beta"
	c <- "gamma"
	close(c)
}

func main() {
	c := make(chan string)
	go getWords(c)
	for word := range c {
		fmt.Println(word)
	}
}
