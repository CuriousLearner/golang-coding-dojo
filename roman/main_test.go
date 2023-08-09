package main

import (
	"errors"
	"testing"
)

func TestConvertToRoman(t *testing.T) {
	testCases := []struct {
		name        string
		input       int
		expected    string
		expectedErr error
	}{
		{"NegativeNumber", -9, "", errors.New("No representation found. Numbers start with 1 in roman numeral")},
		{"TestCase0", 0, "", errors.New("No representation found. Numbers start with 1 in roman numeral")},
		{"TestCase1", 1, "I", nil},
		{"TestCase2", 2, "II", nil},
		{"TestCase3", 3, "III", nil},
		{"TestCase4", 4, "IV", nil},
		{"TestCase5", 5, "V", nil},
		{"TestCase6", 6, "VI", nil},
		{"TestCase7", 7, "VII", nil},
		{"TestCase8", 8, "VIII", nil},
		{"TestCase9", 9, "IX", nil},
		{"TestCase10", 10, "X", nil},
		{"TestCase40", 40, "XL", nil},
		{"TestCase50", 50, "L", nil},
		{"TestCase90", 90, "XC", nil},
		{"TestCase100", 100, "C", nil},
		{"TestCase400", 400, "CD", nil},
		{"TestCase500", 500, "D", nil},
		{"TestCase900", 900, "CM", nil},
		{"TestCase999", 999, "CMXCIX", nil},
		{"TestCase1000", 1000, "M", nil},
		{"TestCase2023", 2023, "MMXXIII", nil},
		{"TestCase999999", 999999, "", errors.New("Numerals larger than 3999 don't have a consistent Roman Numeral representation.")},
		{"TestCase3999", 3999, "MMMCMXCIX", nil},
		{"TestCase4000", 4000, "", errors.New("Numerals larger than 3999 don't have a consistent Roman Numeral representation.")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ConvertToRoman(tc.input)
			if err != nil && err.Error() != tc.expectedErr.Error() {
				t.Fatalf("Expected error: %v, got: %v", tc.expectedErr, err)
			}

			if result != tc.expected {
				t.Fatalf("For %s, expected: %s, but got: %s", tc.name, tc.expected, result)
			}
		})
	}
}
