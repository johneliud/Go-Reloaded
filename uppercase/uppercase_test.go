package uppercase

import (
	"strings"
	"testing"
)

func TestConvertToUpperCase(t *testing.T) {
	input := []string{"Ready,", "set,", "go", "(up)", "!"}
	expected := []string{"Ready,", "set,", "GO", "!"}
	output := ConvertToUpperCase(input)

	expectedString := strings.Join(expected, " ")
	outputString := strings.Join(output, " ")

	if outputString != expectedString {
		t.Errorf(`Expected: %v
Output: %v`, expectedString, outputString)
	}
}
