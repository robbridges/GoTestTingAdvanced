package Form

import (
	"fmt"
	"reflect"
)

func fields(strct interface{}) []field {
	// god help me for what I'm about to do.
	rv := reflect.ValueOf(strct)
	t := rv.Type()
	var ret []field
	for i := 0; i < t.NumField(); i++ {
		tf := t.Field(i)
		rvf := rv.Field(i)
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
