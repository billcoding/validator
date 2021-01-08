package funcs

import "reflect"

// VFunc interface
type VFunc interface {
	Accept(typ reflect.Type) bool
	Pass(value reflect.Value) bool
}
