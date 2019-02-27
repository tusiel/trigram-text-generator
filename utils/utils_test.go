package utils

import (
	"testing"
)

func TestRemoveUnwantedCharacters(t *testing.T) {
	input := "This i@£%@£s a st^^$ring!"

	expected := "This is a string!"

	output, err := RemoveUnwantedCharacters(input)
	if err != nil {
		t.Errorf("Error calling RemoveUnwantedCharacters: %+v", err)
	}

	if output != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, output)
	}

}

func TestIsSentenceEnd(t *testing.T) {
	sentence1 := "This is a sentence."

	r, err := IsSentenceEnd(sentence1)
	if err != nil {
		t.Errorf("Error calling IsSentenceEnd: %+v", err)
	}

	if !r {
		t.Errorf("Expected '%s' to return true, but returned %t", sentence1, r)
	}

	sentence2 := "This is another sentence"

	r, err = IsSentenceEnd(sentence2)
	if err != nil {
		t.Errorf("Error calling IsSentenceEnd: %+v", err)
	}

	if r {
		t.Errorf("Expected '%s' to return false, but returned %t", sentence2, r)
	}
}
