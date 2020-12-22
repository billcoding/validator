package validator

import (
	"reflect"
	"strings"
)

//Define Result struct
type Result struct {
	StructPtr interface{}
	Passed    bool
	Items     []*ResultItem
}

//Define ResultItem struct
type ResultItem struct {
	Field   *reflect.StructField
	Passed  bool
	Message string
}

//Messages
func (r *Result) Messages() string {
	msgs := make([]string, 0)
	for _, item := range r.Items {
		if !item.Passed && item.Message != "" {
			msgs = append(msgs, item.Message)
		}
	}
	return strings.Join(msgs, ",")
}
