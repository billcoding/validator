package funcs

import (
	"github.com/billcoding/reflectx"
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

// Pass handler
func (f *minFunc) Pass(value reflect.Value) bool {
	if f.min <= 0 {
		return true
	}
	switch {
	case reflectx.IsUint(value.Type()):
		return f.min <= float64(value.Uint())
	case reflectx.IsInt(value.Type()):
		return f.min <= float64(value.Int())
	case reflectx.IsFloat(value.Type()):
		return f.min <= value.Float()
	case reflectx.IsSlice(value.Type()) || reflectx.IsArray(value.Type()):
		for i := 0; i < value.Len(); i++ {
			passed := f.Pass(value.Index(i))
			if !passed {
				return false
			}
		}
	}
	return true
}
