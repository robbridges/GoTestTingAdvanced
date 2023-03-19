package Form_test

import (
	"GoTestingAdvanced/Form"
	"html/template"
	"testing"
)

var (
	// this template calls template.Must if it is not created correctly it will panic. We then call in parse to create
	//in a string to parse the template
	tplTypeNameValue = template.Must(template.New("").Parse(`<input type="{{.Type}}" name="{{.Name}}"{{with .Value}} value="{{.}}"{{end}}>`))
)

func TestParseTags(t *testing.T) {

}

func TestHTML(t *testing.T) {
	tests := map[string]struct {
		tpl     *template.Template
		strct   interface{}
		want    template.HTML
		wantErr error
	}{
		"A basic form with values": {
			tpl: tplTypeNameValue,
			strct: struct {
				Name  string
				Email string
			}{
				Name:  "Michael Scott",
				Email: "michael@dundermifflin.com",
			},
			want: `<input type="text" name="Name" value="Michael Scott">` +
				`<input type="text" name="Email" value="michael@dundermifflin.com">`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := Form.HTML(tc.tpl, tc.strct)
			if err != tc.wantErr {
				t.Fatalf("HTML(tc.tpl, tc.strct) err = %v, wanted %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("HTML(tc.tpl, tc.strct) got = %q, want %q", got, tc.want)
			}
		})
	}
}
