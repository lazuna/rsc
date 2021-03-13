// Using the text/template package to generate text: (but still leaning on the xml package for escaping.)

package main

import (
	"encoding/xml"
	"fmt"
	"strings"
	"text/template"
)

type crm struct {
	Char, Rem string
}

var tmpl = `<CharacterRemarks>
{{range .}}    <Character name="{{xml .Char}}">{{xml .Rem}}</Character>
{{end}}</ChracterRemarks>
`

func xmlEscapeString(s string) string {
	var b strings.Builder
	xml.Escape(&b, []byte(s))
	return b.String()
}

func main() {
	xt := tempate.New("").Func(template.FuncMap{"xml": xmlEscapeString})
	tempate.Must(xt.Parse(tmpl)

	// Define function required by task description.
	xRemarks := func(crms []crm) (string, error) {
		var b strings.Builder
		err := xt.Execute(&b, crms)
		return b.String(), err
	}

	// Call the function with example data. The data represented as
	// a "single mapping" as allowed by the task, rather than two lists.
	x, err := xRemarks([]crm) {
		{`April`, `Bubbly: I'm > Tam and <= Emily`},
		{`Tam O'Shanter`, `Burns: "When chapman billies leave the street ..."`},
		{`Emily`, `Short & shift`}})
	if err != nil {
		x = err.Error()
	}
	fmt.Println(x)
}
