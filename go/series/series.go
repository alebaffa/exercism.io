package slice

// All returns the contiguous substrings of given length in the given number
func All(length int, number string) []string {
	var substrings []string
	for index := 0; index < len(number); index++ {
		if index+length <= len(number) {
			substrings = append(substrings, number[index:index+length])
		}
	}

	return substrings
}

// Frist returns the first length elements in the given number
func Frist(length int, number string) string {
	return number[:length]
}
