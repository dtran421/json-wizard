package providers_test

import (
	"testing"

	"github.com/dtran421/json-wizard/providers"
	"github.com/dtran421/json-wizard/types"
	"github.com/dtran421/json-wizard/utils"
)

var outputFilepath types.Filepath = utils.Rootpath().Append("/output/yaml/output.yaml")
var inputFilepath types.Filepath = types.NewFilepath("../test/input.json")

var cmd providers.ConvertCmd

func setupTest() {
	cmd = providers.ConvertCmd{}
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
			in:   "xml",
			want: types.XML,
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

func TestValidateInputFile_HappyPath(t *testing.T) {
	setupTest()

	cmd.SetInputFile(inputFilepath.String())

	if err := cmd.ValidateInputFile(); err != nil {
		t.Errorf("ValidateInputFile() == %v, want nil", err)
	}
}

func TestValidateInputFile_Ignored(t *testing.T) {
	cases := []struct {
		inputFile types.Filepath
		input     string
	}{
		{
			inputFile: types.NewFilepath(""),
			input:     "",
		},
		{
			inputFile: inputFilepath,
			input:     `{"key": "value"}`,
		},
		{
			inputFile: types.NewFilepath(""),
			input:     `{"key": "value"}`,
		},
	}

	for _, c := range cases {
		setupTest()

		cmd.SetInputFile(c.inputFile.String())
		cmd.SetInput([]byte(c.input))

		if err := cmd.ValidateInputFile(); err != nil {
			t.Errorf("ValidateInputFile() == %v, want nil", err)
		}
	}

}

func TestValidateInputFile_Error(t *testing.T) {
	cases := []struct {
		inputFile types.Filepath
	}{
		{
			inputFile: types.NewFilepath("invalid.json"),
		},
		{
			inputFile: types.NewFilepath("../test/invalid.json"),
		},
		{
			inputFile: types.NewFilepath("../test/strategy/convert/invalid/input"),
		},
		{
			inputFile: types.NewFilepath("../test/strategy/convert/invalid/input.txt"),
		},
	}

	for _, c := range cases {
		setupTest()

		cmd.SetInputFile(c.inputFile.String())

		if err := cmd.ValidateInputFile(); err == nil {
			t.Errorf("ValidateInputFile() == nil, want error")
		}
	}
}

func TestValidateOutputFile_HappyPath(t *testing.T) {
	cases := []struct {
		outputFile types.Filepath
	}{
		{
			outputFile: types.NewFilepath("output.yaml"),
		},
	}

	for _, c := range cases {
		setupTest()

		cmd.SetOutputFile(c.outputFile.String())

		if err := cmd.ValidateOutputFile(); err != nil {
			t.Errorf("ValidateOutputFile() == %v, want nil", err)
		}
	}
}

func TestValidateOutputFile_Ignored(t *testing.T) {
	setupTest()

	cmd.SetOutputFile("")

	if err := cmd.ValidateOutputFile(); err != nil {
		t.Errorf("ValidateOutputFile() == %v, want nil", err)
	}
}

func TestValidateOutputFile_Error(t *testing.T) {
	cases := []struct {
		outputFile types.Filepath
	}{
		{
			outputFile: types.NewFilepath("output"),
		},
		{
			outputFile: types.NewFilepath("output.txt"),
		},
	}

	for _, c := range cases {
		setupTest()

		cmd.SetOutputFile(c.outputFile.String())

		if err := cmd.ValidateOutputFile(); err == nil {
			t.Errorf("ValidateOutputFile() == nil, want error")
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
	cmd.SetInputFile("../test/input.json")
	cmd.SetOutputFile("output.yaml")

	if err := cmd.ValidateFlags(); err != nil {
		t.Errorf("ValidateFlags() == %v, want nil", err)
	}
}

func TestValidateFlags_Error(t *testing.T) {
	cases := []struct {
		rawOutputFormat string
		inputFile       types.Filepath
		outputFile      string
	}{
		{
			rawOutputFormat: "invalid",
			inputFile:       inputFilepath,
			outputFile:      "output.yaml",
		},
		{
			rawOutputFormat: "yaml",
			inputFile:       types.NewFilepath("invalid.json"),
			outputFile:      "output.yaml",
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
		cmd.SetOutputFile(c.outputFile)

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
			t.Errorf("ValidateFn() == nil, want error")
		}
	}
}

func TestConvertJSON_toYAML_HappyPath(t *testing.T) {
	setupTest()

	cmd.SetInput([]byte(`{"key": "value"}`))

	if err := cmd.ConvertJSON(); err != nil {
		t.Errorf("ConvertJSON() == %v, want nil", err)
	}
}

func TestConvertJSON_Error(t *testing.T) {
	setupTest()

	cmd.SetOutputFormat(types.OutputFormat("invalid"))
	cmd.SetInput([]byte(`{"key": [}}`))

	if err := cmd.ConvertJSON(); err == nil {
		t.Errorf("ConvertJSON() == nil, want error")
	}
}
