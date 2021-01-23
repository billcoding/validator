package validator

import (
	"encoding/json"
	"github.com/billcoding/validator/funcs"
	"reflect"
	"time"
)

// Item struct
type Item struct {
	Required  bool      `alias:"required"`
	Min       float64   `alias:"min"`
	Max       float64   `alias:"max"`
	MinLength int       `alias:"minlength"`
	MaxLength int       `alias:"maxlength"`
	Length    int       `alias:"length"`
	Fixed     string    `alias:"fixed"`
	Enums     []string  `alias:"enums"`
	Regex     string    `alias:"regex"`
	Before    time.Time `alias:"before"`
	After     time.Time `alias:"after"`
	Message   string    `alias:"message"`
}

// vfuncs return VFunc list
func (i *Item) vfuncs() []funcs.VFunc {
	return []funcs.VFunc{
		funcs.RequiredFunc(),
		funcs.MinFunc(i.Min),
		funcs.MaxFunc(i.Max),
		funcs.LengthFunc(i.Length),
		funcs.FixedFunc(i.Fixed),
		funcs.EnumsFunc(i.Enums...),
		funcs.MinLengthFunc(i.MinLength),
		funcs.MaxLengthFunc(i.MaxLength),
		funcs.RegexFunc(i.Regex),
		funcs.BeforeFunc(i.Before),
		funcs.AfterFunc(i.After),
	}
}

// Validate by fields
func (i *Item) Validate(field *reflect.StructField, value *reflect.Value) (bool, string) {
	if !i.Required {
		return true, i.Message
	}
	passed := true
	vfuncs := i.vfuncs()
	for _, vFunc := range vfuncs {
		if !vFunc.Accept(field.Type) {
			continue
		}

		passed = value != nil
		if !passed {
			break
		}

		passed = vFunc.Pass(*value)
		if !passed {
			break
		}
	}
	return passed, i.Message
}

// String return json
func (i *Item) String() string {
	bytes, _ := json.Marshal(i)
	return string(bytes)
}
