package functionality

import "strings"

func TextNormalizer(s string) string {
	var builder strings.Builder
	inBracket := false

	for _, val := range s {
		switch val {
		case '(':
			inBracket = true
			builder.WriteRune(val)
		case ')':
			inBracket = false
			builder.WriteRune(val)
		case ' ':
			if !inBracket {
				builder.WriteRune(val)
			}
		default:
			builder.WriteRune(val)
		}
	}

	return builder.String()
}
