package main

import (
	"fmt"
	"math"

	"volkmann"
)

// Geometry ...
type Geometry interface {
	area() float64
	name() string
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) name() string {
	return "rect"
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

func printArea(g Geometry) {
	volkmann.LogVar("area", g.area())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	var g Geometry

	//printArea(g) // panic: runtime error: invalid memory address or nil pointer dereference

	g = r
	fmt.Println("g is a", g.name())

	//switch value := g.(type) {
	switch g.(type) {
	case circle:
		fmt.Println("We have a circle.")
	case rect:
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
	volkmann.LogVar("myRect", myRect)
	volkmann.LogVar("ok", ok)

	printArea(g)
	g = c
	fmt.Println("g is a", g.name())
	printArea(g)
}
