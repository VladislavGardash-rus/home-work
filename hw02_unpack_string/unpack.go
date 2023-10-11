package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputString string) (string, error) {
	resultString := ""
	inputStringSymbols := strings.Split(inputString, "")
	for i := range inputStringSymbols {
		number, isNumber := SymbolIsNumber(inputStringSymbols[i])
		if isNumber && number != 0 {
			resultString += strings.Repeat(inputStringSymbols[i-1], number-1)
			continue
		} else if isNumber && number == 0 {
			resultString = resultString[:len(resultString)-1]
			continue
		}
		resultString += inputStringSymbols[i]
	}
	return resultString, nil
}

func SymbolIsNumber(symbol string) (int, bool) {
	number, err := strconv.Atoi(symbol)
	if err != nil {
		return 0, false
	}
	return number, true
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
