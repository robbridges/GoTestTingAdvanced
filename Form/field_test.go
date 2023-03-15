package Form

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFields(t *testing.T) {
	tests := []struct {
		strct interface{}
		want  []field
	}{
		{
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
		{
			strct: struct {
				FullName string
				Label    string
			}{},
			want: []field{
				{
					Label:       "FullName",
					Name:        "FullName",
					Type:        "text",
					Placeholder: "FullName",
					Value:       "",
				},
				{
					Label:       "Label",
					Name:        "Label",
					Type:        "text",
					Placeholder: "Label",
					Value:       "",
				},
			},
		},
		{
			strct: struct {
				FullName string
				Email    string
			}{
				FullName: "Rob Bridges",
				Email:    "test@Admin.com",
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
