package convert_test

import (
	"testing"

	"github.com/dtran421/json-wizard/providers/convert"
	"github.com/dtran421/json-wizard/types"
	"github.com/dtran421/json-wizard/utils"
)

var outputFilepath types.Filepath = utils.TestPathname().Append("/output/yaml/output.yaml")
var inputFilepath types.Filepath = utils.TestPathname().Append("/input.json")

var cmd convert.ConvertCmd

func setupTest() {
	cmd = convert.ConvertCmd{}
	cmd.SetOutputFormat(types.YAML)
	cmd.SetOutputFile(outputFilepath.String())
}

func TestValidateOutputFormat_HappyPath(t *testing.T) {
	cases := []struct {
		in   string
		want types.OutputFormat
	}{
		{
			in:   "yaml",
			want: types.YAML,
		},
		{
			in:   "ts",
			want: types.TS,
		},
		{
			in:   "go",
			want: types.GO,
		},
		{
			in:   "rs",
			want: types.RS,
		},
	}

	for _, c := range cases {
		setupTest()

		cmd.SetRawOutputFormat(c.in)
		cmd.ValidateOutputFormat()

		if got := cmd.OutputFormat(); got != c.want {
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

func TestValidateIndentSize_HappyPath(t *testing.T) {
	cases := []struct {
		indentSize int
	}{
		{
			indentSize: 0,
		},
		{
			indentSize: 2,
		},
		{
			indentSize: 4,
		},
	}

	for _, c := range cases {
		setupTest()

		cmd.SetIndentSize(c.indentSize)

		if err := cmd.ValidateIndentSize(); err != nil {
			t.Errorf("ValidateIndentSize() == %v, want nil", err)
		}
	}
}

func TestValidateIndentSize_Error(t *testing.T) {
	cases := []struct {
		indentSize int
	}{
		{
			indentSize: -1,
		},
	}

	for _, c := range cases {
		setupTest()

		cmd.SetIndentSize(c.indentSize)

		if err := cmd.ValidateIndentSize(); err == nil {
			t.Errorf("ValidateIndentSize() == nil, want error")
		}
	}
}

func TestValidateFlags_HappyPath(t *testing.T) {
	setupTest()

	cmd.SetRawOutputFormat("yaml")
	cmd.SetInputFile(inputFilepath.String())
	cmd.SetOutputFile(outputFilepath.String())

	if err := cmd.ValidateFlags(); err != nil {
		t.Errorf("ValidateFlags() == %v, want nil", err)
	}
}

func TestValidateFlags_Error(t *testing.T) {
	cases := []struct {
		rawOutputFormat string
		inputFile       types.Filepath
		outputFile      types.Filepath
	}{
		{
			rawOutputFormat: "invalid",
			inputFile:       inputFilepath,
			outputFile:      outputFilepath,
		},
		{
			rawOutputFormat: "yaml",
			inputFile:       types.NewFilepath("invalid.json"),
			outputFile:      outputFilepath,
		},
		{
			rawOutputFormat: "yaml",
			inputFile:       inputFilepath,
			outputFile:      "output",
		},
	}

	for _, c := range cases {
		setupTest()

		cmd.SetRawOutputFormat(c.rawOutputFormat)
		cmd.SetInputFile(c.inputFile.String())
		cmd.SetOutputFile(c.outputFile.String())

		if err := cmd.ValidateFlags(); err == nil {
			t.Errorf("ValidateFlags() == nil, want error")
		}
	}
}

func TestValidateFn_HappyPath(t *testing.T) {
	cases := []struct {
		rawOutputFormat string
		args            []string
		inputFile       types.Filepath
	}{
		{
			rawOutputFormat: "yaml",
			args:            []string{`{"key": "value"}`},
			inputFile:       types.NewFilepath(""),
		},
		{
			rawOutputFormat: "yaml",
			args:            []string{},
			inputFile:       inputFilepath,
		},
		{
			rawOutputFormat: "yaml",
			args:            []string{`{"key": "value"}`},
			inputFile:       inputFilepath,
		},
	}

	for _, c := range cases {
		setupTest()

		cmd.SetRawOutputFormat(c.rawOutputFormat)
		cmd.SetInputFile(c.inputFile.String())

		if err := cmd.ValidateFn(nil, c.args); err != nil {
			t.Errorf("ValidateFn() == %v, want nil", err)
		}
	}
}

func TestValidateFn_Error(t *testing.T) {
	cases := []struct {
		rawOutputFormat string
		inputFile       types.Filepath
		args            []string
	}{
		{
			rawOutputFormat: "yaml",
			inputFile:       types.NewFilepath(""),
			args:            []string{""},
		},
		{
			rawOutputFormat: "invalid",
			inputFile:       types.NewFilepath(""),
			args:            []string{`{"key": "value"}`},
		},
		{
			rawOutputFormat: "yaml",
			inputFile:       types.NewFilepath(""),
			args:            []string{`{"key": [}}`},
		},
		{
			rawOutputFormat: "yaml",
			inputFile:       types.NewFilepath("invalid.json"),
			args:            []string{},
		},
	}

	for _, c := range cases {
		setupTest()

		cmd.SetRawOutputFormat(c.rawOutputFormat)
		cmd.SetInputFile(c.inputFile.String())

		if err := cmd.ValidateFn(nil, c.args); err == nil {
			t.Errorf("ValidateFn(%s, %s, %v) == nil, want error", c.rawOutputFormat, c.inputFile, c.args)
		}
	}
}

func TestExecute_toYAML_HappyPath(t *testing.T) {
	setupTest()

	cmd.SetInput([]byte(`{"key": "value"}`))

	if err := cmd.Execute(); err != nil {
		t.Errorf("Execute() == %v, want nil", err)
	}
}

func TestExecute_Error(t *testing.T) {
	setupTest()

	cmd.SetOutputFormat(types.OutputFormat("invalid"))
	cmd.SetInput([]byte(`{"key": [}}`))

	if err := cmd.Execute(); err == nil {
		t.Errorf("Execute() == nil, want error")
	}
}
