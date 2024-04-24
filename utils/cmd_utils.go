package utils

import (
	"fmt"

	"github.com/dtran421/json-wizard/types"
	"github.com/spf13/cobra"
)

type RunEFn func(*cobra.Command, []string) error

func CmdRunE(cmdStruct types.Cmd) RunEFn {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%s called", cmdStruct.CmdName())

		if err := cmdStruct.Execute(); err != nil {
			return err
		}

		return nil
	}
}
