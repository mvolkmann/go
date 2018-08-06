package main

import (
	"fmt"
	"reflect"
)

type any interface{}

func whatAmI(value any) {
	kind := reflect.TypeOf(value).Kind()
	switch kind {
	case reflect.Array:
		fmt.Println("I am an array.")
	case reflect.Slice:
		fmt.Println("I am a slice.")
	default:
		fmt.Println("I am something else.")
	}
}

func variadicFn(args ...int) {
	whatAmI(args)
}

func main() {
	s := []int{1, 2, 3}
	whatAmI(s)

	a := [3]int{1, 2, 3}
	whatAmI(a)

	i := 1
	whatAmI(i)

	variadicFn(1, 2, 3)
}
