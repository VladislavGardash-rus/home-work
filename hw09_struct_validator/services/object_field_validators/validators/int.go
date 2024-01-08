package validators

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	minValidatorName = "min"
	maxValidatorName = "max"
	inValidatorName  = "in"
)

type IntValidator struct{}

func NewIntValidator() *IntValidator {
	return new(IntValidator)
}

func (validator *IntValidator) Validate(objectValueInfo reflect.Value, objectTypeInfo reflect.StructField) error {
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

		switch validatorStrings[0] {
		case minValidatorName:
			err = validator.min(objectValueInfo.Int(), validatorStrings[1])
			if err != nil {
				stringsBuilder.WriteString(fmt.Sprintf(" %s;", err.Error()))
			}
			break
		case maxValidatorName:
			err = validator.max(objectValueInfo.Int(), validatorStrings[1])
			if err != nil {
				stringsBuilder.WriteString(fmt.Sprintf(" %s;", err.Error()))
			}
			break
		case inValidatorName:
			err = validator.in(objectValueInfo.Int(), validatorStrings[1])
			if err != nil {
				stringsBuilder.WriteString(fmt.Sprintf(" %s;", err.Error()))
			}
			break
		default:
			continue
		}
	}

	if err != nil {
		return errors.New(stringsBuilder.String())
	}

	return nil
}

func (validator *IntValidator) min(value int64, tagValue string) error {
	intTagValue, err := strconv.Atoi(tagValue)
	if err != nil {
		return nil
	}

	if value < int64(intTagValue) {
		return errors.New("value less than min")
	}

	return nil
}

func (validator *IntValidator) max(value int64, tagValue string) error {
	intTagValue, err := strconv.Atoi(tagValue)
	if err != nil {
		return nil
	}

	if value > int64(intTagValue) {
		return errors.New("value more than max")
	}

	return nil
}

func (validator *IntValidator) in(value int64, tagValue string) error {
	tagValues := strings.Split(tagValue, ",")
	for _, tagValue := range tagValues {
		if strconv.FormatInt(value, 10) == tagValue {
			return nil
		}
	}

	return errors.New("value not in " + tagValue)
}
