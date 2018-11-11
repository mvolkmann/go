package main

import (
	"fmt"
	"time"

	"github.com/mvolkmann/encapsulation"
)

func main() {
	birthday := time.Date(1961, time.April, 16, 0, 0, 0, 0, time.UTC)
	me := encapsulation.NewPerson("Mark", birthday)
	fmt.Printf("%s was born on %s.\n", me.Name, me.FormattedBirthday())
	fmt.Printf("%s is %d years old.\n", me.Name, me.Age())
}
