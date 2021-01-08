package funcs

import "reflect"

// fixedFunc struct
type fixedFunc struct {
	fixed string
}

// FixedFunc method
func FixedFunc(fixed string) VFunc {
	return &fixedFunc{fixed}
}

// Accept method
func (f *fixedFunc) Accept(typ reflect.Type) bool {
	return typ.Kind() == reflect.String ||
		(typ.Kind() == reflect.Slice && typ.Elem().Kind() == reflect.String) ||
		(typ.Kind() == reflect.Array && typ.Elem().Kind() == reflect.String)
}

// Pass method
func (f *fixedFunc) Pass(value reflect.Value) bool {
	if f.fixed == "" {
		return true
	}
	if value.Type().Kind() == reflect.String {
		return f.fixed == value.String()
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
