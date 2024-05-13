package apostrophe

import (
	"strings"
	"unicode"
)

func CheckApostrophe(words []string) []string {
	wordString := strings.Join(words, " ")
	wordRune := []rune(wordString)

	for i := 0; i < len(wordRune)-1; i++ {
		// First apostrophe
		if unicode.IsPunct(wordRune[i]) && wordRune[i+1] == ' ' && wordRune[i] == '\'' {
			wordRune[i+1], wordRune[i] = wordRune[i], wordRune[i+1]
		}
	}

	words = strings.Fields(string(wordRune))
	return words
}
