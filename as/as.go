package as

import (
	"fmt"
	"reflect"
)

var stringPtrType = reflect.TypeOf((*string)(nil))

// String tries to convert obj to a string.
// It returns a string and true if obj is a string or it implements fmt.Stringer.
// If obj implements fmt.Stringer but is nil, it returns an empty string.
// Otherwise it returns false.
func String(obj any) (string, bool) {
	if obj == nil {
		return "", false
	}
	switch o := obj.(type) {
	case string:
		return o, true
	case *string:
		if o == nil {
			return "", true
		}
		return *o, true
	case fmt.Stringer:
		v := reflect.ValueOf(o)
		if v.Kind() == reflect.Pointer {
			if v.IsNil() {
				return "", true
			}
		}
		return o.String(), true
	}
	v := reflect.ValueOf(obj)
	k := v.Kind()
	if k == reflect.String {
		return v.String(), true
	}
	if k == reflect.Pointer {
		e := v.Elem()
		if e.Kind() == reflect.String {
			return e.String(), true
		}
		if v.IsNil() {
			if v.Type().ConvertibleTo(stringPtrType) {
				return "", true
			}
			return "", false
		}
		return String(e.Interface())
	}
	return "", false
}
