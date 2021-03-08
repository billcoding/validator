package funcs

import "reflect"

var numberKindMap = map[reflect.Kind]struct{}{
	// int types
	reflect.Int8:  {},
	reflect.Int16: {},
	reflect.Int32: {},
	reflect.Int:   {},
	reflect.Int64: {},

	// uint types
	reflect.Uint8:  {},
	reflect.Uint16: {},
	reflect.Uint32: {},
	reflect.Uint:   {},
	reflect.Uint64: {},

	// float types
	reflect.Float32: {},
	reflect.Float64: {},
}
