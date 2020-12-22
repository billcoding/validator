package funcs

import "reflect"

//Define minLengthFunc struct
type minLengthFunc struct {
	minLength int
}

//MinLengthFunc
func MinLengthFunc(minLength int) VFunc {
	return &minLengthFunc{minLength}
}

//Accept
func (f *minLengthFunc) Accept(typ reflect.Type) bool {
	return typ.Kind() == reflect.String ||
		(typ.Kind() == reflect.Slice && typ.Elem().Kind() == reflect.String) ||
		(typ.Kind() == reflect.Array && typ.Elem().Kind() == reflect.String)
}

//Pass
func (f *minLengthFunc) Pass(value reflect.Value) bool {
	if f.minLength <= 0 {
		return true
	}
	if value.Type().Kind() == reflect.String {
		return f.minLength <= len(value.String())
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
