package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

func Valid(number string) bool {
	number = removeSpaces(number)
	return containsValidDigits(number) && (calculateTotalSum(number))%10 == 0

}

func AddCheck(number string) string {
	partialNumber := removeSpaces(number)
	partialNumber += "0"
	sum := calculateTotalSum(partialNumber)
	checkDigit := calculateCheckDigit(sum)
	number += strconv.Itoa(checkDigit)
	return number
}

func removeSpaces(input string) string {
	return strings.Replace(input, " ", "", -1)
}

func calculateCheckDigit(sum int) int {
	checkDigit := 0
	if sum%10 != 0 {
		checkDigit = 10 - (sum % 10)
	}
	return checkDigit
}

func containsValidDigits(number string) bool {
	match, _ := regexp.MatchString("([1-9]+)", number)
	if match == false {
		return false
	}
	return true
}

func calculateTotalSum(number string) int {
	return sumOfEvenDigits(number) + sumOfOddDigits(number)
}

func sumOfEvenDigits(number string) int {
	totalEven := 0
	for index := len(number) - 2; index >= 0; index -= 2 {
		evenNumber, _ := strconv.Atoi(string(number[index]))
		evenNumber *= 2
		if evenNumber >= 10 {
			evenNumber -= 9
		}
		totalEven += evenNumber
	}
	return totalEven
}

func sumOfOddDigits(number string) int {
	totalOdd := 0
	for index := len(number) - 1; index >= 0; index -= 2 {
		oddNumber, _ := strconv.Atoi(string(number[index]))
		totalOdd += oddNumber
	}
	return totalOdd
}
