package apostrophe

import (
	"strings"
	"testing"
)

func TestCheckApostrophe(t *testing.T) {
	input := []string{"I", "am", "exactly", "how", "they", "describe", "me:", "' awesome '"}
	expected := []string{"I", "am", "exactly", "how", "they", "describe", "me:", "'awesome '"}
	output := CheckApostrophe(input)

	expectedString := strings.Join(expected, " ")
	outputString := strings.Join(output, " ")

	if outputString != expectedString {
		t.Errorf(`Expected: %v
Output: %v`, expectedString, outputString)
	}
}
