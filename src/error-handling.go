package main

import (
	"errors"
	"fmt"
)

func divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

type divideByZero struct {
	numerator float32
}

//type DivideByZero float32; // alternate approach

func (err divideByZero) Error() string {
	return fmt.Sprintf("tried to divide %v by zero", err.numerator)
	//return fmt.Sprintf("tried to divide %v by zero", float32(err)) // alternate approach
}

func divide2(a, b float32) (float32, error) {
	if b == 0 {
		return 0, DivideByZero{a}
		//return 0, DivideByZero(a) // alternate approach
	}
	return a / b, nil
}

func main() {
	fn := divide2
	q, err := fn(7, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(q)
	}

	q, err = fn(7, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(q)
	}
}
