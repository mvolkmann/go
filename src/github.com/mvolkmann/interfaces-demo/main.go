// To run this, cd to $GOPATH and enter
// go run src/github.com/mvolkmann/interfaces-demo/main.go
package main

import (
	"fmt"
	"math"

	"github.com/mvolkmann/logger"
)

// Shape ...
type Shape interface {
	area() float64
	name() string
}

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) name() string {
	return "rectangle"
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) name() string {
	return "circle"
}
func (c circle) diameter() float64 {
	return c.radius * 2
}

func printArea(g Shape) {
	logger.LogVar("area", g.area())
}

func main() {
	//r := rectangle{width: 3, height: 4}
	var r Shape = rectangle{width: 3, height: 4}
	c := circle{radius: 5}
	var g Shape

	//printArea(g) // panic: runtime error: invalid memory address or nil pointer dereference

	g = r
	fmt.Println("g is a", g.name())

	//switch value := g.(type) {
	switch g.(type) {
	case circle:
		fmt.Println("We have a circle.")
	case rectangle:
		fmt.Println("We have a rect.")
	default:
		fmt.Println("We have something else.")
	}

	// This is a "type assertion" that asserts that
	// g refers to a circle, which it currently does not.
	//myRect := g.(circle) // triggers a panic assuming no compile errors

	// By accepting a second return value,
	// the type assertion will not panic if it fails.
	myRect, ok := g.(circle)
	logger.LogVar("myRect", myRect)
	logger.LogVar("ok", ok)

	printArea(g)
	g = c
	fmt.Println("g is a", g.name())
	printArea(g)
	fmt.Println("diameter =", g.(circle).diameter())
}
