package middleware

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateStruct(s any) []ValidationError {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	var validationErrors []ValidationError
	for _, e := range err.(validator.ValidationErrors) {
		validationErrors = append(validationErrors, ValidationError{
			Field:   strings.ToLower(e.Field()),
			Message: formatMessage(e),
		})
	}

	return validationErrors
}

func formatMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", strings.ToLower(e.Field()))
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", strings.ToLower(e.Field()), e.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", strings.ToLower(e.Field()), e.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", strings.ToLower(e.Field()))
	default:
		return fmt.Sprintf("%s is invalid", strings.ToLower(e.Field()))
	}
}
