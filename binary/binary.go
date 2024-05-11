package binary

import (
	"fmt"
	"strconv"
)

func BinaryToDecimal(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(bin)" {
			convertedBinary, err := strconv.ParseInt(words[i-1], 2, 64)
			if err != nil {
				fmt.Println("Error converting to binary!", err)
			}
			words[i-1] = strconv.Itoa(int(convertedBinary))
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}
