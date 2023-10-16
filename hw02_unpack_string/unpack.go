package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputString string) (string, error) {
	if !ValidateString(inputString) {
		return "", ErrInvalidString
	}

	resultRunes := make([]rune, 0)
	inputRunes := []rune(inputString)
	for i := range []rune(inputString) {
		number, err := strconv.Atoi(string(inputRunes[i]))
		if err == nil && number != 0 {
			for j := 1; j < number; j++ {
				resultRunes = append(resultRunes, inputRunes[i-1])
			}
			continue
		} else if err == nil && number == 0 {
			resultRunes = resultRunes[:len(resultRunes)-1]
			continue
		}
		resultRunes = append(resultRunes, inputRunes[i])
	}

	return string(resultRunes), nil
}

func ValidateString(inputString string) bool {
	for i := 0; i < 10; i++ {
		if strings.Contains(inputString, fmt.Sprintf("0%d", i)) || strings.HasPrefix(inputString, strconv.Itoa(i)) {
			return false
		}
	}
	for i := 10; i < 100; i++ {
		if strings.Contains(inputString, strconv.Itoa(i)) {
			return false
		}
	}
	return true
}
