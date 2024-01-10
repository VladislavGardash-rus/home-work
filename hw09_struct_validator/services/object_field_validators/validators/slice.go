//nolint:all
package validators

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type SliceValidator struct{}

func NewSliceValidator() *SliceValidator {
	return new(SliceValidator)
}

func (validator *SliceValidator) Validate(objectValueInfo reflect.Value, objectTypeInfo reflect.StructField) error {
	tag := objectTypeInfo.Tag.Get("validate")
	if tag == "" {
		return nil
	}

	var err error
	stringsBuilder := strings.Builder{}
	stringsBuilder.WriteString(fmt.Sprintf(" field %s:", objectTypeInfo.Name))

	validatorStrings := strings.Split(tag, "|")
	for _, validatorString := range validatorStrings {
		validatorStrings := strings.Split(validatorString, ":")
		if len(validatorStrings) != 2 {
			continue
		}

		for i := 0; i < objectValueInfo.Len(); i++ {
			if objectValueInfo.Index(i).Kind() == reflect.String {
				validator := NewStringValidator()
				err = validator.Validate(objectValueInfo.Index(i), objectTypeInfo)
				if err != nil {
					stringsBuilder.WriteString(fmt.Sprintf("%s", strings.ReplaceAll(err.Error(), fmt.Sprintf(" field %s:", objectTypeInfo.Name), fmt.Sprintf(" value %s:", objectValueInfo.Index(i).String()))))
				}
			}

			if objectValueInfo.Index(i).Kind() == reflect.Int {
				validator := NewIntValidator()
				err = validator.Validate(objectValueInfo.Index(i), objectTypeInfo)
				if err != nil {
					stringsBuilder.WriteString(fmt.Sprintf("%s", strings.ReplaceAll(err.Error(), fmt.Sprintf(" field %s:", objectTypeInfo.Name), fmt.Sprintf(" value %d:", objectValueInfo.Index(i).Int()))))
				}
			}
		}
	}

	if err != nil {
		return errors.New(stringsBuilder.String())
	}

	return nil
}
