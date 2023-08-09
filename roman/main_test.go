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
		{"NegativeNumber", -9, "", errors.New("No representation found. Numbers start with 1 in Roman Numeral representation.")},
		{"TestCase0", 0, "", errors.New("No representation found. Numbers start with 1 in Roman Numeral representation.")},
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
		{"TestCase20", 20, "XX", nil},
		{"TestCase29", 29, "XXIX", nil},
		{"TestCase30", 30, "XXX", nil},
		{"TestCase40", 40, "XL", nil},
		{"TestCase50", 50, "L", nil},
		{"TestCase60", 60, "LX", nil},
		{"TestCase70", 70, "LXX", nil},
		{"TestCase80", 80, "LXXX", nil},
		{"TestCase90", 90, "XC", nil},
		{"TestCase100", 100, "C", nil},
		{"TestCase200", 200, "CC", nil},
		{"TestCase294", 294, "CCXCIV", nil},
		{"TestCase300", 300, "CCC", nil},
		{"TestCase400", 400, "CD", nil},
		{"TestCase500", 500, "D", nil},
		{"TestCase600", 600, "DC", nil},
		{"TestCase700", 700, "DCC", nil},
		{"TestCase800", 800, "DCCC", nil},
		{"TestCase900", 900, "CM", nil},
		{"TestCase999", 999, "CMXCIX", nil},
		{"TestCase1000", 1000, "M", nil},
		{"TestCase1234", 1234, "MCCXXXIV", nil},
		{"TestCase1987", 1987, "MCMLXXXVII", nil},
		{"TestCase2019", 2019, "MMXIX", nil},
		{"TestCase2023", 2023, "MMXXIII", nil},
		{"TestCase3999", 3999, "MMMCMXCIX", nil},
		{"TestCase4000", 4000, "", errors.New("Numerals larger than 3999 don't have a consistent Roman Numeral representation.")},
		{"TestCase999999", 999999, "", errors.New("Numerals larger than 3999 don't have a consistent Roman Numeral representation.")},
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
