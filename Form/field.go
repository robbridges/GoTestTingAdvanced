package Form

import (
	"fmt"
	"reflect"
)

func fields(strct interface{}) []field {
	// god help me for what I'm about to do.
	rv := reflect.ValueOf(strct)
	// if it's a pointer, just set it to the element that it is pointing to. Cake and pie.
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		panic("form: invalid value; only structs are supported")
	}
	t := rv.Type()
	var ret []field
	for i := 0; i < t.NumField(); i++ {
		tf := t.Field(i)
		rvf := rv.Field(i)
		if !rvf.CanInterface() {
			continue
		}
		f := field{
			Label:       tf.Name,
			Name:        tf.Name,
			Type:        "text",
			Placeholder: tf.Name,
			Value:       rvf.Interface(),
		}
		ret = append(ret, f)
	}
	fmt.Printf("%v\n", t)

	return ret
}

type field struct {
	Label       string
	Name        string
	Type        string
	Placeholder string
	Value       interface{}
}
