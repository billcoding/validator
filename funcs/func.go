package funcs

import "reflect"

// VFunc
type VFunc interface {
	Accept(typ reflect.Type) bool
	Pass(value reflect.Value) bool
}
