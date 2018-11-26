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
	Area() float64
	Name() string
}

// Rectangle describes a rectangle.
type Rectangle struct {
	width, height float64
}

// Area returns the area of the Rectangle.
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Name returns the name of the Rectangle.
func (r Rectangle) Name() string {
	return "rectangle"
}

// Circle describes a circle.
type Circle struct {
	radius float64
}

// Area returns the area of the Rectangle.
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Name returns the name of the Circle.
func (c Circle) Name() string {
	return "circle"
}

// Diameter returns the diameter of the Circle.
func (c Circle) Diameter() float64 {
	return c.radius * 2
}

func printArea(g Shape) {
	logger.LogVar("area", g.Area())
}

func main() {
	//r := rectangle{width: 3, height: 4}
	var r Shape = Rectangle{width: 3, height: 4}
	c := Circle{radius: 5}
	var g Shape

	//printArea(g) // panic: runtime error: invalid memory address or nil pointer dereference

	g = r
	fmt.Println("g is a", g.Name())
	fmt.Printf("g currently holds a %T\n", g) // main.rectangle

	//switch value := g.(type) {
	switch g.(type) {
	case Circle:
		fmt.Println("We have a circle.")
	case Rectangle:
		fmt.Println("We have a rect.")
	default:
		fmt.Println("We have something else.")
	}

	// This is a "type assertion" that asserts that
	// g refers to a circle, which it currently does not.
	//myRect := g.(circle) // triggers a panic assuming no compile errors

	// By accepting a second return value,
	// the type assertion will not panic if it fails.
	myRect, ok := g.(Circle)
	logger.LogVar("myRect", myRect)
	logger.LogVar("ok", ok)

	printArea(g)
	g = c
	fmt.Println("g is a", g.Name())
	printArea(g)
	fmt.Println("diameter =", g.(Circle).Diameter())
}
