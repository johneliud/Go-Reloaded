package hexadecimal

import (
	"fmt"
	"strconv"
)

func HexadecimalToDecimal(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" {
			convertedHexadecimal, err := strconv.ParseInt(words[i-1], 16, 64)
			if err != nil {
				fmt.Println("Error converting to hexadecimal!", err)
				break
			}
			words[i-1] = strconv.Itoa(int(convertedHexadecimal))
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}
