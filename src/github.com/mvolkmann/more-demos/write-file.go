package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	data := []byte("This is a test.\nLine #2")
	mode := os.FileMode(0644)
	err := ioutil.WriteFile("new-file.txt", data, mode)
	if err != nil {
    log.Fatal(err)
	}
}