package types_test

import (
	"testing"

	"github.com/dtran421/json-wizard/types"
	"github.com/dtran421/json-wizard/utils"
)

func TestHomeDirpath(t *testing.T) {
	homedirpath := types.HomeDirpath()
	expected := types.NewFilepath("/Users/dtran")

	if homedirpath != expected {
		t.Errorf("HomeDirpath() == %q, want %q", homedirpath, expected)
	}
}

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
		expected types.Filepath
	}{
		{
			in:       "test",
			expected: types.HomeDirpath() + "/test",
		},
		{
			in:       "test2",
			expected: types.HomeDirpath() + "/test2",
		},
		{
			in:       "test3/",
			expected: types.HomeDirpath() + "/test3",
		},
		{
			in:       "",
			expected: types.HomeDirpath(),
		},
	}

	for _, c := range cases {
		got := types.NewFilepath(c.in)
		if got != c.expected {
			t.Errorf("New(%q) == %q, want %q", c.in, got, c.expected)
		}
	}
}

func TestNewFilepathFromAbsPath(t *testing.T) {
	cases := []struct {
		in       string
		expected types.Filepath
	}{
		{
			in:       "/test",
			expected: "/test",
		},
		{
			in:       "/test2",
			expected: "/test2",
		},
		{
			in:       "/test3/",
			expected: "/test3",
		},
		{
			in:       "test4/test5",
			expected: "/test4/test5",
		},
		{
			in:       "/",
			expected: "/",
		},
		{
			in:       "",
			expected: "/",
		},
	}

	for _, c := range cases {
		got := types.NewFilepathFromAbsPath(c.in)
		if got != c.expected {
			t.Errorf("NewFilepathFromAbsPath(%q) == %q, want %q", c.in, got, c.expected)
		}
	}
}

func TestWithPrefix(t *testing.T) {
	cases := []struct {
		in       types.Filepath
		prefix   types.Filepath
		expected types.Filepath
	}{
		{
			in:       types.NewFilepath("test"),
			prefix:   types.NewFilepathFromAbsPath("test"),
			expected: types.NewFilepathFromAbsPath("test/test"),
		},
		{
			in:       types.NewFilepath("test2"),
			prefix:   types.NewFilepathFromAbsPath("test"),
			expected: types.NewFilepathFromAbsPath("test/test2"),
		},
		{
			in:       types.NewFilepath("test3"),
			prefix:   types.NewFilepathFromAbsPath("test2"),
			expected: types.NewFilepathFromAbsPath("test2/test3"),
		},
		{
			in:       types.NewFilepath("test4").WithPrefix("test"),
			prefix:   types.HomeDirpath(),
			expected: types.NewFilepath("test/test4"),
		},
		{
			in:       types.NewFilepath("test5").WithPrefix(""),
			prefix:   utils.Rootpath(),
			expected: types.NewFilepathFromAbsPath("test5").WithPrefix(utils.Rootpath()),
		},
		{
			in:       types.NewFilepathFromAbsPath("/"),
			prefix:   types.HomeDirpath(),
			expected: types.HomeDirpath(),
		},
		{
			in:       types.NewFilepathFromAbsPath("/test6"),
			prefix:   types.NewFilepathFromAbsPath("/"),
			expected: types.NewFilepathFromAbsPath("/test6"),
		},
	}

	for _, c := range cases {
		got := c.in.WithPrefix(c.prefix)
		if got != c.expected {
			t.Errorf("WithPrefix(%q, %q) == %q, want %q", c.in, c.prefix, got, c.expected)
		}
	}
}

func TestWithExtension(t *testing.T) {
	cases := []struct {
		in       types.Filepath
		expected types.Filepath
	}{
		{
			in:       types.NewFilepath("test").WithExtension("yaml"),
			expected: types.NewFilepath("test.yaml"),
		},
		{
			in:       types.NewFilepath("test2.json").WithExtension("json"),
			expected: types.NewFilepath("test2.json"),
		},
		{
			in:       types.NewFilepath("test3/").WithExtension("ts"),
			expected: types.NewFilepath("test3.ts"),
		},
		{
			in:       types.NewFilepath("test4/test5").WithExtension("yaml"),
			expected: types.NewFilepath("test4/test5.yaml"),
		},
		{
			in:       types.NewFilepath("test6.yaml").WithExtension("yaml"),
			expected: types.NewFilepath("test6.yaml"),
		},
	}

	for _, c := range cases {
		got := c.in
		if got != c.expected {
			t.Errorf("WithExtension(%q) == %q, want %q", c.in, got, c.expected)
		}
	}
}

