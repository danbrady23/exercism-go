package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	var out = map[string]int{}

	for key, value := range in {
		for _, letter := range value {
			lowerLetter := strings.ToLower(letter)
			out[lowerLetter] = key
		}
	}

	return out
}
