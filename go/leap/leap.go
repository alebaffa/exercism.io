package leap

// TestVersion is the testing version number.
const TestVersion = 1

// IsLeapYear checks if the input year is leap or not.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
