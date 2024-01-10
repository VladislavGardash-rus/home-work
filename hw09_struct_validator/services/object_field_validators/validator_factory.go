//nolint:all
package object_field_validators

import (
	"errors"
	"github.com/gardashvs/home-work/hw09_struct_validator/services/object_field_validators/validators"
	"reflect"
)

type FieldValidator interface {
	Validate(reflect.Value, reflect.StructField) error
}

type ValidatorFactory struct {
	Validators map[reflect.Kind]FieldValidator
}

func NewValidatorFactory() *ValidatorFactory {
	validatorFactory := new(ValidatorFactory)
	validatorFactory.Validators = make(map[reflect.Kind]FieldValidator)
	validatorFactory.Validators[reflect.Int] = validators.NewIntValidator()
	validatorFactory.Validators[reflect.String] = validators.NewStringValidator()
	validatorFactory.Validators[reflect.Slice] = validators.NewSliceValidator()
	return validatorFactory
}

func (validatorFactory *ValidatorFactory) GetFieldValidator(kind reflect.Kind) (FieldValidator, error) {
	switch kind {
	case reflect.Int:
		validator := validators.NewIntValidator()
		return validator, nil
	case reflect.String:
		validator := validators.NewStringValidator()
		return validator, nil
	case reflect.Slice:
		validator := validators.NewSliceValidator()
		return validator, nil
	default:
		return nil, errors.New("field validator not found")
	}
}
