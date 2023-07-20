package main

import "testing"

func TestMaskLinks(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "Here's my spammy page: http://hehefouls.netHAHAHA see you.",
			expected: "Here's my spammy page: http://******************* see you.",
		},
		{
			input:    "Check out my website: https://www.example.com",
			expected: "Check out my website: https://***************.com",
		},
		{
			input:    "No links here!",
			expected: "No links here!",
		},
	}

	for _, c := range cases {
		result := maskLinks(c.input)
		if result != c.expected {
			t.Errorf("Expected %q, but got %q", c.expected, result)
		}
	}
}
