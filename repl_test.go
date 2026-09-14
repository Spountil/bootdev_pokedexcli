package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " test again number two ",
			expected: []string{"test", "again", "number", "two"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected: %d words, got: %d words", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expected_word := c.expected[i]
			if word != expected_word {
				t.Errorf("Expected: %v, got: %v", expected_word, word)
				t.Fail()
			}
		}
	}
}
