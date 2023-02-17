package utils

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("not-blank", notBlank)
}

func notBlank(fl validator.FieldLevel) bool {
	return strings.Trim(fl.Field().String(), " ") != ""
}

// 验证对象，返回第一个发现的错误
func GetFirestErr(o any) (err error) {

	if o == nil {
		return errors.New("validate object is nil")
	}

	err = validate.Struct(o)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, err = range err.(validator.ValidationErrors) {

			return err
			// fmt.Println("namespace:", err.Namespace())
			// fmt.Println("field:", err.Field())
			// fmt.Println("structnamespace:", err.StructNamespace())
			// fmt.Println("structfield:", err.StructField())
			// fmt.Println("tag:", err.Tag())
			// fmt.Println("actualtag:", err.ActualTag())
			// fmt.Println("kind:", err.Kind())
			// fmt.Println("type:", err.Type())
			// fmt.Println("value:", err.Value())
			// fmt.Println("param:", err.Param())
			// fmt.Println("err", err.Error())
		}

	}
	return nil
}
