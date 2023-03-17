package Form

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFields(t *testing.T) {

	var nilTestTypePtr *struct {
		FullName string
		Age      int
	}

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
		"Nil-Pointers to structs should be supported": {
			strct: nilTestTypePtr,
			want: []field{
				{
					Label:       "FullName",
					Name:        "FullName",
					Type:        "text",
					Placeholder: "FullName",
					Value:       "",
				},
				{
					Label:       "Age",
					Name:        "Age",
					Type:        "text",
					Placeholder: "Age",
					Value:       0,
				},
			},
		},
		"Pointers to structs should be valid": {
			strct: &struct {
				FullName string
				Age      int
			}{
				FullName: "Rob Bridges",
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
		"Pointers within struct fields should be valid": {
			strct: struct {
				FullName *string
				Age      *int
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
					Label:       "Age",
					Name:        "Age",
					Type:        "text",
					Placeholder: "Age",
					Value:       0,
				},
			},
		},
		"Nested structs should be supported": {
			strct: struct {
				FullName string
				Address  struct {
					Street string
					Zip    int
				}
			}{
				FullName: "Rob Bridges",
				Address: struct {
					Street string
					Zip    int
				}{
					Street: "123 fake st",
					Zip:    00001,
				},
			},
			want: []field{
				{
					Label:       "FullName",
					Name:        "FullName",
					Type:        "text",
					Placeholder: "FullName",
					Value:       "",
				},
				{
					Label:       "Street",
					Name:        "Address.Street",
					Type:        "text",
					Placeholder: "Street",
					Value:       "123 fake St",
				},
				{
					Label:       "Zip",
					Name:        "Address.Zip",
					Type:        "text",
					Placeholder: "Zip",
					Value:       00001,
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

func TestFields_InvalidValues(t *testing.T) {
	tests := []struct {
		notAstruct interface{}
	}{
		{"This is a string"},
		{123},
		{nil},
	}
	// test for a panic, first by defering a function to catch the panic then call the function that we expect to panic
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%T", tc.notAstruct), func(t *testing.T) {
			defer func() {
				if err := recover(); err == nil {
					t.Errorf("Fields() did not panic: with %v", tc.notAstruct)
				}
			}()
			fields(tc.notAstruct)
		})
	}
}
