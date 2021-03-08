package funcs

import (
	"reflect"
)

// minFunc struct
type minFunc struct {
	min float64
}

// MinFunc method
func MinFunc(min float64) VFunc {
	return &minFunc{min}
}

// Accept method
func (f *minFunc) Accept(typ reflect.Type) bool {
	_, have := numberKindMap[typ.Kind()]
	if !have {
		if typ.Kind() == reflect.Slice || typ.Kind() == reflect.Array {
			_, have = numberKindMap[typ.Elem().Kind()]
		}
	}
	return have
}

// Pass
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
