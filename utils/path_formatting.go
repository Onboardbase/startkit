package utils

import (
	"regexp"
	"strings"
)

func NormalizeToKebabOrSnakeCase(str string) string {
	STRING_DASHERIZE_REGEXP := regexp.MustCompile(`\s`)
	STRING_DECAMELIZE_REGEXP := regexp.MustCompile(`([a-z\d])([A-Z])`)
	// Replace camelCase with dashes.
	str = STRING_DECAMELIZE_REGEXP.ReplaceAllString(str, "$1-$2")
	str = strings.ToLower(str)
	// Replace spaces with dashes.
	str = STRING_DASHERIZE_REGEXP.ReplaceAllString(str, "-")
	return str
}
