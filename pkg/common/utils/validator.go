package utils

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni              *ut.UniversalTranslator
	translator       ut.Translator
	validateInstance *validator.Validate
)

func init() {

	// NOTE: ommitting allot of error checking for brevity

	en := en.New()
	uni = ut.New(en, en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	translator, _ = uni.GetTranslator("en")

	validateInstance = validator.New()
	en_translations.RegisterDefaultTranslations(validateInstance, translator)

	translateOverride(translator)
}

func translateOverride(trans ut.Translator) {

	validateInstance.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	validateInstance.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} must be a valid email address!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())

		return t
	})

	// validateInstance.RegisterTranslation("len", trans, func(ut ut.Translator) error {
	// 	return ut.Add("email", "{0} must be a valid email address!", true)
	// }, func(ut ut.Translator, fe validator.FieldError) string {
	// 	t, _ := ut.T("email", fe.Field())

	// 	return t
	// })

}

func Validate(value interface{}) (bool, *validator.ValidationErrors) {

	err := validateInstance.Struct(value)

	if err != nil {

		errs := err.(validator.ValidationErrors)

		fmt.Println(errs.Translate(translator))

		return false, &errs
	}

	return true, nil
}
