package vowel

func CorrectVowels(words []string) []string {
	vowelsIncludingH := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
	consonants := []string{"b", "c", "d", "f", "g", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z", "B", "C", "D", "F", "G", "J", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "V", "W", "X", "Y", "Z"}

	for i, word := range words {
		for _, vowel := range vowelsIncludingH {
			// Check if current article is "a" and the first letter of the next word is a vowel
			if word == "a" && string(words[i+1][0]) == vowel {
				words[i] = "an"
			} else if word == "A" && string(words[i+1][0]) == vowel {
				words[i] = "An"
			}

			// Check if current article is "an" and the first letter of the next word is a vowel
			if word == "an" && string(words[i+1][0]) == vowel {
				words[i] = "an"
			} else if word == "An" && string(words[i+1][0]) == vowel {
				words[i] = "An"
			}
		}
	}

	for i, word := range words {
		for _, consonant := range consonants {
			// Check if current article is "an" and the first letter of the next word is not a vowel
			if word == "an" && string(words[i+1][0]) == consonant {
				words[i] = "a"
			} else if word == "An" && string(words[i+1][0]) == consonant {
				words[i] = "A"
			}
	}
}

	return words
}