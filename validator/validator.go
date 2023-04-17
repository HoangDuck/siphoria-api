package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"strings"
)

type StructValidator struct {
	Validator *validator.Validate
	Uni       *ut.UniversalTranslator
	Trans     ut.Translator
}

// NewStructValidator create new validator object
func NewStructValidator() *StructValidator {
	translator := en.New()
	uni := ut.New(translator, translator)
	trans, _ := uni.GetTranslator("en")
	return &StructValidator{
		Validator: validator.New(),
		Uni:       uni,
		Trans:     trans,
	}
}

// Validate setup validate
func (cv *StructValidator) Validate(i interface{}) error {
	//cv.RegisterValidate()
	err := cv.Validator.Struct(i)
	if err == nil {
		return nil
	}

	transErrors := make([]string, 0)
	for _, e := range err.(validator.ValidationErrors) {
		transErrors = append(transErrors, e.Translate(cv.Trans))
	}
	return errors.Errorf("%s", strings.Join(transErrors, " \n "))
}
