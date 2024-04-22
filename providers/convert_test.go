package providers_test

import (
	"fmt"
	"testing"

	"github.com/dtran421/json-wizard/providers"
	"github.com/dtran421/json-wizard/types/convert"
)

var cmd providers.ConvertCmd

func setupTest() {
	cmd = providers.ConvertCmd{}
}

func TestValidateOutputFormat_HappyPath(t *testing.T) {
	cases := []struct {
		in   string
		want convert.OutputFormat
	}{
		{
			in:   "yaml",
			want: convert.YAML,
		},
		{
			in:   "xml",
			want: convert.XML,
		},
		{
			in:   "ts",
			want: convert.TS,
		},
		{
			in:   "go",
			want: convert.GO,
		},
		{
			in:   "rs",
			want: convert.RS,
		},
	}

	for _, c := range cases {
		setupTest()

		cmd.SetRawOutputFormat(c.in)
		cmd.ValidateOutputFormat()

		if got := cmd.GetOutputFormat(); got != c.want {
			t.Errorf("ValidateOutputFormat(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestValidateOutputFormat_Error(t *testing.T) {
	cases := []struct {
		in string
	}{
		{
			in: "invalid",
		},
		{
			in: "",
		},
	}

	for _, c := range cases {
		setupTest()

		cmd.SetRawOutputFormat(c.in)

		if err := cmd.ValidateOutputFormat(); err == nil {
			t.Errorf("ValidateOutputFormat(%q) == nil, want error", c.in)
		}
	}
}

func TestValidateInputFile_HappyPath(t *testing.T) {
	setupTest()

	cmd.SetInputFile("../test/input.json")

	if err := cmd.ValidateInputFile(); err != nil {
		t.Errorf("ValidateInputFile() == %v, want nil", err)
	}
}

func TestValidateInputFile_Ignored(t *testing.T) {
	setupTest()

	cmd.SetInputFile("")

	if err := cmd.ValidateInputFile(); err != nil {
		t.Errorf("ValidateInputFile() == %v, want nil", err)
	}

	cmd.SetInputFile("../test/input.json")
	cmd.SetInput([]byte(`{"key": "value"}`))

	if err := cmd.ValidateInputFile(); err != nil {
		t.Errorf("ValidateInputFile() == %v, want nil", err)
	}

	cmd.SetInputFile("")
	cmd.SetInput([]byte(`{"key": "value"}`))

	if err := cmd.ValidateInputFile(); err != nil {
		t.Errorf("ValidateInputFile() == %v, want nil", err)
	}
}

func TestValidateInputFile_Error(t *testing.T) {
	setupTest()

	cmd.SetInputFile("invalid.json")

	if err := cmd.ValidateInputFile(); err == nil {
		t.Errorf("ValidateInputFile() == nil, want error")
	}

	cmd.SetInputFile("../test/invalid.json")

	if err := cmd.ValidateInputFile(); err == nil {
		t.Errorf("ValidateInputFile() == nil, want error")
	}

	cmd.SetInputFile("../test/input")

	if err := cmd.ValidateInputFile(); err == nil {
		t.Errorf("ValidateInputFile() == nil, want error")
	}

	cmd.SetInputFile("../test/input.txt")

	if err := cmd.ValidateInputFile(); err == nil {
		t.Errorf("ValidateInputFile() == nil, want error")
	}
}

func TestValidateFlags_HappyPath(t *testing.T) {
	setupTest()

	cmd.SetRawOutputFormat("yaml")
	cmd.SetInputFile("../test/input.json")

	if err := cmd.ValidateFlags(); err != nil {
		t.Errorf("ValidateFlags() == %v, want nil", err)
	}
}

func TestValidateFlags_Error(t *testing.T) {
	setupTest()

	cmd.SetRawOutputFormat("invalid")
	cmd.SetInputFile("../test/input.json")

	if err := cmd.ValidateFlags(); err == nil {
		t.Errorf("ValidateFlags() == nil, want error")
	}

	cmd.SetRawOutputFormat("yaml")
	cmd.SetInputFile("invalid.json")

	if err := cmd.ValidateFlags(); err == nil {
		t.Errorf("ValidateFlags() == nil, want error")
	}
}

func TestValidateFn_HappyPath(t *testing.T) {
	setupTest()

	args := []string{`{"key": "value"}`}
	cmd.SetRawOutputFormat("yaml")

	if err := cmd.ValidateFn(nil, args); err != nil {
		t.Errorf("ValidateFn() == %v, want nil", err)
	}

	cmd.SetInputFile("../test/input.json")

	if err := cmd.ValidateFn(nil, args); err != nil {
		t.Errorf("ValidateFn() == %v, want nil", err)
	}

	args = []string{}

	if err := cmd.ValidateFn(nil, args); err != nil {
		t.Errorf("ValidateFn() == %v, want nil", err)
	}
}

func TestValidateFn_Error(t *testing.T) {
	setupTest()

	cmd.SetRawOutputFormat("yaml")

	args := []string{}

	if err := cmd.ValidateFn(nil, args); err == nil {
		t.Errorf("ValidateFn() == nil, want error")
	}

	cmd.SetRawOutputFormat("invalid")

	args = []string{`{"key": "value"}`}

	if err := cmd.ValidateFn(nil, args); err == nil {
		t.Errorf("ValidateFn() == nil, want error")
	}

	args = []string{`{"key": [}}`}

	if err := cmd.ValidateFn(nil, args); err == nil {
		fmt.Printf("%v\n", err)
		t.Errorf("ValidateFn() == nil, want error")
	}
}
