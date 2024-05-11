package lowercase

import (
	"strings"
	"testing"
)

func TestConvertToLowerCase(t *testing.T) {
	input := []string{"I", "should", "stop", "SHOUTING", "(low)"}
	expected := []string{"I", "should", "stop", "shouting"}
	output := ConvertToLowerCase(input)

	expectedString := strings.Join(expected, " ")
	outputString := strings.Join(output, " ")

	if outputString != expectedString {
		t.Errorf(`Expected: %v
Output: %v`, expectedString, outputString)
	}
}
