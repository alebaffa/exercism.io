package secret

//Handshake takes a decimal number, and converts it to the appropriate sequence of events for a secret handshake.
func Handshake(code int) []string {

	var result []string
	events := []string{"wink", "double blink", "close your eyes", "jump"}

	if code < 0 {
		return result
	}
	/*Here I can directly manipulates the bits in the input number.
		Ex.: 19 = 10011. I can shift the number 1 to the number of positions I want to check in the input number
		(1<<x, x being the num of positions. It will give 2^x value corresponding to the position-x, in case there is 1 in that position).
	  Then, I can AND it with the input number. If it is greater than zero, it means that the bit is set in that position.
	*/
	for index, _ := range events {
		i := uint(index)
		if code&(1<<i) > 0 {
			result = append(result, events[index])
		}
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
