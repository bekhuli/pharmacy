package user

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewUserValidator() *Validator {
	v := validator.New()

	return &Validator{validate: v}
}

func (v *Validator) Validate(s any) error {
	return v.validate.Struct(s)
}
