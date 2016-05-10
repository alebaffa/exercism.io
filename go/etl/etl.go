package etl

import "strings"

func Transform(oldMap map[int][]string) map[string]int {
	newMap := map[string]int{}

	for num, letters := range oldMap {
		for _, letter := range letters {
			newMap[strings.ToLower(letter)] = num
		}
	}
	return newMap
}
