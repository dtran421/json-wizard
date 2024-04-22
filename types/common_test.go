package types_test

import (
	"testing"

	"github.com/dtran421/json-wizard/types"
)

func TestGetExtension(t *testing.T) {
	cases := []struct {
		in       types.OutputFormat
		expected string
	}{
		{
			in:       types.YAML,
			expected: ".yaml",
		},
		{
			in:       types.XML,
			expected: ".xml",
		},
		{
			in:       types.TS,
			expected: ".ts",
		},
		{
			in:       types.GO,
			expected: ".go",
		},
		{
			in:       types.RS,
			expected: ".rs",
		},
		{
			in:       types.OutputFormat(""),
			expected: "",
		},
	}

	for _, c := range cases {
		got := c.in.GetExtension()
		if got != types.Extension(c.expected) {
			t.Errorf("GetExtension(%q) == %q, want %q", c.in, got, c.expected)
		}
	}
}

func TestNewFilepath(t *testing.T) {
	cases := []struct {
		in       string
		expected string
	}{
		{
			in:       "test",
			expected: "test",
		},
		{
			in:       "test2",
			expected: "test2",
		},
		{
			in:       "test3/",
			expected: "test3",
		},
		{
			in:       "",
			expected: ".",
		},
	}

	for _, c := range cases {
		got := types.NewFilepath(c.in)
		if got.String() != c.expected {
			t.Errorf("New(%q) == %q, want %q", c.in, got, c.expected)
		}
	}
}

func TestWithExtension(t *testing.T) {
	cases := []struct {
		in       types.Filepath
		expected string
	}{
		{
			in:       types.NewFilepath("test").WithExtension("yaml"),
			expected: "test.yaml",
		},
		{
			in:       types.NewFilepath("test2.json").WithExtension("json"),
			expected: "test2.json",
		},
		{
			in:       types.NewFilepath("test3/").WithExtension("ts"),
			expected: "test3.ts",
		},
	}

	for _, c := range cases {
		got := c.in
		if got.String() != c.expected {
			t.Errorf("WithExtension(%q) == %q, want %q", c.in, got, c.expected)
		}
	}
}

func TestAppend(t *testing.T) {
	cases := []struct {
		in       types.Filepath
		pathname types.Filepath
		expected string
	}{
		{
			in:       types.NewFilepath("test"),
			pathname: types.NewFilepath("test"),
			expected: "test/test",
		},
		{
			in:       types.NewFilepath("test2"),
			pathname: types.NewFilepath("test2"),
			expected: "test2/test2",
		},
	}

	for _, c := range cases {
		got := c.in.Append(c.pathname)
		if got.String() != c.expected {
			t.Errorf("Append(%q) == %q, want %q", c.in, got, c.expected)
		}
	}
}

func TestDirectory(t *testing.T) {
	cases := []struct {
		in       types.Filepath
		expected string
	}{
		{
			in:       types.NewFilepath("test"),
			expected: ".",
		},
		{
			in:       types.NewFilepath("test/test2"),
			expected: "test",
		},
		{
			in:       types.NewFilepath("test/test3.txt"),
			expected: "test",
		},
		{
			in:       types.NewFilepath("test/test4/test5"),
			expected: "test/test4",
		},
	}

	for _, c := range cases {
		got := c.in.Directory()
		if got != c.expected {
			t.Errorf("Directory(%q) == %q, want %q", c.in, got, c.expected)
		}
	}
}

func TestBase(t *testing.T) {
	cases := []struct {
		in       types.Filepath
		expected string
	}{
		{
			in:       types.NewFilepath("test"),
			expected: "test",
		},
		{
			in:       types.NewFilepath("test/test2"),
			expected: "test2",
		},
		{
			in:       types.NewFilepath("test/test3.txt"),
			expected: "test3.txt",
		},
		{
			in:       types.NewFilepath("test/test4/test5"),
			expected: "test5",
		},
		{
			in:       types.NewFilepath(""),
			expected: "",
		},
	}

	for _, c := range cases {
		got := c.in.Base()
		if got != c.expected {
			t.Errorf("Base(%q) == %q, want %q", c.in, got, c.expected)
		}
	}
}

func TestExtension(t *testing.T) {
	cases := []struct {
		in       types.Filepath
		expected string
	}{
		{
			in:       types.NewFilepath("test"),
			expected: "",
		},
		{
			in:       types.NewFilepath("test2.json"),
			expected: ".json",
		},
		{
			in:       types.NewFilepath("test3.ts"),
			expected: ".ts",
		},
	}

	for _, c := range cases {
		got := c.in.Extension()
		if got != c.expected {
			t.Errorf("GetExtension(%q) == %q, want %q", c.in, got, c.expected)
		}
	}
}

func TestEmpty(t *testing.T) {
	cases := []struct {
		in       types.Filepath
		expected bool
	}{
		{
			in:       types.NewFilepath("test"),
			expected: false,
		},
		{
			in:       types.NewFilepath(""),
			expected: true,
		},
	}

	for _, c := range cases {
		got := c.in.IsEmpty()
		if got != c.expected {
			t.Errorf("IsEmpty(%q) == %t, want %t", c.in, got, c.expected)
		}
	}
}

func TestString(t *testing.T) {
	cases := []struct {
		in       types.Filepath
		expected string
	}{
		{
			in:       types.Filepath("test"),
			expected: "test",
		},
		{
			in:       types.Filepath("test2"),
			expected: "test2",
		},
	}

	for _, c := range cases {
		got := c.in.String()
		if got != c.expected {
			t.Errorf("String(%q) == %q, want %q", c.in, got, c.expected)
		}
	}
}
