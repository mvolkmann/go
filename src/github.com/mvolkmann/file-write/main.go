package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func writeString(file *os.File, text string) {
	bytes, err := io.WriteString(file, text)
	check(err)
	fmt.Printf("wrote %v bytes\n", bytes)
}

func main() {
	file, err := os.Create("out-file.txt")
	check(err)
	defer file.Close()

	writeString(file, "first line\n")
	writeString(file, "second line\n")
}
