package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	seen := make(map[rune]bool)
	for _, s := range word {
		if seen[s] {
			return false
		}
		if unicode.IsLetter(s) {
			seen[s] = true
		}
	}
	return true
}
