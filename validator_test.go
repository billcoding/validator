package validator

import (
	"fmt"
	"testing"
)

type model4 struct {
	ID   int    `validate:"required(T) min(20) message(model4.Id为空)"`
	Name string `validate:"required(T) minlength(22) message(model4.Name为空)"`
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
		Model2: model2{
			ID:   1110,
			Name: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			Model3: model3{
				ID:   850,
				Name: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
				Model4: model4{
					ID:   1098,
					Name: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
				},
			},
		},
	}
	validator := New(m)
	result := validator.Validate()
	fmt.Println(result.Passed)
	if !result.Passed {
		fmt.Println(result.Messages())
	}
}
