package lowercase

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertToLowerCase(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		} else if words[i] == "(low," {
			convertedInteger, err := strconv.Atoi(words[i+1][:len(words[i+1])-1])
			if err != nil {
				fmt.Println("Error converting derivative!", err)
			}

			// Handle derivative errors
			if convertedInteger < 1 || convertedInteger > len(words[:i]) {
				fmt.Println("Invalid derivative passed!")
				break
			}

			for j := i - convertedInteger; j < i; j++ {
				words[j] = strings.ToLower(words[j])
			}
			words = append(words[:i], words[i+2:]...)
		}
	}
	return words
}
