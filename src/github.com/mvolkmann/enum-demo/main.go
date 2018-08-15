package main

import "fmt"

const (
	red   = iota // 0
	green        // 1
	blue         // 2
)

const (
	apple  = 9        // iota = 0
	banana = 8        // iota = 1
	cherry = iota + 3 // iota = 2, value = 2 + 3 = 5
	date              // iota = 3, value = 3 + 6 = 6
)

const (
	north = iota + 1 // iota = 0, value = 0 + 1 = 1
	south            // iota = 1, value = 1 + 1 = 2
	east             // iota = 2, value = 2 + 1 = 3
	west             // iota = 3, value = 3 + 1 = 4
)

const (
	t1 = iota * 3 // iota = 0, value = 0 * 3 = 0
	t2            // iota = 1, value = 1 * 3 = 3
	t3            // iota = 2, value = 2 * 3 = 6
)

const (
	_        = iota             // iota = 0, ignore first value
	kb int64 = 1 << (10 * iota) // iota = 1, value = 1 shifted left 10 places
	mb                          // iota = 2, value = 1 shifted left 20 places
	gb                          // iota = 3, value = 1 shifted left 30 places
	tb                          // iota = 4, value = 1 shifted left 40 places
)

func main() {
	fmt.Println("colors:", red, green, blue)
	fmt.Println("fruits:", apple, banana, cherry, date)
	fmt.Println("directions:", north, south, east, west)
	fmt.Println("doubles:", t1, t2, t3)
	fmt.Println("sizes:", kb, mb, gb, tb)
}
