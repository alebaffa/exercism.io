package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(words []string) FreqMap {
	channel := make(chan FreqMap, len(words))

	for _, value := range words {
		// anonymous function calls Frequency() 3 times in parallel
		// and puts the 3 results in the channel of size 3
		go func(v string) { channel <- Frequency(v) }(value)
	}

	return formatResult(words, channel)
}

func formatResult(words []string, channel chan FreqMap) FreqMap {
	frequency := FreqMap{}
	// loops 3 times because channel of size 3
	for range words {
		for key, value := range <-channel {
			frequency[key] += value
		}
	}
	return frequency
}
