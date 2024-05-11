package capitalize

import (
	"strings"
	"testing"
)

func TestCapitalize(t *testing.T) {
	input := []string{"Welcome", "to", "the", "Brooklyn", "bridge", "(cap)"}
	expected := []string{"Welcome", "to", "the", "Brooklyn", "Bridge"}
	output := Capitalize(input)

	expectedString := strings.Join(expected, " ")
	outputString := strings.Join(output, " ")

	if outputString != expectedString {
		t.Errorf(`Expected: %v
Output: %v`, expectedString, outputString)
	}
}
