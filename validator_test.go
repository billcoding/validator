package validator

import (
	"fmt"
	"testing"
)

type model4 struct {
	ID []byte `validate:"required(T) enums(1,2,3) message(ID不合法)"`
	//Name []string `validate:"required(T) enums(a,b,c) message(Name不合法)"`
}

type model3 struct {
	ID     int    `validate:"required(T) min(30) message(model3.Id为空)"`
	Name   string `validate:"required(T) minlength(23) message(model3.Name为空)"`
	Model4 model4 `validate:"required(T)"`
}

type model2 struct {
	ID     int    `validate:"required(T) min(11) message(model2.Id为空)"`
	Name   string `validate:"required(T) minlength(20) message(model2.Name为空)"`
	Model3 model3 `validate:"required(T)"`
}

type model struct {
	ID     int    `validate:"required(T) min(10) message(Id为空)"`
	Name   string `validate:"required(T) minlength(10) message(Name为空)"`
	Model2 model2 `validate:"required(T)"`
}

func TestValidator(t *testing.T) {
	m := &model{
		ID:   10,
		Name: "122222222222",
	}
	validator := New(m)
	result := validator.Validate()
	fmt.Println(result.Passed)
	if !result.Passed {
		fmt.Println(result.Messages())
	}
}

func TestValidator2(t *testing.T) {
	m := &model4{
		ID: []byte{11},
		//Name: []string{"cxx"},
	}
	validator := New(m)
	result := validator.Validate()
	fmt.Println(result.Passed)
	if !result.Passed {
		fmt.Println(result.Messages())
	}
}
