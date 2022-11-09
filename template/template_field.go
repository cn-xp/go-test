package template

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name               string
	nonExportdAgeField string
}

func TemplateFieldMain() {
	t := template.New("hello")
	t, _ = t.Parse("hello {{.Name}}!")
	p := &Person{Name: "wtt", nonExportdAgeField: "31"}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("error:", err.Error())
	}
}
