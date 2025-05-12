package user

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

var phoneRegex = regexp.MustCompile(`^\+998\d{9}$`)

func NewUserValidator() *Validator {
	v := validator.New()
	v.RegisterValidation("phone", validatePhone)

	return &Validator{validate: v}
}

func validatePhone(fl validator.FieldLevel) bool {
	return phoneRegex.MatchString(fl.Field().String())
}

func (v *Validator) Validate(s any) error {
	return v.validate.Struct(s)
}
