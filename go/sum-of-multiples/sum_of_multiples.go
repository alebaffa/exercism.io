package summultiples

func MultipleSummer(divisors ...int) func(int) int {
	return func(limit int) int {
		count := 0
		for i := 1; i < limit; i++ {
			for _, divisor := range divisors {
				if i%divisor == 0 {
					count += i
					break
				}
			}
		}
		return count
	}
}
