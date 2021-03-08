package funcs

import (
	"fmt"
	"reflect"
	"strings"
)

// enumsFunc struct
type enumsFunc struct {
	enums string
}

// EnumsFunc method
func EnumsFunc(enums string) VFunc {
	return &enumsFunc{enums}
}

// Accept method
func (f *enumsFunc) Accept(typ reflect.Type) bool {
	_, have := numberKindMap[typ.Kind()]
	if have {
		return true
	}
	if typ.Kind() == reflect.Slice || typ.Kind() == reflect.Array {
		_, have = numberKindMap[typ.Elem().Kind()]
		if have {
			return true
		}
	}
	return typ.Kind() == reflect.String ||
		(typ.Kind() == reflect.Slice && typ.Elem().Kind() == reflect.String) ||
		(typ.Kind() == reflect.Array && typ.Elem().Kind() == reflect.String)
}

// Pass method
func (f *enumsFunc) Pass(value reflect.Value) bool {
	if f.enums == "" {
		return true
	}
	set := make(map[string]struct{}, 0)
	for _, enum := range strings.Split(f.enums, ",") {
		set[enum] = struct{}{}
	}
	switch value.Type().Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
		_, have := set[fmt.Sprintf("%d", value.Int())]
		return have
	case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
		_, have := set[fmt.Sprintf("%d", value.Uint())]
		return have
	case reflect.Float32, reflect.Float64:
		_, have := set[fmt.Sprintf("%f", value.Float())]
		return have
	case reflect.String:
		_, have := set[value.String()]
		return have
	case reflect.Array, reflect.Slice:
		_, have := numberKindMap[value.Type().Elem().Kind()]
		if value.Type().Elem().Kind() == reflect.String || have {
			for i := 0; i < value.Len(); i++ {
				if !f.Pass(value.Index(i)) {
					return false
				}
			}
		}
	}
	return true
}
