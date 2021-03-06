package validator

import (
	"github.com/billcoding/reflectx"
	"reflect"
)

// Validator struct
type Validator struct {
	structPtr interface{}
	fields    []*reflect.StructField
	values    []*reflect.Value
	items     []interface{}
}

// New return new Validator
func New(structPtr interface{}) *Validator {
	fields, values, tags := reflectx.ParseTag(structPtr, new(Item), "alias", "validate", true)
	return &Validator{
		structPtr: structPtr,
		fields:    fields,
		values:    values,
		items:     tags,
	}
}

// Validate start
func (v *Validator) Validate() *Result {
	resultItems := make([]*ResultItem, len(v.fields))
	passedCount := 0
	for pos := range v.fields {
		field := v.fields[pos]
		value := v.values[pos]
		item := v.items[pos].(*Item)
		resultItem := validate(item, field, value)
		resultItems[pos] = resultItem
		if resultItem.Passed {
			passedCount++
		}
	}
	return &Result{
		StructPtr: v.structPtr,
		Passed:    len(v.items) == passedCount,
		Items:     resultItems,
	}
}

func validate(item *Item, field *reflect.StructField, value *reflect.Value) *ResultItem {
	passed, msg := item.Validate(*field, *value)
	return &ResultItem{Field: field, Passed: passed, Message: msg}
}
