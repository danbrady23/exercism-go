package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {

	// Remove any spaces
	normId := strings.ReplaceAll(id, " ", "")

	// Strings with length less than 2 are not valid
	if len(normId) < 2 {
		return false
	}

	var checkSum int = 0
	var pos int = 0
	for i := len(normId) - 1; i >= 0; i-- {

		// Try and convert string to digit
		digit, err := strconv.Atoi(string(normId[i]))

		// If a non-digit is encountered then the sequence isn't valid
		if err != nil {
			return false
		}

		// If we're in an even position then just add the number
		if pos%2 == 0 {
			checkSum += digit
			// If we're in an odd position then before adding double and subtract 9 if needed
		} else {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
			checkSum += digit
		}
		pos++
	}

	return checkSum%10 == 0
}
