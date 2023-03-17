package Form

import (
	"fmt"
	"reflect"
)

func valueOf(v interface{}) reflect.Value {
	var rv reflect.Value
	switch value := v.(type) {
	case reflect.Value:
		rv = value
	default:
		rv = reflect.ValueOf(v)
	}

	// With any pointers we really want to just work with their underlying
	// type.
	if rv.Kind() == reflect.Ptr {
		// The underlying type is pretty useless if it is nil, so we need to
		// instantiate a new copy of whatever that is before using it.
		if rv.IsNil() {
			rv = reflect.New(rv.Type().Elem())
		}
		rv = rv.Elem()
	}
	return rv
}

func fields(strct interface{}) []field {
	// god help me for what I'm about to do.
	rv := valueOf(strct)
	if rv.Kind() != reflect.Struct {
		panic("form: invalid value; only structs are supported")
	}
	t := rv.Type()
	var ret []field
	for i := 0; i < t.NumField(); i++ {
		tf := t.Field(i)
		rvf := valueOf(rv.Field(i))
		if !rvf.CanInterface() {
			continue
		}
		if rvf.Kind() == reflect.Struct {
			nestedFields := fields(rvf.Interface())
			for i, nf := range nestedFields {
				nestedFields[i].Name = tf.Name + "." + nf.Name
			}
			ret = append(ret, nestedFields...)
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
