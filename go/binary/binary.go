package binary

import "strconv"
import "regexp"
import "fmt"

const testVersion = 2

/* Using the standard library:
	i, err := strconv.ParseInt(input, 2, 64)
}*/

func ParseBinary(input string) (int, error) {

	if !isValid(input) {
		return 0, fmt.Errorf("found not valid characters")
	}

	result := 0
	i := len(input) - 1
	for _, n := range input {
		num, _ := strconv.Atoi(string(n))
		result += num << uint(i)
		i--
	}
	return result, nil
}

func isValid(input string) bool {
	regex := regexp.MustCompile("[a-z]+|[2]+")
	word := regex.FindAllString(input, -1)
	return len(word) == 0
}
