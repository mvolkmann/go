package main

import (
	"fmt"
	"os"
	"text/template"
)

// Person describes a person.
type Person struct {
	FirstName string
	LastName  string
	Colors    []string
}

func main() {
	person := Person{
		FirstName: "Mark",
		LastName:  "Volkmann",
		Colors:    []string{"red", "yellow", "orange"},
	}

	content := `{{- $longest := "foo"}}
The favorite colors of{{" "}}
{{- .FirstName}} {{.LastName}} are
{{- range .Colors}}
	{{- if gt (len .) (len $longest)}}
		{{- $longest = .}}
  {{- end}}
	{{.}}
{{- end}}
The longest color name is {{$longest}}.`
	/*
	 */

	// Separate steps.
	//myTemplate := template.New("my template")
	//myTemplate, err := myTemplate.Parse(content)

	/*
		// Checking for errors.
			myTemplate, err := template.New("my template").Parse(content)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
	*/

	// Panic on errors.
	myTemplate := template.Must(template.New("my template").Parse(content))

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
