package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(?:ERR|DBG|TRC|INF|WRN|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[-*~=]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) (count int) {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)

	linesOut := make([]string, len(lines))

	for i, line := range lines {
		result := re.FindStringSubmatch(line)
		if result == nil {
			linesOut[i] = line
		} else {
			linesOut[i] = fmt.Sprintf("[USR] %s %s", result[1], line)
		}

	}
	return linesOut
}
