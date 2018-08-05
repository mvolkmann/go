// To run this, cd to $GOPATH and enter
// go run src/github.com/mvolkmann/exec-demo/main.go
package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	date, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("date = %s\n", date)
}
