package Form

import (
	"html/template"
	"strings"
)

// HTML is used to generate HTML forms/inputs from Go structs. Given a
// template that looks something like this:
//
//     <input type="{{.Type}}" name="{{.Name}}" {{with .Value}}value="{{.}}"{{end}}>
//
// And a struct like this:
//
//     struct{
//	     Name string
//       Email string
//     }{
// 			 Name: "Michael Scott",
//       Email: "michael@dundermifflin.com",
//     }
//
// The HTML function will return a template.HTML of:
//
//     <input type="text" name="Name" value="Michael Scott">
//     <input type="text" name="Email" value="michael@dundermifflin.com">
//
// An example similar to this is shown as the first test case in TestHTML
// in the html_test.go source file.

func HTML(t *template.Template, strct interface{}) (template.HTML, error) {
	var inputs []string
	for _, field := range fields(strct) {
		var sb strings.Builder
		err := t.Execute(&sb, field)
		if err != nil {
			return "", err
		}
		inputs = append(inputs, sb.String())
	}
	return template.HTML(strings.Join(inputs, "")), nil
}
