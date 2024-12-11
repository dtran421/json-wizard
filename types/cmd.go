package types

import (
	"encoding/json"

	"github.com/spf13/cobra"
)

type InputValidator interface {
	ValidateInputFile(Cmd) error
	ValidateOutputFile(Cmd) error
}

type Cmd interface {
	CmdName() string

	Input() json.RawMessage
	InputFile() Filepath
	OutputFile() Filepath
	OutputFormat() OutputFormat
	IndentSize() int

	SetOutputFile(string)

	/*
	 * Validate the arguments passed to the command.
	 */
	ValidateFn(*cobra.Command, []string) error

	/*
	 * Validate the flags passed to the command.
	 */
	ValidateFlags() error

	/*
	 * Execute the command.
	 */
	Execute() error
}
