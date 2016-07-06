package romannumerals

import (
	"bytes"
	"fmt"
)

const testVersion = 2

type arabicToRoman struct {
	arabic int
	roman  string
}

func ToRomanNumeral(number int) (string, error) {
	buffer := bytes.NewBufferString("")

	if number <= 0 || number >= 4000 {
		return "", fmt.Errorf("Number %d is not allowed", number)
	}

	dictionary := []arabicToRoman{
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

	for _, item := range dictionary {
		for number >= item.arabic {
			buffer.WriteString(item.roman)
			number -= item.arabic
		}
	}
	return buffer.String(), nil
}
