package convert

import (
	"encoding/json"

	"github.com/dtran421/json-wizard/types"
)

type ConvertStrategy interface {
	SetInput(input json.RawMessage)

	SetInputFile(inputFile types.Filepath)

	SetOutputFile(outputFile types.Filepath)

	SetIndentSize(indentSize int)

	/*
	 * Convert JSON to the specified output format.
	 */
	Convert() error
}
