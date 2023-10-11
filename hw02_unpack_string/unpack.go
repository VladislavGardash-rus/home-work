package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputString string) (string, error) {
	if len(inputString) == 0 {
		return "", nil
	}

	if !ValidateString(inputString) {
		return "", ErrInvalidString
	}

	resultString := RepeatSymbols(inputString)

	return CutZeroSymbolsFromString(resultString), nil
}

func RepeatSymbols(inputString string) string {
	resultString := ""
	inputStringSymbols := strings.Split(inputString, "")
	for i := range inputStringSymbols {
		number, isNumber := SymbolIsNumber(inputStringSymbols[i])
		if isNumber && number != 0 {
			resultString += strings.Repeat(inputStringSymbols[i-1], number-1)
			continue
		}

		resultString += inputStringSymbols[i]
	}

	return resultString
}

func CutZeroSymbolsFromString(str string) string {
	if strings.Contains(str, "0") {
		zeroIndex := strings.Index(str, "0")
		b := strings.Builder{}
		b.WriteString(str[:zeroIndex-1] + str[zeroIndex+1:])
		str = CutZeroSymbolsFromString(b.String())
	}

	return str
}

func SymbolIsNumber(symbol string) (int, bool) {
	number, err := strconv.Atoi(symbol)
	if err != nil {
		return 0, false
	}

	return number, true
}

func ValidateString(inputString string) bool {
	if strings.Contains(inputString, "00") {
		return false
	}

	for i := 0; i < 10; i++ {
		if strings.Contains(inputString, fmt.Sprintf("0%d", i)) {
			return false
		}

		if strings.HasPrefix(inputString, strconv.Itoa(i)) {
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
