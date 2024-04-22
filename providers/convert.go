package providers

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	convertStrategy "github.com/dtran421/json-wizard/strategy/convert"
	"github.com/dtran421/json-wizard/types/convert"
	"github.com/spf13/cobra"
)

// ConvertCmd represents the convert command
type ConvertCmd struct {
	rawOutputFormat string

	input json.RawMessage

	outputFormat convert.OutputFormat
	inputFile    string
	outputFile   string
}

func (cmdStruct ConvertCmd) GetRawOutputFormat() string {
	return cmdStruct.rawOutputFormat
}

func (cmdStruct *ConvertCmd) SetRawOutputFormat(rawOutputFormat string) {
	cmdStruct.rawOutputFormat = rawOutputFormat
}

func (cmdStruct ConvertCmd) GetInput() json.RawMessage {
	return cmdStruct.input
}

func (cmdStruct *ConvertCmd) SetInput(input json.RawMessage) {
	cmdStruct.input = input
}

func (cmdStruct ConvertCmd) GetOutputFormat() convert.OutputFormat {
	return cmdStruct.outputFormat
}

func (cmdStruct *ConvertCmd) SetOutputFormat(outputFormat convert.OutputFormat) {
	cmdStruct.outputFormat = outputFormat
}

func (cmdStruct ConvertCmd) GetInputFile() string {
	return cmdStruct.inputFile
}

func (cmdStruct *ConvertCmd) SetInputFile(inputFile string) {
	cmdStruct.inputFile = inputFile
}

func (cmdStruct ConvertCmd) GetOutputFile() string {
	return cmdStruct.outputFile
}

func (cmdStruct *ConvertCmd) SetOutputFile(outputFile string) {
	cmdStruct.outputFile = outputFile
}

func (cmdStruct ConvertCmd) ValidateFn(cmd *cobra.Command, args []string) error {
	if err := cmdStruct.ValidateFlags(); err != nil {
		return err
	}

	if err := cobra.MinimumNArgs(1)(cmd, args); err != nil && cmdStruct.inputFile == "" {
		return err
	}

	if cmdStruct.inputFile != "" {
		return nil
	}

	// Validate the JSON input
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

	return nil
}

func (cmdStruct *ConvertCmd) ValidateOutputFormat() error {
	if cmdStruct.rawOutputFormat == "" {
		return fmt.Errorf("output format is required")
	}

	switch cmdStruct.rawOutputFormat {
	case string(convert.YAML), string(convert.XML), string(convert.TS), string(convert.GO), string(convert.RS):
		cmdStruct.outputFormat = convert.OutputFormat(cmdStruct.rawOutputFormat)
		return nil
	default:
		return fmt.Errorf("invalid output format specified: %s", cmdStruct.rawOutputFormat)
	}
}

func (cmdStruct ConvertCmd) ValidateInputFile() error {
	if cmdStruct.inputFile == "" {
		return nil
	}

	if cmdStruct.input != nil {
		fmt.Println("ignoring input file as JSON input is provided")
		return nil
	}

	if _, err := os.Stat(cmdStruct.inputFile); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("input file does not exist: %s", cmdStruct.inputFile)
	}

	if extension := filepath.Ext(cmdStruct.inputFile); extension != ".json" {
		return fmt.Errorf("input file must be a JSON file")
	}

	return nil
}

func (cmdStruct ConvertCmd) ConvertJSON() error {
	fmt.Printf("Converting %s to %s\n", cmdStruct.input, cmdStruct.outputFormat)

	switch cmdStruct.outputFormat {
	case convert.YAML:
		fmt.Println("Converting to YAML")

		yamlConvert := convertStrategy.YAMLConverter{}
		yamlConvert.SetInput(cmdStruct.input)
		yamlConvert.SetInputFile(cmdStruct.inputFile)

		yamlConvert.Convert()

	case convert.XML:
		fmt.Println("Converting to XML")
	case convert.TS:
		fmt.Println("Converting to TypeScript")
	case convert.GO:
		fmt.Println("Converting to Go struct")
	case convert.RS:
		fmt.Println("Converting to Rust struct")
	}

	return nil
}
