package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

func main() {
	text := "FooBarBaz"
	matched, err := regexp.MatchString("Bar", text)
	fmt.Println(matched, err) // true nil
	matched, err = regexp.MatchString("^Foo", text)
	fmt.Println(matched, err) // true nil
	matched, err = regexp.MatchString("Baz$", text)
	fmt.Println(matched, err) // true nil
	matched, err = regexp.MatchString("bad[", text)
	fmt.Println(matched, err) // false error parsing regexp: missing closing ]

	bytes, err := ioutil.ReadFile("haiku.txt")
	if err != nil {
		log.Fatal(err)
	}
	matched, err = regexp.Match("whole sky", bytes)
	fmt.Println(matched, err) // true nil

	// Panics if the regular expression cannot be parsed.
	bingoRE := regexp.MustCompile("^[BINGO]\\d{1,2}$")
	callout := "G57"
	matched = bingoRE.MatchString(callout)
	fmt.Println(matched, err) // true nil

	bingoRE2 := regexp.MustCompile("^([BINGO])(\\d{1,2})$")
	//matches := bingoRE2.FindStringSubmatch("G57")
	matches := bingoRE2.FindStringSubmatch("5G7")
	fmt.Printf("matches = %v\n", matches) // [G57 G 57]

	text = "ab1c23def456g"
	digitsRE := regexp.MustCompile("\\d+")
	parts := digitsRE.Split(text, -1) // -1 to return all parts
	fmt.Printf("parts = %v\n", parts) // [ab c def g]
}
