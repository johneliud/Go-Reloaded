package binary

import (
	"strings"
	"testing"
)

func TestBinaryToDecimal(t *testing.T) {
	input := []string{"It", "has", "been", "10", "(bin)", "years"}
	expected := []string{"It", "has", "been", "2", "years"}
	output := BinaryToDecimal(input)

	expectedString := strings.Join(expected, " ")
	outputString := strings.Join(output, " ")

	if outputString != expectedString {
		t.Errorf(`Expected: %v
Output: %v`, expectedString, outputString)
	}
}
