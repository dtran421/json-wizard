package convert

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/dtran421/json-wizard/types"
	"github.com/dtran421/json-wizard/utils"
)

type XMLConverter struct {
	input json.RawMessage

	inputFile  types.Filepath
	outputFile types.Filepath

	indentSize int
}

func (c *XMLConverter) SetInput(input json.RawMessage) {
	c.input = input
}

func (c *XMLConverter) SetInputFile(inputFile types.Filepath) {
	c.inputFile = inputFile
}

func (c *XMLConverter) SetOutputFile(outputFile types.Filepath) {
	c.outputFile = outputFile
}

func (c *XMLConverter) SetIndentSize(indentSize int) {
	c.indentSize = indentSize
}

func (c *XMLConverter) Convert() error {
	outputPath := filepath.Join(".", "output", string(types.XML))

	if err := os.MkdirAll(outputPath, os.ModePerm); err != nil {
		return err
	}

	fo, err := os.Create(fmt.Sprintf("output%s", string(types.XML.GetExtension())))
	if err != nil {
		return err
	}

	defer func() error {
		if err := fo.Close(); err != nil {
			return err
		}

		return nil
	}()

	var output interface{}
	if err := json.Unmarshal(c.input, &output); err != nil {
		return err
	}

	if err := c.convertToXML(fo, output); err != nil {
		return err
	}

	return nil
}

func (c XMLConverter) convertToXML(fo *os.File, output interface{}) error {
	fo.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")

	// TODO: handle array of objects

	for _, keyValuePair := range utils.SortedMap(output.(map[string]interface{})) {
		k, v := keyValuePair.Key, keyValuePair.Value
		if err := c.convertToXMLHelper(fo, k, v, 0, -1); err != nil {
			return err
		}
	}

	return nil
}

func (c XMLConverter) convertToXMLHelper(fo *os.File, key string, value interface{}, level int, idx int) error {
	indent := utils.GetCustomIndent(level, c.indentSize)

	switch value := value.(type) {
	case map[string]interface{}:
		outputKey := key
		if key == "-" {
			outputKey = fmt.Sprintf("- %d", idx)
		}

		fo.WriteString(fmt.Sprintf("%s%s:\n", indent, outputKey))

		for _, keyValuePair := range utils.SortedMap(value) {
			k, v := keyValuePair.Key, keyValuePair.Value
			c.convertToXMLHelper(fo, k, v, level+1, -1)
		}

	case []interface{}:
		fo.WriteString(fmt.Sprintf("%s%s:\n", indent, key))

		for idx, v := range value {
			c.convertToXMLHelper(fo, "-", v, level+1, idx)
		}
	default:
		outputValue := value
		if reflect.TypeOf(outputValue).String() == "string" {
			outputValue = fmt.Sprintf("\"%v\"", value)
		}

		if key == "-" {
			fo.WriteString(fmt.Sprintf("%s- %v\n", indent, outputValue))
			return nil
		}

		fo.WriteString(fmt.Sprintf("%s%s: %v\n", indent, key, outputValue))
	}

	return nil
}
