package utils_test

import (
	"testing"

	"github.com/dtran421/json-wizard/utils"
)

func TestGetIndent(t *testing.T) {
	cases := []struct {
		in       int
		expected string
	}{
		{
			in:       0,
			expected: "",
		},
		{
			in:       1,
			expected: " ",
		},
		{
			in:       2,
			expected: "  ",
		},
		{
			in:       3,
			expected: "   ",
		},
	}

	for _, c := range cases {
		got := utils.GetIndent(c.in)
		if got != c.expected {
			t.Errorf("GetIndent(%q) == %q, want %q", c.in, got, c.expected)
		}
	}
}

func TestGetCustomIndent(t *testing.T) {
	cases := []struct {
		in       int
		indent   int
		expected string
	}{
		{
			in:       0,
			indent:   1,
			expected: "",
		},
		{
			in:       1,
			indent:   1,
			expected: " ",
		},
		{
			in:       2,
			indent:   1,
			expected: "  ",
		},
		{
			in:       3,
			indent:   1,
			expected: "   ",
		},
		{
			in:       0,
			indent:   2,
			expected: "",
		},
		{
			in:       1,
			indent:   2,
			expected: "  ",
		},
		{
			in:       2,
			indent:   2,
			expected: "    ",
		},
		{
			in:       3,
			indent:   2,
			expected: "      ",
		},
		{
			in:       0,
			indent:   4,
			expected: "",
		},
		{
			in:       1,
			indent:   4,
			expected: "    ",
		},
		{
			in:       2,
			indent:   4,
			expected: "        ",
		},
		{
			in:       3,
			indent:   4,
			expected: "            ",
		},
	}

	for _, c := range cases {
		got := utils.GetCustomIndent(c.in, c.indent)
		if got != c.expected {
			t.Errorf("GetCustomIndent(%q) == %q, want %q", c.in, got, c.expected)
		}
	}
}
