/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"encoding/json"

	"github.com/spf13/cobra"
)

type OutputFormat string

const (
	JSON OutputFormat = "json"
	YAML              = "yaml"
	XML               = "xml"
	TS                = "ts"
	GO                = "go"
	RS                = "rs"
)

var rawOutputFormat string

var input json.RawMessage
var outputFormat OutputFormat
var inputFile string

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

	Args: validateFn,

	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("convert called")
		if err := ConvertJSON(); err != nil {
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
	convertCmd.Flags().StringVarP(&rawOutputFormat, "outputFormat", "o", "",
		"Output format (required, one of: yaml, xml, ts, go, rs)")
	rootCmd.MarkFlagRequired("output")

	convertCmd.Flags().StringVarP(&inputFile, "inputFile", "i", "",
		"Input file to convert (will ignore if input is provided)")
}

/*
 * Validate the arguments passed to the command.
 */
func validateFn(cmd *cobra.Command, args []string) error {
	if err := validateFlags(); err != nil {
		return err
	}

	if len(args) == 0 && inputFile == "" {
		return fmt.Errorf("JSON input is required if no file specified")
	}

	if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
		return err
	}

	if inputFile != "" {
		fmt.Println("ignoring input file as JSON input is provided")
	}

	// Validate the JSON input
	if json.Valid([]byte(args[0])) {
		input = json.RawMessage(args[0])
		return nil
	}

	return fmt.Errorf("invalid JSON provided: %s", args[0])
}

/*
 * Validate the flags passed to the command.
 */
func validateFlags() error {
	if err := validateOutputFormat(); err != nil {
		return err
	}

	return nil
}

/*
 * Validate the output format flag.
 */
func validateOutputFormat() error {
	if rawOutputFormat == "" {
		return fmt.Errorf("output format is required")
	}

	switch OutputFormat(rawOutputFormat) {
	case JSON, YAML, XML, TS, GO, RS:
		outputFormat = OutputFormat(rawOutputFormat)
		return nil
	default:
		return fmt.Errorf("invalid output format specified: %s", rawOutputFormat)
	}
}

/*
 * Convert JSON to the specified output format.
 */
func ConvertJSON() error {
	fmt.Printf("Converting %s to %s\n", input, outputFormat)
	return nil
}
