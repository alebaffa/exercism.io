package accumulate

const testVersion = 1

func Accumulate(items []string, converter func(string) string) []string {
	convertedItems := []string{}

	for _, item := range items {
		convertedItems = append(convertedItems, converter(item))
	}
	return convertedItems
}
