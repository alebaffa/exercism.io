package etl

import "strings"

/* Transform extracts some scrabble scores from a legacy system. */
func Transform(oldMap map[int][]string) map[string]int {
	newMap := map[string]int{}

	for key, values := range oldMap {
		for _, value := range values {
			newMap[strings.ToLower(value)] = key
		}
	}
	return newMap
}
