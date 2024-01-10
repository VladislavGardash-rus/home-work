//nolint:all
package object_field_validators

import (
	"errors"
	"reflect"
)

type ObjectValidatorService struct {
	objectValueInfo reflect.Value
	objectTypeInfo  reflect.Type
}

func NewObjectValidatorService(object interface{}) (*ObjectValidatorService, error) {
	objectValueInfo := reflect.Indirect(reflect.ValueOf(object))
	if objectValueInfo.Kind() != reflect.Struct {
		return nil, errors.New("input values type is not struct")
	}

	objectValidatorService := new(ObjectValidatorService)
	objectValidatorService.objectValueInfo = objectValueInfo
	objectValidatorService.objectTypeInfo = reflect.TypeOf(object)

	return objectValidatorService, nil
}

func (service *ObjectValidatorService) ValidateField(number int) error {
	fieldValidator, err := NewValidatorFactory().GetFieldValidator(service.objectTypeInfo.Field(number).Type.Kind())
	if err != nil {
		return nil
	}

	err = fieldValidator.Validate(service.GetFieldValueInfo(number), service.GetFieldTypeInfo(number))
	if err != nil {
		return err
	}

	return nil
}

func (service *ObjectValidatorService) GetFieldName(number int) string {
	return service.objectTypeInfo.Field(number).Name
}

func (service *ObjectValidatorService) GetFieldValueInfo(number int) reflect.Value {
	return service.objectValueInfo.Field(number)
}

func (service *ObjectValidatorService) GetFieldTypeInfo(number int) reflect.StructField {
	return service.objectTypeInfo.Field(number)
}

func (service *ObjectValidatorService) GetFieldCount() int {
	return service.objectValueInfo.NumField()
}
