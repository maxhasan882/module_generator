package string_formatter

import "testing"

func TestFormatFloat(t *testing.T) {
	testCases := []struct {
		input    float64
		expected string
	}{
		{123.456, "123.46"},
		{123.4500, "123.45"},
		{0.001, "0"},
		{5.001, "5"},
		{12.0, "12"},
		{123456.789, "123456.79"},
		{123456.0044, "123456"},
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			result := FormatFloat(testCase.input)
			if result != testCase.expected {
				t.Errorf("For input %.3f, expected: %s, Got: %s", testCase.input, testCase.expected, result)
			}
		})
	}
}
