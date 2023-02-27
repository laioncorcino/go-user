package service

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en2 "github.com/go-playground/validator/translations/en"
	"github.com/go-playground/validator/v10"
	"github.com/laioncorcino/go-user/rest_error"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		langEN := en.New()
		universalTranslator := ut.New(langEN, langEN)
		transl, _ = universalTranslator.GetTranslator("en")
		_ = en2.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validationError error) *rest_error.StandardError {
	var jsonError *json.UnmarshalTypeError
	var jsonValidatorError validator.ValidationErrors

	if errors.As(validationError, &jsonError) {
		return rest_error.NewBadRequestError("Invalid fields type")

	} else if errors.As(validationError, &jsonValidatorError) {
		var errorsCauses []rest_error.Cause

		for _, e := range validationError.(validator.ValidationErrors) {
			cause := rest_error.Cause{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return rest_error.NewValidationError("Some fields are invalid", errorsCauses)

	} else {
		return rest_error.NewBadRequestError("Error trying to convert fields")
	}
}
