package tolo

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

const FAILED_VALIDATION = "failed validation"
const FORBIDDEN = "forbidden"
const NOT_FOUND = "not found"

func ValidatorTranslate(data any, errs error) map[string]string {
	validatorErrs, ok := errs.(validator.ValidationErrors)
	if !ok {
		return map[string]string{
			"default": errs.Error(),
		}
	}

	errors := make(map[string]string)
	for _, e := range validatorErrs {
		field, _ := reflect.TypeOf(data).FieldByName(e.StructField())
		name := field.Tag.Get("name")

		switch e.Tag() {
		case "required":
			errors[name] = fmt.Sprintf("%s is required", name)
		case "email":
			errors[name] = "Invalid email format"
		case "min":
			errors[name] = fmt.Sprintf("%s must be at least %s characters", name, e.Param())
		default:
			errors[name] = "Invalid value"
		}
	}
	return errors
}

func Validator() *validator.Validate {
	return validator.New()
}
