package validator

import "github.com/go-playground/validator/v10"

func (v *Validator) ParseValidationErrors(err error) Errors {
	errorsMap := make(Errors)
	for _, vErr := range err.(validator.ValidationErrors) {

		errorsMap.Add(vErr.Field(), v.errorMessage(vErr))
	}
	return errorsMap
}

func (v *Validator) errorMessage(err validator.FieldError) string {

	return err.Tag()
}
