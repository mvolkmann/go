package main

import (
	"fmt"
	"time"
)

var x = 1

const blackjack = 21

type any interface{}

//func log(args ...interface{}) {
func log(args ...any) {
	fmt.Println(args)
}
func logInts(args ...int) {
	for _, v := range args {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func demo() {
	x = 3
	log(x, x*2)

	type person struct {
		name string
		age  int8
	}

	p := Person{}
}

func myAsync(done chan<- bool) {
	// Do some asynchronous thing.
	time.Sleep(time.Second * 3)
	// When it completes, do this:
	done <- true
}

func selectDemo() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	//for i := 0; i < 2; i++ {
myLabel:
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
			//return; // works
			break myLabel // error
		}
	}
}

func main() {
	/*
		x = 2
		demo()
		log(x)

		ptr := &x
		*ptr = 4
		log(*ptr)

		log(blackjack)

		// Create an array of ints.
		rgb := [3]int{100, 50, 234}
		log(rgb[1])
		rgb[1] = 75
		log(rgb)
		logInts(rgb[:]...)

		mySlice := []int{2, 4, 6}
		log(mySlice)

		mySlice = append(mySlice, 8, 10)
		log(mySlice)

		double := func(v int) int { return v * 2 }
		log(mapOverInts(rgb[:], double))

		println("text", '%', 19, true)

		name := "Mark"
		println(len(name))
	*/

	/*
			println("calling myAsync")
			done := make(chan bool)
			go myAsync(done)
		  <-done
			println("returned from myAsync")
	*/

	selectDemo()

	var me = struct {
		name string
		age  int8
	}{
		"Mark",
		57,
	}
	fmt.Println(me.name) // Mark
	me.age++
	fmt.Println(me.age) // 58

	/*
		days := []string{"Sat", "Sun"}
		log(len(days))
		log(cap(days))
	*/
}
