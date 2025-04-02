package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Validator struct {
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	Trans    ut.Translator
}

var Validate *Validator

func InitlizeValidator() *Validator {
	Validate = &Validator{}
	Validate.newValidator()
	Validate.newTranslator()

	return Validate
}

func (v *Validator) newValidator() {
	v.validate = validator.New(func(v *validator.Validate) {
		v.SetTagName("binding")
	})
	v.validate.RegisterTagNameFunc(func(fld reflect.StructField) string {

		return strings.Split(fld.Tag.Get("json"), ",")[0]
	})
}

func (v *Validator) newTranslator() {
	en := en.New()
	v.uni = ut.New(en, en)
	trans, _ := v.uni.GetTranslator("en")
	v.Trans = trans

	en_translations.RegisterDefaultTranslations(v.validate, trans)
}
