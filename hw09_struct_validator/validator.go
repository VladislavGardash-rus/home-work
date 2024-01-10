//nolint:all
package hw09structvalidator

import (
	"github.com/gardashvs/home-work/hw09_struct_validator/services/object_field_validators"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (validationErrors ValidationErrors) Error() string {
	stringsBuilder := strings.Builder{}

	stringsBuilder.WriteString("Validation errors:")
	if len(validationErrors) == 0 {
		stringsBuilder.WriteString(" none")
	}

	for _, validationError := range validationErrors {
		stringsBuilder.WriteString(validationError.Err.Error())
	}

	return stringsBuilder.String()
}

func Validate(v interface{}) error {
	if v == nil {
		return nil
	}

	var validationErrors ValidationErrors
	validationErrors = make([]ValidationError, 0)

	objectValidatorService, err := object_field_validators.NewObjectValidatorService(v)
	if err != nil {
		return err
	}

	for i := 0; i < objectValidatorService.GetFieldCount(); i++ {
		err = objectValidatorService.ValidateField(i)
		if err != nil {
			validationErrors = append(validationErrors, ValidationError{objectValidatorService.GetFieldName(i), err})
		}
	}

	return validationErrors
}
