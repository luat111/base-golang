package middlewares

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

type ErrorField struct {
	ErrorMessage string `json:"errorMessage"`
	FieldName    string `json:"fieldName"`
}

func ValidateStruct(payload interface{}) (errs []ErrorField) {
	err := Validator.Struct(payload)

	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, e := range err.(validator.ValidationErrors) {
				errs = append(errs, ErrorField{
					FieldName:    e.Field(),
					ErrorMessage: e.Tag(),
				})
			}
			return errs
		}
	}
	return nil
}
