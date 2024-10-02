package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()

	validate.RegisterValidation("uniqueEmail", UniqueEmail)
	validate.RegisterValidation("uniquePhoneNumber", UniquePhoneNumber)
	validate.RegisterValidation("indonesianPhoneNumber", IndonesianPhoneNumber)
}

func ValidateStruct(data interface{}) map[string]string {
	err := validate.Struct(data)
	if err != nil {
		errs := map[string]string{}

		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()
			errs[field] = getErrorMessage(field, tag)
		}

		return errs
	}

	return nil
}

func getErrorMessage(field, tag string) string {
	switch tag {
	case "required":
		return field + " is required"
	case "email":
		return field + " must be a valid email"
	case "uniqueEmail":
		return field + " this email is already taken"
	case "uniquePhoneNumber":
		return field + " this phone number is already taken"
	case "indonesianPhoneNumber":
		return field + " must be a valid Indonesian phone number"
	default:
		return field + " is not valid"
	}
}
