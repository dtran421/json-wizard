/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/dtran421/json-wizard/providers"
	"github.com/spf13/cobra"
)

var cmdStruct = providers.ConvertCmd{}

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert JSON to other formats.",
	Long: `
	Convert JSON to other formats, including:

	- YAML
	- XML
	- TypeScript
	- Go struct
	- Rust struct
	`,

	Args: cmdStruct.ValidateFn,

	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("convert called")
		if err := cmdStruct.ConvertJSON(); err != nil {
			return err
		}
		return nil
	},
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

	convertCmd.Flags().StringVarP(&rawOutputFormat, "outputFormat", "o", "",
		"output format (required, one of: yaml, xml, ts, go, rs)")
	rootCmd.MarkFlagRequired("output")

	cmdStruct.SetRawOutputFormat(rawOutputFormat)

	convertCmd.Flags().StringVarP(&inputFile, "inputFile", "i", "",
		"input file to convert (will ignore if input is provided)")

	cmdStruct.SetInputFile(inputFile)

	convertCmd.Flags().StringVarP(&outputFile, "outputFile", "f", "",
		"output file to write the converted JSON")

	cmdStruct.SetOutputFile(outputFile)
}
