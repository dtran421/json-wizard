package format_test

import (
	"testing"

	"github.com/dtran421/json-wizard/providers/format"
	"github.com/dtran421/json-wizard/providers/input_validator"
	"github.com/dtran421/json-wizard/types"
	"github.com/dtran421/json-wizard/utils"
)

var inputFilepath types.Filepath = utils.TestPathname().Append("/input.json")
var outputFilepath types.Filepath = utils.TestPathname().Append("/output/json/output.json")

var cmd format.FormatCmd

func setupTest() {
	validator := input_validator.InputValidator{}
	cmd = *format.New(validator)
	cmd.SetOutputFile(outputFilepath.String())
}

func TestValidateFlags_HappyPath(t *testing.T) {
	setupTest()

	cmd.SetInputFile(inputFilepath.String())
	cmd.SetOutputFile(outputFilepath.String())

	if err := cmd.ValidateFlags(); err != nil {
		t.Errorf("ValidateFlags(%s, %s) == %v, want nil", inputFilepath, outputFilepath, err)
	}
}

func TestValidateFlags_Error(t *testing.T) {
	cases := []struct {
		inputFile  types.Filepath
		outputFile types.Filepath
	}{
		{
			inputFile:  types.NewFilepath("invalid.yaml"),
			outputFile: outputFilepath,
		},
		{
			inputFile:  inputFilepath,
			outputFile: "output",
		},
	}

	for _, c := range cases {
		setupTest()

		cmd.SetInputFile(c.inputFile.String())
		cmd.SetOutputFile(c.outputFile.String())

		if err := cmd.ValidateFlags(); err == nil {
			t.Errorf("ValidateFlags(%s, %s) == nil, want error", c.inputFile, c.outputFile)
		}
	}
}

func TestValidateFn_HappyPath(t *testing.T) {
	cases := []struct {
		args      []string
		inputFile types.Filepath
	}{
		{
			args:      []string{`{"key": "value"}`},
			inputFile: types.NewFilepath(""),
		},
		{
			args:      []string{},
			inputFile: inputFilepath,
		},
		{
			args:      []string{`{"key": "value"}`},
			inputFile: inputFilepath,
		},
	}

	for _, c := range cases {
		setupTest()

		cmd.SetInputFile(c.inputFile.String())

		if err := cmd.ValidateFn(nil, c.args); err != nil {
			t.Errorf("ValidateFn(%s, %s) == %v, want nil", c.args, c.inputFile, err)
		}
	}
}

func TestValidateFn_Error(t *testing.T) {
	cases := []struct {
		inputFile types.Filepath
		args      []string
	}{
		{
			inputFile: types.NewFilepath(""),
			args:      []string{""},
		},
		{
			inputFile: types.NewFilepath(""),
			args:      []string{`{"key": [}}`},
		},
		{
			inputFile: types.NewFilepath("invalid.json"),
			args:      []string{},
		},
	}

	for _, c := range cases {
		setupTest()

		cmd.SetInputFile(c.inputFile.String())

		if err := cmd.ValidateFn(nil, c.args); err == nil {
			t.Errorf("ValidateFn(%s, %v) == nil, want error", c.inputFile, c.args)
		}
	}
}

func TestExecute_HappyPath(t *testing.T) {
	setupTest()

	cmd.SetInput([]byte(`{"key": "value"}`))

	if err := cmd.Execute(); err != nil {
		t.Errorf("Execute() == %v, want nil", err)
	}
}

func TestExecute_Error(t *testing.T) {
	setupTest()

	cmd.SetInput([]byte(`{"key": [}}`))

	if err := cmd.Execute(); err == nil {
		t.Errorf("Execute() == nil, want error")
	}
}
