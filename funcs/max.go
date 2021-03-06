package funcs

import (
	"reflect"
)

// minFunc struct
type maxFunc struct {
	max float64
}

// MaxFunc method
func MaxFunc(max float64) VFunc {
	return &maxFunc{max}
}

// Accept method
func (f *maxFunc) Accept(typ reflect.Type) bool {
	_, have := numberKindMap[typ.Kind()]
	if !have {
		if typ.Kind() == reflect.Slice || typ.Kind() == reflect.Array {
			_, have = numberKindMap[typ.Elem().Kind()]
		}
	}
	return have
}

// Pass method
func (f *maxFunc) Pass(value reflect.Value) bool {
	if f.max <= 0 {
		return true
	}
	switch value.Type().Kind() {
	default:
	case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
		return f.max >= float64(value.Int())
	case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
		return f.max >= float64(value.Uint())
	case reflect.Float32, reflect.Float64:
		return f.max >= value.Float()
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
