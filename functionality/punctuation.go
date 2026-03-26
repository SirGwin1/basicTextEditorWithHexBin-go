package functionality

import "regexp"

func Punctuation(str string) string {
	// Remove spaces before punctuation
	reBefore := regexp.MustCompile(`\s+([,.;:!?])`)
	str = reBefore.ReplaceAllString(str, `$1`)

	// Ensure single space after punctuation
	reAfter := regexp.MustCompile(`([,.;:!?])([^\s])`)
	str = reAfter.ReplaceAllString(str, `$1 $2`)

	// Fix spaces around quotes
	reQuote := regexp.MustCompile(`\s*(['"])\s*`)
	str = reQuote.ReplaceAllString(str, `$1`)

	return str
}
