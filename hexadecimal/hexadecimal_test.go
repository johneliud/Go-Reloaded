package hexadecimal

import (
	"strings"
	"testing"
)

func TestHexadecimalToDecimal(t *testing.T) {
	input := []string{"1E", "(hex)", "files", "were", "added"}
	expected := []string{"30", "files", "were", "added"}
	output := HexadecimalToDecimal(input)

	expectedString := strings.Join(expected, " ")
	outputString := strings.Join(output, " ")

	if outputString != expectedString {
		t.Errorf(`Expected: %v
Output: %v`, expectedString, outputString)
	}
}
