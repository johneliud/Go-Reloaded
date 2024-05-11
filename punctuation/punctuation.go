package punctuation

import (
	"strings"
	"unicode"
)

func CheckPunctuation(words []string) []string {
	wordString := strings.Join(words, " ")
	wordRune := []rune(wordString)

	for i := 0; i < len(wordRune)-1; i++ {
		if unicode.IsPunct(wordRune[i+1]) && wordRune[i] == ' ' {
			// Swap positions
			wordRune[i+1], wordRune[i] = wordRune[i], wordRune[i+1]
		}

		// Last apostrophe
		if unicode.IsPunct(wordRune[i+1]) && wordRune[i] == ' ' && wordRune[i+1] == '\'' {
			wordRune[i+1], wordRune[i] = wordRune[i], wordRune[i+1]
		}
	}

	words = strings.Fields(string(wordRune))
	return words
}
