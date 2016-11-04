package binary

import "errors"

import "regexp"

const testVersion = 2

/* Using the standard library:
	i, err := strconv.ParseInt(input, 2, 64)
}*/

func ParseBinary(input string) (int, error) {

	if !isValid(input) {
		return 0, errors.New("found not valid characters")
	}

	result := 0
	for i, n := range input {
		if n == '1' {
			result += 1 << uint(len(input)-i-1)
		}
	}
	return result, nil
}

func isValid(input string) bool {
	return regexp.MustCompile("\\A[01]*\\z").MatchString(input)
}
