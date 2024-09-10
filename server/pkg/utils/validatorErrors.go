package utils

import "github.com/go-playground/validator/v10"

func ValidatorErrors(err validator.ValidationErrors) map[string]string {
	errors := make(map[string]string)
	for _, fieldErr := range err {
		errors[fieldErr.Field()] = fieldErr.Tag()
	}
	return errors
}
