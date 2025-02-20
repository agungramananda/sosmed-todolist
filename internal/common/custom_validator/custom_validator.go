package custom_validator

import "github.com/go-playground/validator/v10"

type Validator struct {
	validate *validator.Validate
}

func NewCustomValidator(v *validator.Validate) *Validator{
	return &Validator{
		validate: v,
	}
}

func (v *Validator) Validate(i interface{}) error {
	return v.validate.Struct(i)
}
