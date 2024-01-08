package validators

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	lenValidatorName           = "len"
	regexpValidatorName        = "regexp"
	inStringArrayValidatorName = "in"
)

type StringValidator struct{}

func NewStringValidator() *StringValidator {
	return new(StringValidator)
}

func (validator *StringValidator) Validate(objectValueInfo reflect.Value, objectTypeInfo reflect.StructField) error {
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
		case lenValidatorName:
			err = validator.len(objectValueInfo.String(), validatorStrings[1])
			if err != nil {
				stringsBuilder.WriteString(fmt.Sprintf(" %s;", err.Error()))
			}
			break
		case regexpValidatorName:
			err = validator.regexp(objectValueInfo.String(), validatorStrings[1])
			if err != nil {
				stringsBuilder.WriteString(fmt.Sprintf(" %s;", err.Error()))
			}
			break
		case inStringArrayValidatorName:
			err = validator.in(objectValueInfo.String(), validatorStrings[1])
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

func (validator *StringValidator) len(value string, tagValue string) error {
	length, err := strconv.Atoi(tagValue)
	if err != nil {
		return nil
	}

	if len(value) != length {
		return errors.New("incorrect value length")
	}

	return nil
}

func (validator *StringValidator) regexp(value string, tagValue string) error {
	reg, err := regexp.Compile(tagValue)
	if err != nil {
		return nil
	}

	if !reg.MatchString(value) {
		return errors.New("value don't match template")
	}

	return nil
}

func (validator *StringValidator) in(value string, tagValue string) error {
	tagValues := strings.Split(tagValue, ",")
	for _, tagValue := range tagValues {
		if value == tagValue {
			return nil
		}
	}

	return errors.New("value not in " + tagValue)
}
