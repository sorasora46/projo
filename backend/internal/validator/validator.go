package vldt

import (
	"github.com/go-playground/validator/v10"
	"github.com/sorasora46/projo/backend/internal/dtos"
	"github.com/sorasora46/projo/backend/pkg/utils"
)

type ReqValidator struct {
	validate *validator.Validate
}

func NewReqValidator(validate *validator.Validate) ReqValidator {
	return ReqValidator{validate: validate}
}

func (v *ReqValidator) Validate(data any) []dtos.ValidationError {
	errs := v.validate.Struct(data)
	if errs != nil {
		errors := []dtos.ValidationError{}
		for _, err := range errs.(validator.ValidationErrors) {
			field := err.Field()
			utils.PascalToCamelCase(&field)
			validationErr := dtos.ValidationError{
				Field:         field,
				Message:       utils.GenValidationErrAtTag(err.Tag()),
				ValueProvided: err.Value(),
			}
			errors = append(errors, validationErr)
		}
		return errors
	}
	return nil
}
