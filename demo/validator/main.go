package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name      string     `validate:"required""`
	Age       uint8      `validate:"gte=0,lte=130"`
	Email     string     `validate:"required,email"`
	Phone     int64      `validate:"number"`
	Addresses []*Address `validate:"required,dive,required"`
}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Phone  string `validate:"required"`
}

// 用 Validate 的单个实例来缓存结构体信息
var validate *validator.Validate

func main() {

	// 创建一个实例

	validate = validator.New()

	address := &Address{
		Street: "Guangzhou 200",
		City:   "Shanghai",
		Phone:  "88888888",
	}

	user := &User{
		Name:      "LiSan",
		Age:       13,
		Email:     "LiSan@world.com",
		Phone:     188 - 888888888,
		Addresses: []*Address{address},
	}

	// 验证结构体
	validateStruct(user)

	// 验证单一变量
	validateVariable()
}

// 复杂结构体的验证
func validateStruct(user *User) {
	err := validate.Struct(user)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
		return
	}

}

func validateVariable() {
	myEmail := "test88888@gmail.com"
	errs := validate.Var(myEmail, "required,email")
	if errs != nil {
		fmt.Println(errs)
		return
	}
}
