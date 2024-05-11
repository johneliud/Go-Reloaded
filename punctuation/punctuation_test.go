package punctuation

import (
	"strings"
	"testing"
)

func TestCheckPunctuation(t *testing.T) {
	input := []string{"I", "was", "sitting", "over", "there", ",and", "then", "BAMM !!"}
	expected := []string{"I", "was", "sitting", "over", "there,", "and", "then", "BAMM!!"}
	output := CheckPunctuation(input)

	expectedString := strings.Join(expected, " ")
	outputString := strings.Join(output, " ")

	if outputString != expectedString {
		t.Errorf(`Expected: %v
Output: %v`, expectedString, outputString)
	}
}
