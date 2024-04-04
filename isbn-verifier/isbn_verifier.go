package isbn

import (
	"strconv"
	"unicode"
)

func IsValidISBN(isbn string) bool {

	d := 10
	var total int
	for _, char := range isbn {
		if char == '-' {
			continue
		}

		if char == 'X' && d == 1 {
			total += 10
		} else if unicode.IsDigit(char) {
			val, _ := strconv.Atoi(string(char))
			total += (val * d)
		} else {
			return false
		}
		d--
	}
	return total%11 == 0 && d == 0
}
