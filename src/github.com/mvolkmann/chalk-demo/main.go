package main

import (
	"fmt"

	"github.com/ttacon/chalk"
)

func main() {
	fmt.Printf("%s%s%s\n", chalk.Magenta, "So pretty!", chalk.Reset)

	red := chalk.Red.Color
	yellow := chalk.Yellow.Color
	blue := chalk.Blue.Color
	fmt.Println(red("Hello,"), yellow("my name is"), blue("Mark!"), chalk.Reset)
}