func TestAppend(t *testing.T) {
	cases := []struct {
		in       types.Filepath
		pathname string
		expected types.Filepath
	}{
		{
			in:       types.NewFilepath("test"),
			pathname: "test",
			expected: types.NewFilepath("test/test"),
		},
		{
			in:       types.NewFilepath("test2"),
			pathname: "test2",
			expected: types.NewFilepath("test2/test2"),
		},
		{
			in:       types.NewFilepathFromAbsPath("/"),
			pathname: "test3",
			expected: types.NewFilepath("/test3"),
		},
	}

	for _, c := range cases {
		got := c.in.Append(c.pathname)
		if got != c.expected {
			t.Errorf("Append(%q, %q) == %q, want %q", c.in, c.pathname, got, c.expected)
		}
	}
}

func TestDirectory(t *testing.T) {
	cases := []struct {
		in       types.Filepath
		expected types.Filepath
	}{
		{
			in:       types.NewFilepath("test"),
			expected: types.HomeDirpath(),
		},
		{
			in:       types.NewFilepath("test/test2"),
			expected: types.HomeDirpath().Append("test"),
		},
		{
			in:       types.NewFilepath("test/test3.txt"),
			expected: types.HomeDirpath().Append("test"),
		},
		{
			in:       types.NewFilepath("test/test4/test5"),
			expected: types.HomeDirpath().Append("test").Append("test4"),
		},
	}

	for _, c := range cases {
		got := c.in.Directory()
		if got != c.expected.String() {
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
			expected: types.HomeDirpath().Base(),
		},
		{
			in:       types.NewFilepathFromAbsPath("/"),
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

func TestHasPrefix(t *testing.T) {
	cases := []struct {
		in       types.Filepath
		prefix   types.Filepath
		expected bool
	}{
		{
			in:       types.NewFilepath("test"),
			prefix:   types.NewFilepath("test"),
			expected: true,
		},
		{
			in:       types.NewFilepath("test2"),
			prefix:   types.NewFilepath("test"),
			expected: true,
		},
		{
			in:       types.NewFilepath("test3"),
			prefix:   types.NewFilepath("test2"),
			expected: false,
		},
		{
			in:       types.NewFilepath("test4"),
			prefix:   types.HomeDirpath(),
			expected: true,
		},
	}

	for _, c := range cases {
		got := c.in.HasPrefix(c.prefix)
		if got != c.expected {
			t.Errorf("HasPrefix(%q, %q) == %t, want %t", c.in, c.prefix, got, c.expected)
		}
	}
}

func TestIsAtHomeDir(t *testing.T) {
	cases := []struct {
		in       types.Filepath
		expected bool
	}{
		{
			in:       types.NewFilepath("test"),
			expected: false,
		},
		{
			in:       types.HomeDirpath(),
			expected: true,
		},
	}

	for _, c := range cases {
		got := c.in.IsAtHomeDir()
		if got != c.expected {
			t.Errorf("IsAtHomeDir(%q) == %t, want %t", c.in, got, c.expected)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	cases := []struct {
		in       types.Filepath
		expected bool
	}{
		{
			in:       types.NewFilepath("test"),
			expected: false,
		},
		{
			in:       types.NewFilepathFromAbsPath(""),
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
		expected types.Filepath
	}{
		{
			in:       types.NewFilepath("test"),
			expected: types.HomeDirpath().Append("test"),
		},
		{
			in:       types.NewFilepath("test2"),
			expected: types.HomeDirpath().Append("test2"),
		},
		{
			in:       types.NewFilepath(""),
			expected: types.HomeDirpath(),
		},
	}

	for _, c := range cases {
		got := c.in.String()
		if got != c.expected.String() {
			t.Errorf("String(%q) == %q, want %q", c.in, got, c.expected)
		}
	}
}

func TestString_Empty(t *testing.T) {
	cases := []struct {
		in       types.Filepath
		expected string
	}{
		{
			in:       types.NewFilepathFromAbsPath(""),
			expected: "/",
		},
		{
			in:       types.NewFilepathFromAbsPath("/"),
			expected: "/",
		},
	}

	for _, c := range cases {
		got := c.in.String()
		if got != c.expected {
			t.Errorf("String_Empty(%q) == %q, want %q", c.in, got, c.expected)
		}
	}
}
