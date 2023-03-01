package underscore

import "testing"

func TestCamel(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Simple test", args{"thisIsACamelCaseString"}, "this_is_a_camel_case_string"},
		{"contains space", args{"with a space"}, "with a space"},
		{"ends with Capital", args{"endsWithA"}, "ends_with_a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Camel(tt.args.str); got != tt.want {
				t.Fatalf("Camel() = %v, want %v", got, tt.want)
			}
		})
	}
}
