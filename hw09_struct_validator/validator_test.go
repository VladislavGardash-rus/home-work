//nolint:all
package hw09structvalidator

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

var wrongUser = User{
	ID:     "123",
	Age:    123,
	Email:  "123.gmail.com",
	Role:   "executor",
	Phones: []string{"03", "10", "11"},
}

var rightUser = User{
	ID:     "850b0bce-2808-47a2-b3c9-3741ddb948f5",
	Name:   "Bob",
	Age:    22,
	Email:  "qwerty@gmail.com",
	Role:   "admin",
	Phones: []string{"88005555535", "84955555535"},
	meta:   []byte("{}"),
}

var wrongApp = App{
	Version: "0",
}

var rightApp = App{
	Version: "0.1.5",
}

var token = Token{
	Header:    []byte{97},
	Payload:   []byte{98},
	Signature: []byte{99},
}

var wrongResponse = Response{
	Code: 503,
	Body: "body",
}

var rightResponse = Response{
	Code: 200,
	Body: "body",
}

func TestValidateSuccess(t *testing.T) {
	tests := []struct {
		name          string
		in            interface{}
		expectedError error
	}{
		{
			name:          "rightUser",
			in:            rightUser,
			expectedError: errors.New("Validation errors: none"),
		},
		{
			name:          "rightApp",
			in:            rightApp,
			expectedError: errors.New("Validation errors: none"),
		},
		{
			name:          "token",
			in:            token,
			expectedError: errors.New("Validation errors: none"),
		},
		{
			name:          "rightResponse",
			in:            rightResponse,
			expectedError: errors.New("Validation errors: none"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.in)
			require.Equal(t, tt.expectedError.Error(), err.Error())
		})
	}
}

func TestValidateFail(t *testing.T) {
	tests := []struct {
		name          string
		in            interface{}
		expectedError error
	}{
		{
			name:          "wrongUser",
			in:            wrongUser,
			expectedError: errors.New("Validation errors: field ID: incorrect value length; field Age: value more than max; field Email: value don't match template; field Role: value not in admin,stuff; field Phones: value 03: incorrect value length; value 10: incorrect value length; value 11: incorrect value length;"),
		},
		{
			name:          "wrongApp",
			in:            wrongApp,
			expectedError: errors.New("Validation errors: field Version: incorrect value length;"),
		},
		{
			name:          "wrongResponse",
			in:            wrongResponse,
			expectedError: errors.New("Validation errors: field Code: value not in 200,404,500;"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.in)
			require.Equal(t, tt.expectedError.Error(), err.Error())
		})
	}
}
