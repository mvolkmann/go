package main

import (
  "fmt"
  "os"
)

func main() {
  args := os.Args[1:]
  if len(args) != 2 {
    fmt.Println("usage: greet {first-name} {last-name}")
    os.Exit(1)
  }
  firstName := args[0]
  lastName := args[1]
  fmt.Printf("Hello %s %s!\n", firstName, lastName)
}
