package common

import (
	"regexp"
	"strings"
)

func CleanString(input string) string {
	// Remove all tab characters
	result := strings.ReplaceAll(input, "\t", "")

	// Replace multiple spaces with a single space
	spaceRegex := regexp.MustCompile(`\s{2,}`)
	result = spaceRegex.ReplaceAllString(result, " ")
	result = strings.ReplaceAll(result, "\n", " ")

	// Trim leading and trailing spaces
	return strings.TrimSpace(result)
}
