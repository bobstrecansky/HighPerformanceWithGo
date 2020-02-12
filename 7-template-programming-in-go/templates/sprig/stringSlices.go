package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/Masterminds/sprig"
)

func main() {
	tpl := `{{$v := "Hands.On.High.Performance.In.Go" | splitn "." 5}}{{$v._3}}`

	functionMap := sprig.TxtFuncMap()
	t := template.Must(template.New("String Split").Funcs(functionMap).Parse(tpl))

	fmt.Print("String Split into Dict (word 3): ")
	err := t.Execute(os.Stdout, tpl)
	if err != nil {
		fmt.Printf("Couldn't create template: %s", err)
		return
	}

	alphaSort := `{{ list "Foo" "Bar" "Baz" | sortAlpha}}`
	s := template.Must(template.New("sortAlpha").Funcs(functionMap).Parse(alphaSort))
	fmt.Print("\nAlpha Tuple: ")
	alphaErr := s.Execute(os.Stdout, tpl)
	if alphaErr != nil {
		fmt.Printf("Couldn't create template: %s", err)
		return
	}

	fmt.Print("\nString Slice Functions Completed\n")
}
