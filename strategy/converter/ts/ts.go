package ts

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/dtran421/json-wizard/providers"
	"github.com/dtran421/json-wizard/types"
	"github.com/dtran421/json-wizard/utils"
)

type TSConverter struct {
	quicktypeWrapper providers.QuicktypeWrapper

	input      json.RawMessage
	inputFile  types.Filepath
	outputFile types.Filepath
	indentSize int
}

func NewTSConverter() *TSConverter {
	return &TSConverter{
		quicktypeWrapper: providers.NewQuicktypeWrapper(types.JSON, types.TS),
	}
}

func (c *TSConverter) SetInput(input json.RawMessage) {
	c.input = input
}

func (c *TSConverter) SetInputFile(inputFile types.Filepath) {
	c.inputFile = inputFile
}

func (c *TSConverter) SetOutputFile(outputFile types.Filepath) {
	c.outputFile = outputFile
}

func (c *TSConverter) SetIndentSize(indentSize int) {
	c.indentSize = indentSize
}

func (c *TSConverter) Convert() error {
	outputFilepath := c.outputFile.WithExtension(types.TS)

	if err := os.MkdirAll(filepath.Dir(outputFilepath.String()), os.ModePerm); err != nil {
		return err
	}

	fo, err := os.Create(outputFilepath.String())
	if err != nil {
		return err
	}

	defer func() error {
		if err := fo.Close(); err != nil {
			return err
		}

		return nil
	}()

	c.quicktypeWrapper.SetOutFile(outputFilepath)

	var output interface{}
	if err := json.Unmarshal(c.input, &output); err != nil {
		return err
	}

	ft, err := utils.CreateTempFile()
	if err != nil {
		return err
	}

	defer utils.RemoveTempFile(ft)

	if _, err := ft.Write(c.input); err != nil {
		return err
	}

	c.quicktypeWrapper.SetInputFile(ft)

	// c.quicktypeWrapper

	c.quicktypeWrapper.Execute()

	return nil
}
