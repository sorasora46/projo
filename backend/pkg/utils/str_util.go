package utils

import "strings"

func PascalToCamelCase(s *string) {
	if s == nil || *s == "" {
		return
	}

	first := strings.ToLower(string((*s)[0]))
	if len(*s) > 1 {
		*s = first + (*s)[1:]
	} else {
		*s = first
	}
}

func GenValidationErrAtTag(tag string) string {
	if tag == "" {
		return ""
	}
	return "validation error at \"" + tag + "\" tag"
}
