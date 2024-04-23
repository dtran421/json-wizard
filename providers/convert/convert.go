package convert

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/dtran421/json-wizard/strategy/convert"
	"github.com/dtran421/json-wizard/types"
	"github.com/spf13/cobra"
)

// ConvertCmd represents the convert command
type ConvertCmd struct {
	Converter convert.ConverterFactory

	rawOutputFormat string

	input json.RawMessage

	outputFormat types.OutputFormat
	inputFile    types.Filepath
	outputFile   types.Filepath

	indentSize int
}

func (cmdStruct *ConvertCmd) SetRawOutputFormat(rawOutputFormat string) {
	cmdStruct.rawOutputFormat = rawOutputFormat
}

func (cmdStruct *ConvertCmd) SetInput(input json.RawMessage) {
	cmdStruct.input = input
}

func (cmdStruct ConvertCmd) OutputFormat() types.OutputFormat {
	return cmdStruct.outputFormat
}

func (cmdStruct *ConvertCmd) SetOutputFormat(outputFormat types.OutputFormat) {
	cmdStruct.outputFormat = outputFormat
}

func (cmdStruct *ConvertCmd) SetInputFile(inputFile string) {
	cmdStruct.inputFile = types.NewFilepath(inputFile)
}

func (cmdStruct *ConvertCmd) SetOutputFile(outputFile string) {
	cmdStruct.outputFile = types.NewFilepath(outputFile)
}

func (cmdStruct *ConvertCmd) SetIndentSize(indentSize int) {
	cmdStruct.indentSize = indentSize
}

func (cmdStruct ConvertCmd) ValidateFn(cmd *cobra.Command, args []string) error {
	if err := cmdStruct.ValidateFlags(); err != nil {
		return err
	}

	if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
		if cmdStruct.inputFile.IsEmpty() {
			return err
		}

		return nil
	}

	if json.Valid([]byte(args[0])) {
		cmdStruct.input = json.RawMessage(args[0])
		return nil
	}

	return fmt.Errorf("invalid JSON provided: %s", args[0])
}

func (cmdStruct ConvertCmd) ValidateFlags() error {
	if err := cmdStruct.ValidateOutputFormat(); err != nil {
		return err
	}

	if err := cmdStruct.ValidateInputFile(); err != nil {
		return err
	}

	if err := cmdStruct.ValidateOutputFile(); err != nil {
		return err
	}

	return nil
}

func (cmdStruct *ConvertCmd) ValidateOutputFormat() error {
	if cmdStruct.rawOutputFormat == "" {
		return fmt.Errorf("output format is required")
	}

	switch cmdStruct.rawOutputFormat {
	case string(types.YAML), string(types.TS), string(types.GO), string(types.RS):
		cmdStruct.outputFormat = types.OutputFormat(cmdStruct.rawOutputFormat)
		return nil
	default:
		return fmt.Errorf("invalid output format specified: %s", cmdStruct.rawOutputFormat)
	}
}

func (cmdStruct ConvertCmd) ValidateInputFile() error {
	if cmdStruct.inputFile.IsEmpty() || cmdStruct.inputFile.IsAtHomeDir() {
		return nil
	}

	if cmdStruct.input != nil {
		fmt.Println("ignoring input file as JSON input is provided")
		return nil
	}

	if _, err := os.Stat(cmdStruct.inputFile.String()); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("input file does not exist: %s", cmdStruct.inputFile)
	}

	if extension := cmdStruct.inputFile.Extension(); extension != ".json" {
		return fmt.Errorf("input file must be a JSON file")
	}

	return nil
}

func (cmdStruct *ConvertCmd) ValidateOutputFile() error {
	if types.NewFilepathFromAbsPath(cmdStruct.outputFile.Base()).IsEmpty() {
		cmdStruct.outputFile = types.NewFilepath("output").WithExtension(cmdStruct.outputFormat)
		return nil
	}

	var outputFormatExtension = cmdStruct.outputFormat.GetExtension()
	if extension := cmdStruct.outputFile.Extension(); extension != string(outputFormatExtension) {
		return fmt.Errorf("output file must have the extension %s, got %s", outputFormatExtension, extension)
	}

	return nil
}

func (cmdStruct ConvertCmd) ValidateIndentSize() error {
	if cmdStruct.indentSize < 0 {
		return fmt.Errorf("indent size must be a positive integer")
	}

	return nil
}

func (cmdStruct ConvertCmd) ConvertJSON() error {
	fmt.Printf("Converting %s to %s\n", cmdStruct.input, cmdStruct.outputFormat)

	convertStrategy, err := cmdStruct.Converter.BuildConverter(cmdStruct.outputFormat)
	if err != nil {
		return err
	}

	convertStrategy.SetInput(cmdStruct.input)
	convertStrategy.SetInputFile(cmdStruct.inputFile)
	convertStrategy.SetOutputFile(cmdStruct.outputFile)

	convertStrategy.Convert()

	return nil
}
