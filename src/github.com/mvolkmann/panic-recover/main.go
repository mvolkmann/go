package main

import "fmt"

func mustDivide(m, n float64) float64 {
	if n == 0 {
		panic(fmt.Sprintf("cannot divide %f by zero", m))
	}
	return m / n
}

func divideWithRecover(m, n float64) (result float64) {
	defer func() {
		message := recover()
		if message != nil {
			fmt.Printf("%s; recovering\n", message)
			result = 0
		}
	}()
	return mustDivide(m, n)
}

func main() {
	fmt.Println(divideWithRecover(7, 2))
	fmt.Println(divideWithRecover(7, 0))
}
