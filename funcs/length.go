package funcs

import "reflect"

//Define lengthFunc struct
type lengthFunc struct {
	length int
}

//LengthFunc
func LengthFunc(length int) VFunc {
	return &lengthFunc{length}
}

//Accept
func (f *lengthFunc) Accept(typ reflect.Type) bool {
	return typ.Kind() == reflect.String ||
		(typ.Kind() == reflect.Slice && typ.Elem().Kind() == reflect.String) ||
		(typ.Kind() == reflect.Array && typ.Elem().Kind() == reflect.String)
}

//Pass
func (f *lengthFunc) Pass(value reflect.Value) bool {
	if f.length <= 0 {
		return true
	}
	if value.Type().Kind() == reflect.String {
		return f.length == len(value.String())
	} else if value.Type().Kind() == reflect.Slice && value.Type().Elem().Kind() == reflect.String ||
		(value.Type().Kind() == reflect.Array && value.Type().Elem().Kind() == reflect.String) {
		for i := 0; i < value.Len(); i++ {
			if !f.Pass(value.Index(i)) {
				return false
			}
		}
	}
	return true
}
