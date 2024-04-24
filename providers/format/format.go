package format

import (
	"encoding/json"
	"fmt"

	"github.com/dtran421/json-wizard/providers/iofile_validator"
	"github.com/dtran421/json-wizard/strategy/formatter"
	"github.com/dtran421/json-wizard/types"
	"github.com/spf13/cobra"
)

type FormatCmd struct {
	IOFileValidator iofile_validator.IOFileValidator
	Formatter       formatter.Formatter

	cmdName string

	input     json.RawMessage
	inputFile types.Filepath

	outputFormat types.OutputFormat
	outputFile   types.Filepath

	indentSize int
}

func New(ioFileValidator iofile_validator.IOFileValidator) *FormatCmd {
	return &FormatCmd{
		IOFileValidator: ioFileValidator,
		Formatter:       *formatter.NewFormatter(),

		cmdName: "format",

		outputFormat: types.JSON,
	}
}

func (cmdStruct FormatCmd) CmdName() string {
	return cmdStruct.cmdName
}

func (cmdStruct FormatCmd) OutputFormat() types.OutputFormat {
	return cmdStruct.outputFormat
}

func (cmdStruct FormatCmd) Input() json.RawMessage {
	return cmdStruct.input
}

func (cmdStruct *FormatCmd) SetInput(input json.RawMessage) {
	cmdStruct.input = input
}

func (cmdStruct FormatCmd) InputFile() types.Filepath {
	return cmdStruct.inputFile
}

func (cmdStruct *FormatCmd) SetInputFile(inputFile string) {
	cmdStruct.inputFile = types.NewFilepath(inputFile)
}

func (cmdStruct FormatCmd) OutputFile() types.Filepath {
	return cmdStruct.outputFile
}

func (cmdStruct *FormatCmd) SetOutputFile(outputFile string) {
	cmdStruct.outputFile = types.NewFilepath(outputFile)
}

func (cmdStruct FormatCmd) IndentSize() int {
	return cmdStruct.indentSize
}

func (cmdStruct *FormatCmd) SetIndentSize(indentSize int) {
	cmdStruct.indentSize = indentSize
}

func (cmdStruct FormatCmd) ValidateFn(cmd *cobra.Command, args []string) error {
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

func (cmdStruct FormatCmd) ValidateFlags() error {
	if err := cmdStruct.IOFileValidator.ValidateInputFile(&cmdStruct); err != nil {
		return err
	}

	if err := cmdStruct.IOFileValidator.ValidateOutputFile(&cmdStruct); err != nil {
		return err
	}

	return nil
}

func (cmdStruct FormatCmd) Execute() error {
	fmt.Printf("Formatting %s\n", cmdStruct.input)

	if err := cmdStruct.Formatter.Format(&cmdStruct); err != nil {
		return err
	}

	return nil
}
