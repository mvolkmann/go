package main

import (
	"fmt"

	"github.com/emirpasic/gods/sets/hashset"
)

func main() {
	set := hashset.New()   // empty
	set.Add(1)             // 1
	set.Add(2, 2, 3, 4, 5) // 3, 1, 2, 4, 5 (random order, duplicates ignored)
	set.Remove(4)          // 5, 3, 2, 1
	set.Remove(2, 3)       // 1, 5

	fmt.Println("contains 1?", set.Contains(1))          // true
	fmt.Println("contains 1 and 5?", set.Contains(1, 5)) // true
	fmt.Println("contains 1 and 6?", set.Contains(1, 6)) // false

	values := set.Values()
	fmt.Println("values =", values)    // []int{5,1} (random order)
	fmt.Println("empty?", set.Empty()) // false
	fmt.Println("size =", set.Size())  // 2

	// Must use type assertions to operate on values.
	sum := 0
	for _, v := range values {
		sum += v.(int)
	}
	fmt.Println("sum =", sum)

	set.Clear()                        // empty
	fmt.Println("empty?", set.Empty()) // true
	fmt.Println("size =", set.Size())  // 0
}
