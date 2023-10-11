package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestValidateString(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{input: "a4bc2d5e", expected: true},
		{input: "abccd", expected: true},
		{input: "", expected: true},
		{input: "3abc", expected: false},
		{input: "45", expected: false},
		{input: "aaa10b", expected: false},
		{input: "abccd00abccd", expected: false},
		{input: "abccd03abccd", expected: false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			isValid := ValidateString(tc.input)
			require.Equal(t, tc.expected, isValid)
		})
	}
}

func TestCutZeroSymbolsFromString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a0a0bas2d0", expected: "bas2"},
		{input: "aa0bas2d0", expected: "abas2"},
		{input: "a0a0ba0s2d0", expected: "bs2"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			isValid := CutZeroSymbolsFromString(tc.input)
			require.Equal(t, tc.expected, isValid)
		})
	}
}
