package iofile_validator_test

import (
	"testing"

	"github.com/dtran421/json-wizard/providers/convert"
	"github.com/dtran421/json-wizard/providers/iofile_validator"
	"github.com/dtran421/json-wizard/types"
	"github.com/dtran421/json-wizard/utils"
)

var inputFilepath types.Filepath = utils.TestPathname().Append("/input.json")

var validator iofile_validator.IOFileValidator
var cmd convert.ConvertCmd

func setupTest() {
	validator = iofile_validator.IOFileValidator{}
	cmd = *convert.New(validator)
	cmd.SetOutputFormat(types.JSON)
}

func TestValidateInputFile_HappyPath(t *testing.T) {
	setupTest()

	cmd.SetInputFile(inputFilepath.String())

	if err := validator.ValidateInputFile(&cmd); err != nil {
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

		if err := validator.ValidateInputFile(&cmd); err != nil {
			t.Errorf("ValidateInputFile() == %v, want nil", err)
		}
	}

}

func TestValidateInputFile_Error(t *testing.T) {
	cases := []struct {
		inputFile types.Filepath
	}{
		{
			inputFile: utils.TestPathname().Append("invalid.json"),
		},
		{
			inputFile: utils.TestPathname().Append("invalid/input.txt"),
		},
		{
			inputFile: utils.TestPathname().Append("invalid/input"),
		},
	}

	for _, c := range cases {
		setupTest()

		cmd.SetInputFile(c.inputFile.String())

		if err := validator.ValidateInputFile(&cmd); err == nil {
			t.Errorf("ValidateInputFile() == nil, want error")
		}
	}
}

func TestValidateOutputFile_HappyPath(t *testing.T) {
	cases := []struct {
		outputFile types.Filepath
	}{
		{
			outputFile: types.NewFilepath("output.json"),
		},
	}

	for _, c := range cases {
		setupTest()

		cmd.SetOutputFile(c.outputFile.String())

		if err := validator.ValidateOutputFile(&cmd); err != nil {
			t.Errorf("ValidateOutputFile() == %v, want nil", err)
		}
	}
}

func TestValidateOutputFile_Ignored(t *testing.T) {
	setupTest()

	cmd.SetOutputFile("")

	if err := validator.ValidateOutputFile(&cmd); err != nil {
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

		if err := validator.ValidateOutputFile(&cmd); err == nil {
			t.Errorf("ValidateOutputFile() == nil, want error")
		}

	}
}
