package cmd

import (
	"github.com/dtran421/json-wizard/providers/format"
	"github.com/dtran421/json-wizard/utils"
	"github.com/spf13/cobra"
)

var formatCmdStruct = format.FormatCmd{}

var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Format JSON.",
	Long: `
	Format JSON to pretty print.
	`,

	Args: formatCmdStruct.ValidateFn,

	RunE: utils.CmdRunE(&formatCmdStruct),
}

func init() {
	rootCmd.AddCommand(formatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// formatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	var inputFile string
	var outputFile string

	var indentSize int

	formatCmd.Flags().StringVarP(&inputFile, "inputFile", "in", "",
		"input file to format (will ignore if input is provided)")

	formatCmdStruct.SetInputFile(inputFile)

	formatCmd.Flags().StringVarP(&outputFile, "outputFile", "o", "",
		"output file path to write the formatted JSON")

	formatCmdStruct.SetOutputFile(outputFile)

	formatCmd.Flags().IntVarP(&indentSize, "indent", "i", 2,
		"indent size for the output file")

	formatCmdStruct.SetIndentSize(indentSize)
}
