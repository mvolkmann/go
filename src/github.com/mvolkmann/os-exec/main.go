package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// The command and each argument must be passed as a separate string.
	// Cmd objects can only be executed one.
	// To execute a command again, create a new Cmd object.
	cmd := exec.Command("ls", "-la")
	// Output returns a byte array containing
	// the output to stdout.
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Command =", string(output))

	cmd = exec.Command("ls", "-la")
	// CombinedOutput returns a byte array containing
	// the combined output to stdout and stderr.
	output, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CombinedOutput =", string(output))

	cmd = exec.Command("ls", "-la")

	// Capture output in a Buffer.
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	// Can also assign to cmd.Stdin to specify an input source.

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Run =", stdout.String())

	cmd = exec.Command("ls", "-la")
	var stdout2 bytes.Buffer
	cmd.Stdout = &stdout2
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Start =", stdout2.String())
}
