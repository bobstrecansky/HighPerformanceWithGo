package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/Masterminds/sprig"
)

func main() {
	jsonDict := map[string]interface{}{"JSONExamples": map[string]interface{}{"foo": "bar", "bool": false, "integer": 7}}
	tpl := `{{.JSONExamples | toPrettyJson}}`
	functionMap := sprig.TxtFuncMap()
	t := template.Must(template.New("String Split").Funcs(functionMap).Parse(tpl))
	err := t.Execute(os.Stdout, jsonDict)
	if err != nil {
		fmt.Printf("Couldn't create template: %s", err)
		return
	}
}
