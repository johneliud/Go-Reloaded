package vowel

func CorrectVowels(words []string) []string {
	vowelsIncludingH := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}

	for i, word := range words {
		for _, vowel := range vowelsIncludingH {
			if word == "A" && string(words[i+1][0]) == vowel {
				words[i] = "An"
			} else if word == "a" && string(words[i+1][0]) == vowel {
				words[i] = "an"
			} else if word == "An" && string(words[i+1][0]) != vowel {
				words[i] = "A"
			} else if word == "an" && string(words[i+1][0]) != vowel {
				words[i] = "a"
			}
		}
	}

	return words
}
