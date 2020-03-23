package validator

import "github.com/go-playground/validator"

type (
	Validator struct {
		inner *validator.Validate
	}
)

func (cv *Validator) Validate(i interface{}) error {
	return cv.inner.Struct(i)
}

func NewValidator() *Validator {
	return &Validator{
		inner: validator.New(),
	}
}
