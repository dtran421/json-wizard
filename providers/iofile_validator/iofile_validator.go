package iofile_validator

import (
	"errors"
	"fmt"
	"os"

	"github.com/dtran421/json-wizard/types"
)

type IOFileValidator struct{}

func (v IOFileValidator) ValidateInputFile(cmd types.Cmd) error {
	inputFile := cmd.InputFile()

	if inputFile.IsEmpty() || inputFile.IsAtHomeDir() {
		return nil
	}

	if cmd.Input() != nil {
		fmt.Println("ignoring input file as JSON input is provided")
		return nil
	}

	if _, err := os.Stat(inputFile.String()); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("input file does not exist: %s", inputFile)
	}

	if extension := inputFile.Extension(); extension != ".json" {
		return fmt.Errorf("input file must be a JSON file")
	}

	return nil
}

func (v IOFileValidator) ValidateOutputFile(cmd types.Cmd) error {
	outputFile := cmd.OutputFile()
	outputFormat := cmd.OutputFormat()

	if types.NewFilepathFromAbsPath(outputFile.Base()).IsEmpty() {
		cmd.SetOutputFile(types.NewFilepath("output").WithExtension(outputFormat).String())
		return nil
	}

	if !outputFile.HasExtension(outputFormat) {
		return fmt.Errorf("output file must have the extension %s, got %s",
			outputFormat.GetExtension(), outputFile.Extension())
	}

	return nil
}
