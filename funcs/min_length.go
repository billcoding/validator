package funcs

import (
	"github.com/billcoding/reflectx"
	"reflect"
)

// minLengthFunc struct
type minLengthFunc struct {
	minLength int
}

// MinLengthFunc method
func MinLengthFunc(minLength int) VFunc {
	return &minLengthFunc{minLength}
}

// Accept method
func (f *minLengthFunc) Accept(typ reflect.Type) bool {
	return reflectx.IsString(typ) ||
		((reflectx.IsArray(typ) || reflectx.IsSlice(typ)) && reflectx.IsString(typ.Elem())) ||
		((reflectx.IsArray(typ) || reflectx.IsSlice(typ)) && reflectx.IsStruct(typ.Elem())) ||
		((reflectx.IsArray(typ) || reflectx.IsSlice(typ)) && reflectx.IsPtr(typ.Elem()) && reflectx.IsStruct(typ.Elem().Elem()))
}

// Pass method
func (f *minLengthFunc) Pass(value reflect.Value) bool {
	if f.minLength <= 0 {
		return true
	}
	typ := value.Type()
	switch {
	case reflectx.IsString(typ):
		return f.minLength <= len(value.String())
	case (reflectx.IsArray(typ) || reflectx.IsSlice(typ)) && reflectx.IsString(typ.Elem()):
		for i := 0; i < value.Len(); i++ {
			if !f.Pass(value.Index(i)) {
				return false
			}
		}
	case (reflectx.IsArray(typ) || reflectx.IsSlice(typ)) && reflectx.IsStruct(typ.Elem()):
		return f.minLength <= value.Len()
	case (reflectx.IsArray(typ) || reflectx.IsSlice(typ)) && reflectx.IsPtr(typ.Elem()) && reflectx.IsStruct(typ.Elem().Elem()):
		return f.minLength <= value.Len()
	}
	return true
}
