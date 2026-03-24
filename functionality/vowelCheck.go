package functionality

import (
	"strings"
	"unicode"
)

func Vowelcheck(str string) string {
	words := strings.Fields(str)
	result := []string{}

	for i, word := range words {
		if (word == "a" || word == "A") && i+1 < len(words) {
			next := []rune(words[i+1])

			if len(next) > 0 && isVowel(next[0]) {
				if word == "A" {
					word = "An"
				} else {
					word = "an"
				}
			}
		}
		result = append(result, word)
	}

	return strings.Join(result, " ")
}

func isVowel(r rune) bool {
	r = unicode.ToLower(r)
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'h'
}