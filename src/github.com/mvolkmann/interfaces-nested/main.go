package main

import "fmt"

type shape2d interface {
	area() float32
}

type shape3d interface {
	shape2d
	volume() float32
}

type box struct {
	d float32
	h float32
	w float32
}

// Implements shape2d interface.
// Accepting a pointer so a copy is not made.
func (b *box) area() float32 {
	return b.w * b.h
}

// Implements shape3d interface.
func (b *box) volume() float32 {
	return b.area() * b.d
}

func main() {
	myBox := box{d: 2, h: 3, w: 4}
	fmt.Println("area =", myBox.area())     // 12
	fmt.Println("volume =", myBox.volume()) // 24
}
