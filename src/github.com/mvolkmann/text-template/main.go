package main

import (
	"fmt"
	"os"
	"text/template"
)

// Person describes a person.
type Person struct {
	FirstName        string
	LastName         string
	Salary           int
	PointsPerQuarter []int
	Colors           []string
	Players          map[string]string
}

func main() {
	person := Person{
		FirstName:        "Mark",
		LastName:         "Volkmann",
		Salary:           1234,
		PointsPerQuarter: []int{10, 0, 7, 17},
		Colors:           []string{"red", "yellow", "orange"},
		Players: map[string]string{
			"basketball": "Michael Jordan",
			"hockey":     "Wayne Gretzky",
			"tennis":     "Roger Federer",
		},
	}

	/*
			// Separate steps.
			myTemplate := template.New("my template")
			myTemplate, err := myTemplate.Parse(content)

		  // Checking for errors.
			myTemplate, err := template.New("my template").Parse(content)
			if err != nil {
		    fmt.Fprintln(os.Stderr, err)
		    return
		  }
	*/

	funcMap := template.FuncMap{
		"double": func(n int) int { return n * 2 },
		"sum": func(numbers []int) int {
			result := 0
			for _, n := range numbers {
				result += n
			}
			return result
		},
	}

	fileName := "template.txt"
	myTemplate := template.Must( // panics on errors
		template.
			New(fileName).
			Funcs(funcMap).
			ParseFiles(fileName))

	/*
		var buf bytes.Buffer // implements the io.Writer interface
		err := myTemplate.Execute(&buf, person)
	*/
	err := myTemplate.Execute(os.Stdout, person)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	//fmt.Println(buf.String())
}
