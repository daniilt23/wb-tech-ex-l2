package main

import (
	"errors"
	"testing"

	"github.com/go-playground/assert/v2"
)

type expected struct {
	result string
	err    error
}

func TestStringUnpacking(t *testing.T) {
	testTable := []struct {
		input       string
		exp         expected
		expectedErr bool
	}{
		{
			input: "a4bc2d5e",
			exp: expected{
				result: "aaaabccddddde",
				err:    nil,
			},
			expectedErr: false,
		},
		{
			input: "abcd",
			exp: expected{
				result: "abcd",
				err:    nil,
			},
			expectedErr: false,
		},
		{
			input: "45",
			exp: expected{
				result: "",
				err:    errors.New("строка не может состоять только из цифр"),
			},
			expectedErr: true,
		},
		{
			input: "",
			exp: expected{
				result: "",
				err:    nil,
			},
			expectedErr: false,
		},
		{
			input: "qwe\\4\\5",
			exp: expected{
				result: "qwe45",
				err:    nil,
			},
			expectedErr: false,
		},
		{
			input: "qwe\\45",
			exp: expected{
				result: "qwe44444",
				err:    nil,
			},
			expectedErr: false,
		},
	}

	for _, testCase := range testTable {
		result, err := CheckCondition(testCase.input)
		if testCase.expectedErr {
			assert.Equal(t, testCase.exp.err, err)
		}
		assert.Equal(t, testCase.exp.result, result)
	}
}
