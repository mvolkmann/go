package main

import (
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
    log.Fatal(e)
	}
}

func writeLine(file *os.File, text string) {
	bytes, err := file.Write([]byte(text + "\n"))
	check(err)
	fmt.Printf("wrote %v bytes\n", bytes)
}

func main() {
	var (
		file *os.File
		err error
	)

	file, err = os.Create("out-file.txt")
	check(err)
	defer file.Close()

	writeLine(file, "Line #1")
	writeLine(file, "Line #2")
}