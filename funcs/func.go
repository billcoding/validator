package funcs

import "reflect"

//Define VFunc interface
type VFunc interface {
	Accept(typ reflect.Type) bool
	Pass(value reflect.Value) bool
}
