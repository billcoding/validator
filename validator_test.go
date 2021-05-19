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

type UserDeptPermPost struct {
	Dept []*UserDeptPerm `json:"dept" validate:"required(T) message(dept不能为空)"`
}

// UserDeptPerm struct 用户部门权限
type UserDeptPerm struct {
	// DeptID 部门ID
	DeptID string `db:"dept_id" json:"dept_id" generator:"DB_PRI" validate:"required(T) message(dept_id验证未通过) minlength(1)"`

	// DeptName 部门名称
	DeptName string `db:"dept_name" json:"dept_name" validate:"required(T) message(dept_name验证未通过) minlength(1)"`
}

func TestValidator2(t *testing.T) {
	m := UserDeptPermPost{
		Dept: []*UserDeptPerm{{}, {}, {}},
	}
	validator := New(&m)
	result := validator.Validate()
	fmt.Println(validator.fields)
	fmt.Println(validator.values)
	fmt.Println(m)
	fmt.Println(result.Passed)
	if !result.Passed {
		fmt.Println(result.Messages())
	}
}
