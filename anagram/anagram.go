package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	anagrams := make([]string, 0)

	subject = strings.ToLower(subject)
	for _, candidate := range candidates {
		if isAnagram(subject, strings.ToLower(candidate)) {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}

func isAnagram(sub, cand string) bool {
	return sub != cand && sortString(sub) == sortString(cand)
}

func sortString(in string) string {
	inSplit := strings.Split(in, "")
	sort.Strings(inSplit)
	return strings.Join(inSplit, "")
}
