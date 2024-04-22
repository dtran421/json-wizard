package types

import "github.com/spf13/cobra"

type Cmd interface {
	/*
	 * Validate the arguments passed to the command.
	 */
	ValidateFn(cmd *cobra.Command, args []string) error

	/*
	 * Validate the flags passed to the command.
	 */
	ValidateFlags() error

	/*
	 * Validate the output format flag.
	 */
	ValidateOutputFormat() error

	/*
	 * Convert JSON to the specified output format.
	 */
	Execute() error
}
