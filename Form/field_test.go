package Form

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFields(t *testing.T) {
	tests := map[string]struct {
		strct interface{}
		want  []field
	}{
		"No Values": {
			strct: struct {
				FullName string
			}{},
			want: []field{
				{
					Label:       "FullName",
					Name:        "FullName",
					Type:        "text",
					Placeholder: "FullName",
					Value:       "",
				},
			},
		},
		"Values Added": {
			strct: struct {
				FullName string
				Email    string
				Age      int
			}{
				FullName: "Rob Bridges",
				Email:    "test@Admin.com",
				Age:      123,
			},
			want: []field{
				{
					Label:       "FullName",
					Name:        "FullName",
					Type:        "text",
					Placeholder: "FullName",
					Value:       "Rob Bridges",
				},
				{
					Label:       "Email",
					Name:        "Email",
					Type:        "text",
					Placeholder: "Email",
					Value:       "test@Admin.com",
				},
				{
					Label:       "Age",
					Name:        "Age",
					Type:        "text",
					Placeholder: "Age",
					Value:       123,
				},
			},
		},
		"Unexported fields should be skipped": {
			strct: struct {
				FullName string
				email    string
				Age      int
			}{
				FullName: "Rob Bridges",
				email:    "test@Admin.com",
				Age:      123,
			},
			want: []field{
				{
					Label:       "FullName",
					Name:        "FullName",
					Type:        "text",
					Placeholder: "FullName",
					Value:       "Rob Bridges",
				},
				{
					Label:       "Age",
					Name:        "Age",
					Type:        "text",
					Placeholder: "Age",
					Value:       123,
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%T", tc.strct), func(t *testing.T) {
			got := fields(tc.strct)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("fields() = %v, want %v", got, tc.want)
			}
		})
	}
}
