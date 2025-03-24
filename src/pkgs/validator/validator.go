package validator

import "github.com/go-playground/validator/v10"

type Validator struct {
	validate *validator.Validate
}

var Validate *Validator

func InitlizeValidator() *Validator {
	v := validator.New()

	Validate = &Validator{validate: v}
	return Validate
}

func (v *Validator) ValidateStruct(objStruct any) Errors {
	if err := v.validate.Struct(objStruct); err != nil {
		return v.ParseValidationErrors(err)
	}

	return make(Errors)
}
