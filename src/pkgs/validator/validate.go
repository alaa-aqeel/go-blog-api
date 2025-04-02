package validator

import (
	"encoding/json"
	"io"
)

func (v *Validator) Struct(objStruct any) Errors {
	if err := v.validate.Struct(objStruct); err != nil {

		return v.ParseValidationErrors(err)
	}
	return make(Errors)
}

func MakeValidate[T any](body io.ReadCloser) (*T, Errors) {
	var target T
	err := json.NewDecoder(body).Decode(&target)
	if err != nil {
		return nil, Errors{
			"error": {"invalid body"},
		}
	}

	return &target, Validate.Struct(target)
}
