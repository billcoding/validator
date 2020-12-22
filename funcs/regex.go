package funcs

import (
	"reflect"
	"regexp"
)

//Define regexFunc struct
type regexFunc struct {
	regex string
}

//RegexFunc
func RegexFunc(regex string) VFunc {
	return &regexFunc{regex}
}

//Accept
func (f *regexFunc) Accept(typ reflect.Type) bool {
	return typ.Kind() == reflect.String ||
		(typ.Kind() == reflect.Slice && typ.Elem().Kind() == reflect.String) ||
		(typ.Kind() == reflect.Array && typ.Elem().Kind() == reflect.String)
}

//Pass
func (f *regexFunc) Pass(value reflect.Value) bool {
	if f.regex == "" {
		return true
	}
	re := regexp.MustCompile(f.regex)
	if value.Type().Kind() == reflect.String {
		return re.MatchString(value.String())
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
