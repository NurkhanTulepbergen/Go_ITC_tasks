package reverse

import (
	"testing"
)

func TestReverse(t *testing.T) {
	// The first test is checking for an empty string

	input := ""
	expected := ""
	result := Reverse(input)
	if result != expected {
		t.Errorf("Reverse(%q) == %q, excepted %q", input, result, expected)
	}

	// The second test is checking for a string with one character
	input = "a"
	expected = "a"
	result = Reverse(input)
	if result != expected {
		t.Errorf("Reverse(%q) == %q, excepted %q", input, result, expected)
	}

	// The third test is a check for a string with one word
	input = "hello"
	expected = "olleh"
	result = Reverse(input)
	if result != expected {
		t.Errorf("Reverse(%q) == %q, excepted %q", input, result, expected)
	}

	// The fourth test is a check for a string with several characters
	input = "hello world"
	expected = "dlrow olleh"
	result = Reverse(input)
	if result != expected {
		t.Errorf("Reverse(%q) == %q, excepted %q", input, result, expected)
	}
}
