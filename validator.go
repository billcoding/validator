package validator

import (
	"github.com/billcoding/reflectx"
	"reflect"
)

// Validator
type Validator struct {
	structPtr interface{}
	items     []*Item
	fields    []*reflect.StructField
}

// New
func New(structPtr interface{}) *Validator {
	items := make([]*Item, 0)
	fields := reflectx.CreateFromTag(structPtr, &items, "alias", "validate")
	if len(items) != len(fields) {
		panic("[New]invalid len both items and fields")
	}
	return &Validator{
		structPtr: structPtr,
		items:     items,
		fields:    fields,
	}
}

// Validate
func (v *Validator) Validate() *Result {
	ritems := make([]*ResultItem, len(v.items), len(v.items))
	passedCount := 0
	for pos, item := range v.items {
		field := v.fields[pos]
		value := reflect.ValueOf(v.structPtr).Elem().FieldByName(field.Name)
		resultItem := validate(item, field, value)
		ritems[pos] = resultItem
		if resultItem.Passed {
			passedCount++
		}
	}
	return &Result{
		StructPtr: v.structPtr,
		Passed:    len(v.items) == passedCount,
		Items:     ritems,
	}
}

func validate(item *Item, field *reflect.StructField, value reflect.Value) *ResultItem {
	var innerInterface interface{}
	switch {
	case field.Type.Kind() == reflect.Struct:
		innerInterface = value.Addr().Interface()
	case field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct:
		innerInterface = value.Elem().Addr().Interface()
	default:
		passed, msg := item.Validate(field, value)
		return &ResultItem{Field: field, Passed: passed, Message: msg}
	}
	if !item.Required || innerInterface == nil {
		return &ResultItem{Field: field, Passed: true}
	}
	items := make([]*Item, 0)
	fields := reflectx.CreateFromTag(innerInterface, &items, "alias", "validate")
	if len(items) != len(fields) {
		panic("[New]invalid len both items and fields")
	}
	v := &Validator{
		structPtr: innerInterface,
		items:     items,
		fields:    fields,
	}
	vresult := v.Validate()
	return &ResultItem{
		Field:   field,
		Passed:  vresult.Passed,
		Message: vresult.Messages(),
	}
}
