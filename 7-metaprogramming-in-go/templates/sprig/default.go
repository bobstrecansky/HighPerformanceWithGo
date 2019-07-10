package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/Masterminds/sprig"
)

func main() {

	emptyTemplate := map[string]interface{}{"Name": ""}
	fullTemplate := map[string]interface{}{"Name": "Bob"}
	tpl := `{{empty .Name}}`
	functionMap := sprig.TxtFuncMap()
	t := template.Must(template.New("Empty String").Funcs(functionMap).Parse(tpl))
	fmt.Print("empty template: ")
	emptyErr := t.Execute(os.Stdout, emptyTemplate)
	if emptyErr != nil {
		fmt.Printf("Couldn't create template: %s", emptyErr)
		return
	}

	fmt.Print("\nfull template: ")
	fullErr := t.Execute(os.Stdout, fullTemplate)
	if emptyErr != nil {
		fmt.Printf("Couldn't create template: %s", fullErr)
		return
	}
	fmt.Print("\nEmpty Check Completed\n")

}
