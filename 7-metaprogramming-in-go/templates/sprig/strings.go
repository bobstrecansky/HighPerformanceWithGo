package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/Masterminds/sprig"
)

func main() {
	inStr := map[string]interface{}{"Name": "   -  bob smith"}
	transform := `{{.Name | trim | replace "smith" "strecansky" | trimPrefix "-" | title | repeat 10 | wrapWith 14 "\n"}}`

	functionMap := sprig.TxtFuncMap()
	t := template.Must(template.New("Name Transformation").Funcs(functionMap).Parse(transform))

	err := t.Execute(os.Stdout, inStr)
	if err != nil {
		fmt.Printf("Couldn't create template: %s", err)
		return
	}

}
