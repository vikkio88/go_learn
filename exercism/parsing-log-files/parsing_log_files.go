package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(ERR|WRN|TRC|DBG|INF|FTL)\].+$`)
	res := re.MatchString(text)

	return res
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<(-|\*|~|=)*>`)
	res := re.ReplaceAllString(text, ";")

	return strings.Split(res, ";")
}

func CountQuotedPasswords(lines []string) int {
	// weird syntax for case insensitive
	re := regexp.MustCompile(`(?i)".*password.*"`)
	acc := 0
	for _, line := range lines {
		if re.MatchString(line) {
			acc++
		}
	}

	return acc
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+\b(\w+?)\b`)
	var newLines []string
	for _, line := range lines {
		if !re.MatchString(line) {
			newLines = append(newLines, line)
			continue
		}
		newLines = append(newLines, fmt.Sprintf("[USR] %s %s", re.FindStringSubmatch(line)[1], line))

	}

	return newLines
}
