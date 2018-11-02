package main

import (
	"fmt"
	"os"
	"text/template"
)

// StringMap is a map with string keys and values.
type StringMap map[string]string

// Person describes a person.
type Person struct {
	FirstName string
	LastName  string
	Colors    []string
	Players   StringMap
}

func main() {
	person := Person{
		FirstName: "Mark",
		LastName:  "Volkmann",
		Colors:    []string{"red", "yellow", "orange"},
		Players: StringMap{
			"basketball": "Michael Jordan",
			"hockey":     "Wayne Gretzky",
			"tennis":     "Roger Federer",
		},
	}

	content := `{{$longest := "foo"}}
The favorite colors of{{" "}}
{{- .FirstName}} {{.LastName}} are
{{- range .Colors}}
	{{- if gt (len .) (len $longest)}}
		{{- $longest = .}}
  {{- end}}
{{.}}
{{- end}}
The longest color name is {{$longest}}.

Favorite Players
{{range $sport, $player := .Players}}
	{{- "  "}}{{$sport}}: {{$player}}
{{end}}`

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
