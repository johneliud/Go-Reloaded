package vowel

import (
	"strings"
	"testing"
)

func TestCorrectVowels(t *testing.T) {
	input := []string{"There", "it", "was.", "A", "amazing", "rock!"}
	expected := []string{"There", "it", "was.", "An", "amazing", "rock!"}
	output := CorrectVowels(input)

	expectedString := strings.Join(expected, " ")
	outputString := strings.Join(output, " ")

	if outputString != expectedString {
		t.Errorf(`Expected: %v
Output: %v`, expectedString, outputString)
	}
}
