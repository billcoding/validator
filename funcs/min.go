package funcs

import (
	"reflect"
)

//Define minFunc struct
type minFunc struct {
	min float64
}

//MinFunc
func MinFunc(min float64) VFunc {
	return &minFunc{min}
}

//Accept
func (f *minFunc) Accept(typ reflect.Type) bool {
	var acceptKinds = map[reflect.Kind]byte{
		// int types
		reflect.Int8:  0,
		reflect.Int16: 0,
		reflect.Int32: 0,
		reflect.Int:   0,
		reflect.Int64: 0,

		//uint types
		reflect.Uint8:  0,
		reflect.Uint16: 0,
		reflect.Uint32: 0,
		reflect.Uint:   0,
		reflect.Uint64: 0,

		//float types
		reflect.Float32: 0,
		reflect.Float64: 0,
	}
	_, have := acceptKinds[typ.Kind()]
	if !have {
		if typ.Kind() == reflect.Slice || typ.Kind() == reflect.Array {
			_, have = acceptKinds[typ.Elem().Kind()]
		}
	}
	return have
}

//Pass
func (f *minFunc) Pass(value reflect.Value) bool {
	if f.min <= 0 {
		return true
	}
	switch value.Type().Kind() {
	default:
	case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
		return f.min <= float64(value.Int())
	case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
		return f.min <= float64(value.Uint())
	case reflect.Float32, reflect.Float64:
		return f.min <= value.Float()
	case reflect.Array, reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			passed := f.Pass(value.Index(i))
			if !passed {
				return false
			}
		}
	}
	return true
}
