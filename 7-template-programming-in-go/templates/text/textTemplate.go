package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	type ToField struct {
		Date     string
		Name     string
		Email    string
		InOffice bool
	}

	const note = `
{{/* we can trim whitespace with a {- or a -} respectively */}}
Date: {{- .Date}}
To: {{- .Email | printf "%s"}}
{{.Name}},
{{if .InOffice }}
Thank you for your input yesterday at our meeting.  We are going to go ahead with what you've suggested.
{{- else }}
We were able to get results in our meeting yesterday.  I've emailed them to you.  Enjoy the rest of your time Out of Office!
{{- end}}
Thanks,
Bob
`
	var tofield = []ToField{
		{"07-19-2019", "Mx. Boss", "boss@example.com", true},
		{"07-19-2019", "Mx. Coworker", "coworker@example.com", false},
	}
	t := template.Must(template.New("Email Body").Parse(note))
	for _, k := range tofield {
		err := t.Execute(os.Stdout, k)
		if err != nil {
			fmt.Print(err)
		}
	}
}
