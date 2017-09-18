package twofer

// ShareWith compose the final sentence given the input name.
func ShareWith(name string) string {
	result := "One for "
	if name == "" {
		result += "you"
	} else {
		result += name
	}
	return result + ", one for me."
}
