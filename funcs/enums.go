package funcs

import "reflect"

//Define enumsFunc struct
type enumsFunc struct {
	enums []string
}

//EnumsFunc
func EnumsFunc(enums ...string) VFunc {
	return &enumsFunc{enums}
}

//Accept
func (f *enumsFunc) Accept(typ reflect.Type) bool {
	return typ.Kind() == reflect.String ||
		(typ.Kind() == reflect.Slice && typ.Elem().Kind() == reflect.String) ||
		(typ.Kind() == reflect.Array && typ.Elem().Kind() == reflect.String)
}

//Pass
func (f *enumsFunc) Pass(value reflect.Value) bool {
	if len(f.enums) <= 0 {
		return true
	}
	set := make(map[string]byte, 0)
	for _, enum := range f.enums {
		set[enum] = 0
	}
	if value.Type().Kind() == reflect.String {
		_, have := set[value.String()]
		return have
	} else if value.Type().Kind() == reflect.Slice && value.Type().Elem().Kind() == reflect.String ||
		(value.Type().Kind() == reflect.Array && value.Type().Elem().Kind() == reflect.String) {
		for i := 0; i < value.Len(); i++ {
			if !f.Pass(value.Index(i)) {
				return false
			}
		}
	}
	return true
}
