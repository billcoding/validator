package funcs

import "reflect"

// maxLengthFunc struct
type maxLengthFunc struct {
	maxLength int
}

// MaxLengthFunc method
func MaxLengthFunc(maxLength int) VFunc {
	return &maxLengthFunc{maxLength}
}

// Accept method
func (f *maxLengthFunc) Accept(typ reflect.Type) bool {
	return typ.Kind() == reflect.String ||
		(typ.Kind() == reflect.Slice && typ.Elem().Kind() == reflect.String) ||
		(typ.Kind() == reflect.Array && typ.Elem().Kind() == reflect.String)
}

// Pass method
func (f *maxLengthFunc) Pass(value reflect.Value) bool {
	if f.maxLength <= 0 {
		return true
	}
	if value.Type().Kind() == reflect.String {
		return f.maxLength >= len(value.String())
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
