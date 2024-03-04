package reverse

import (
	"strings"
)

func Reverse(input string) string {
	strSlice := strings.Split(input, "")

	for i, j := 0, len(strSlice)-1; i < j; i, j = i+1, j-1 {
		strSlice[i], strSlice[j] = strSlice[j], strSlice[i]
	}

	return strings.Join(strSlice, "")
}
