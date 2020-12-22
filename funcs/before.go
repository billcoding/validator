package funcs

import (
	"github.com/billcoding/reflectx"
	"reflect"
	"time"
)

//Define beforeFunc struct
type beforeFunc struct {
	before time.Time
}

//BeforeFunc
func BeforeFunc(before time.Time) VFunc {
	return &beforeFunc{before}
}

//Accept
func (f *beforeFunc) Accept(typ reflect.Type) bool {
	return typ.Kind() == reflect.String || typ == reflect.TypeOf(time.Time{})
}

//Pass
func (f *beforeFunc) Pass(value reflect.Value) bool {
	if f.before.IsZero() {
		return true
	}
	if value.Type().Kind() == reflect.String {
		return f.Pass(reflect.ValueOf(reflectx.ParseTime(value.String())))
	} else if value.Type() == reflect.TypeOf(time.Time{}) {
		return f.before.After(value.Interface().(time.Time))
	}
	return true
}
