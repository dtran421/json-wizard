package convert

import (
	"encoding/json"
	"fmt"

	"github.com/dtran421/json-wizard/providers/input_validator"
	"github.com/dtran421/json-wizard/strategy/converter"
	"github.com/dtran421/json-wizard/types"
	"github.com/spf13/cobra"
)

// ConvertCmd represents the convert command
type ConvertCmd struct {
	ConverterFactory converter.ConverterFactoryInstance
	InputValidator   input_validator.InputValidator

	cmdName string

	rawOutputFormat string

	input     json.RawMessage
	inputFile types.Filepath

	outputFormat types.OutputFormat
	outputFile   types.Filepath

	indentSize int
}

func New(ioFileValidator input_validator.InputValidator) *ConvertCmd {
	return &ConvertCmd{
		ConverterFactory: *converter.ConverterFactory(),
		InputValidator:   ioFileValidator,

		cmdName: "convert",
	}
}

func (cmdStruct ConvertCmd) CmdName() string {
	return cmdStruct.cmdName
}

func (cmdStruct *ConvertCmd) SetRawOutputFormat(rawOutputFormat string) {
	cmdStruct.rawOutputFormat = rawOutputFormat
}

func (cmdStruct ConvertCmd) Input() json.RawMessage {
	return cmdStruct.input
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

func (cmdStruct ConvertCmd) InputFile() types.Filepath {
	return cmdStruct.inputFile
}

func (cmdStruct *ConvertCmd) SetInputFile(inputFile string) {
	cmdStruct.inputFile = types.NewFilepath(inputFile)
}

func (cmdStruct ConvertCmd) OutputFile() types.Filepath {
	return cmdStruct.outputFile
}

func (cmdStruct *ConvertCmd) SetOutputFile(outputFile string) {
	cmdStruct.outputFile = types.NewFilepath(outputFile)
}

func (cmdStruct ConvertCmd) IndentSize() int {
	return cmdStruct.indentSize
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

	if err := cmdStruct.InputValidator.ValidateInputFile(&cmdStruct); err != nil {
		return err
	}

	if err := cmdStruct.InputValidator.ValidateOutputFile(&cmdStruct); err != nil {
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

func (cmdStruct ConvertCmd) Execute() error {
	fmt.Printf("Converting %s to %s\n", cmdStruct.input, cmdStruct.outputFormat)

	convertStrategy, err := cmdStruct.ConverterFactory.BuildConverter(cmdStruct.outputFormat)
	if err != nil {
		return err
	}

	convertStrategy.SetInput(cmdStruct.input)
	convertStrategy.SetInputFile(cmdStruct.inputFile)
	convertStrategy.SetOutputFile(cmdStruct.outputFile)

	convertStrategy.Convert()

	return nil
}
