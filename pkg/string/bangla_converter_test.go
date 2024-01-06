package string_formatter

import (
	"fmt"
	"testing"
)

func TestConvertToBanglaNumeric(t *testing.T) {
	testCases := []struct {
		input    int
		expected string
	}{
		{1234567890, "১২৩৪৫৬৭৮৯০"},
		{9876543210, "৯৮৭৬৫৪৩২১০"},
		{0, "০"},
		{-123, "-১২৩"},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Input: %d", testCase.input), func(t *testing.T) {
			result := ConvertToBanglaNumeric(testCase.input)
			if result != testCase.expected {
				t.Errorf("Expected: %s, Got: %s", testCase.expected, result)
			}
		})
	}
}

func TestConvertToBanglaNumericFloatNumber(t *testing.T) {
	testCases := []struct {
		input    float64
		expected string
	}{
		{1.23456789, "১.২৩৪৫৬৭৮৯"},
		{9.87654321, "৯.৮৭৬৫৪৩২১"},
		{0, "০"},
		{-1.23, "-১.২৩"},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Input: %f", testCase.input), func(t *testing.T) {
			result := ConvertToBanglaNumeric(testCase.input)
			if result != testCase.expected {
				t.Errorf("Expected: %s, Got: %s", testCase.expected, result)
			}
		})
	}
}
