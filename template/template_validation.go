package template

import (
	"fmt"
	"text/template"
)

func TemplateValidationMain() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("error:", x)
		}
	}()
	tOk := template.New("ok")
	template.Must(tOk.Parse("/* and a comment */ some static text: {{ .Name }}"))
	fmt.Println("the first one parsed OK.")
	fmt.Println("the next one ought to fail.")
	tErr := template.New("error_template")
	template.Must(tErr.Parse(" some static text {{ .Name }"))
}
