package functionality

import (
	"strconv"
	"strings"
)

func CaseModifier(s string) string {

	words := strings.Split(s, " ")
	lastWord := words[len(words)-1]
	lastWord = strings.Trim(lastWord, "()")

	parts := strings.Split(lastWord, ",")

	if len(parts) != 2 || parts[0] != "up" {
		return s
	}
	n, err := strconv.Atoi(parts[1])
	if err != nil {
		return s
	}
	words = words[:len(words)-1]
	for i := len(words) - n; i < len(words); i++ {
		if i >= 0 {
			words[i] = strings.ToUpper(words[i])
		}
	}
	return strings.Join(words, " ")
}
