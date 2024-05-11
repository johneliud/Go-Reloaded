package uppercase

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertToUpperCase(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		} else if words[i] == "(up," {
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
				words[j] = strings.ToUpper(words[j])
			}
			words = append(words[:i], words[i+2:]...)
		}
	}
	return words
}
