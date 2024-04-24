package cmd

import (
	"github.com/dtran421/json-wizard/providers/convert"
	"github.com/dtran421/json-wizard/providers/iofile_validator"
	"github.com/dtran421/json-wizard/utils"
	"github.com/spf13/cobra"
)

var convertCmdStruct = convert.New(iofile_validator.IOFileValidator{})

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert JSON to other formats.",
	Long: `
	Convert JSON to other formats, including:

	- YAML
	- TypeScript
	- Go struct
	- Rust struct
	`,

	Args: convertCmdStruct.ValidateFn,

	RunE: utils.CmdRunE(convertCmdStruct),
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	var rawOutputFormat string
	var inputFile string
	var outputFile string

	var indentSize int

	convertCmd.Flags().StringVarP(&rawOutputFormat, "outputFormat", "f", "",
		"output format (required, one of: yaml, ts, go, rs)")
	rootCmd.MarkFlagRequired("output")

	convertCmdStruct.SetRawOutputFormat(rawOutputFormat)

	convertCmd.Flags().StringVarP(&inputFile, "inputFile", "in", "",
		"input file to convert (will ignore if input is provided)")

	convertCmdStruct.SetInputFile(inputFile)

	convertCmd.Flags().StringVarP(&outputFile, "outputFile", "o", "",
		"output file path to write the converted JSON")

	convertCmdStruct.SetOutputFile(outputFile)

	convertCmd.Flags().IntVarP(&indentSize, "indent", "i", 2,
		"indent size for the output file")

	convertCmdStruct.SetIndentSize(indentSize)
}
