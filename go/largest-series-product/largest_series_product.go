package lsproduct

import "fmt"

const testVersion = 3

func LargestSeriesProduct(digits string, interval int) (int, error) {
	if interval < 0 {
		return 0, fmt.Errorf("Error: interval cannot be negative: %d", interval)
	}
	if len(digits) < interval {
		return 0, fmt.Errorf("Error: digits of len(%s) < interval: %d < %d", digits, len(digits), interval)
	}

	err := checkValidDigits(digits)
	if err != false {
		return 0, fmt.Errorf("input %q contains non-digits", digits)
	}
	digitsInt := convertStringDigitsToInt(digits)

	return calculateMaxProduct(digitsInt, interval), nil
}

func checkValidDigits(digits string) bool {
	for _, digit := range digits {
		if digit < '0' || digit > '9' {
			return true
		}
	}
	return false
}

func convertStringDigitsToInt(digits string) []int {
	digitsInt := make([]int, len(digits))
	for index, digit := range digits {
		digitsInt[index] = int(digit - '0')
	}
	return digitsInt
}

func calculateMaxProduct(v []int, interval int) int {
	maxProduct := 0
	lastPosition := len(v) - interval + 1

	for i := 0; i < lastPosition; i++ {
		currentProd := 1
		for _, d := range v[i : i+interval] {
			currentProd *= d
		}
		if currentProd > maxProduct {
			maxProduct = currentProd
		}
	}
	return maxProduct
}
