package validator

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"go.uber.org/zap"
	"hotel-booking-api/logger"
	"log"
)

func (cv *StructValidator) CustomValidate() {

	if err := enTranslations.RegisterDefaultTranslations(cv.Validator, cv.Trans); err != nil {
		log.Fatal(err)
	}

	err := cv.Validator.RegisterTranslation("required", cv.Trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} không được để trống", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
	if err != nil {
		logger.Error("Error register translation", zap.Error(err))
	}

	err = cv.Validator.RegisterValidation("max50", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) <= 50
	})
	if err != nil {
		logger.Error("Error register translation", zap.Error(err))
	}
	err = cv.Validator.RegisterTranslation("max50", cv.Trans, func(ut ut.Translator) error {
		return ut.Add("max50", "{0} nhỏ hơn 50 ký tự", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max50", fe.Field())
		return t
	})
	if err != nil {
		logger.Error("Error register translation", zap.Error(err))
	}
}
