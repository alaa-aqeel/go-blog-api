package helpers

import "github.com/go-playground/validator/v10"

func ParseValidationErrors(err error) Map {
	errorsMap := make(Map)
	for _, vErr := range err.(validator.ValidationErrors) {

		errorsMap.Add(vErr.Field(), errorMessage(vErr))
	}
	return errorsMap
}

func errorMessage(err validator.FieldError) string {

	return err.Tag()
}
