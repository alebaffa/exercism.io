package secret

//Handshake takes a decimal number, and converts it to the appropriate sequence of events for a secret handshake.
func Handshake(code int) []string {

	var result []string

	if code < 0 {
		return result
	}

	if code&(1<<0) > 0 {
		result = append(result, "wink")
	}
	if code&(1<<1) > 0 {
		result = append(result, "double blink")
	}
	if code&(1<<2) > 0 {
		result = append(result, "close your eyes")
	}
	if code&(1<<3) > 0 {
		result = append(result, "jump")
	}
	if code&(1<<4) > 0 {
		result = reverse(result)
	}
	return result
}

func reverse(slice []string) []string {
	// Reverse slice
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
