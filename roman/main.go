package main

import (
	"errors"
)

var romanNumeral = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(num int) (roman string, err error) {
	err = nil
	roman = ""
	if num <= 0 {
		err = errors.New("No representation found. Numbers start with 1 in Roman Numeral representation.")
		return
	}
	if num > 3999 {
		err = errors.New("Numerals larger than 3999 don't have a consistent Roman Numeral representation.")
		return
	}
	for _, numeral := range romanNumeral {
		for num >= numeral.value {
			roman += numeral.symbol
			num -= numeral.value
		}
	}
	return roman, err
}
