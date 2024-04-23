package yaml

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/dtran421/json-wizard/types"
	"github.com/dtran421/json-wizard/utils"
)

type YAMLConverter struct {
	input      json.RawMessage
	inputFile  types.Filepath
	outputFile types.Filepath
	indentSize int
}

func NewYAMLConverter() *YAMLConverter {
	return &YAMLConverter{}
}

func (c *YAMLConverter) SetInput(input json.RawMessage) {
	c.input = input
}

func (c *YAMLConverter) SetInputFile(inputFile types.Filepath) {
	c.inputFile = inputFile
}

func (c *YAMLConverter) SetOutputFile(outputFile types.Filepath) {
	c.outputFile = outputFile
}

func (c *YAMLConverter) SetIndentSize(indentSize int) {
	c.indentSize = indentSize
}

func (c *YAMLConverter) Convert() error {
	outputFilepath := c.outputFile.WithExtension(types.YAML)

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

	if c.input != nil {
		if err := c.convertInputToYAML(fo); err != nil {
			return err
		}
	} else {
		if err := c.convertInputFileToYAML(fo); err != nil {
			return err
		}
	}

	return nil
}

// TODO: need to implement this function
func (c YAMLConverter) convertInputFileToYAML(fo *os.File) error {
	fi, err := os.Open(c.inputFile.String())
	if err != nil {
		return err
	}
	defer fi.Close()

	// TODO: format the file

	// TODO: convert the file to YAML line by line

	fo.WriteString("---\n")

	return nil
}

func (c YAMLConverter) convertInputToYAML(fo *os.File) error {
	var output interface{}
	if err := json.Unmarshal(c.input, &output); err != nil {
		return err
	}

	fo.WriteString("---\n")

	// TODO: handle array of objects

	for _, keyValuePair := range utils.SortedMap(output.(map[string]interface{})) {
		k, v := keyValuePair.Key, keyValuePair.Value
		c.convertToYAMLHelper(fo, k, v, 0, -1)
	}

	return nil
}

func (c YAMLConverter) convertToYAMLHelper(fo *os.File, key string, value interface{}, level int, idx int) {
	indent := utils.GetCustomIndent(level, c.indentSize)

	switch value := value.(type) {
	case map[string]interface{}:
		outputKey := key
		if key == "-" {
			outputKey = fmt.Sprintf("%s- %d", indent, idx)
		}

		fo.WriteString(fmt.Sprintf("%s%s:\n", indent, outputKey))

		for _, keyValuePair := range utils.SortedMap(value) {
			k, v := keyValuePair.Key, keyValuePair.Value
			c.convertToYAMLHelper(fo, k, v, level+1, -1)
		}

	case []interface{}:
		fo.WriteString(fmt.Sprintf("%s%s:\n", indent, key))

		for idx, v := range value {
			c.convertToYAMLHelper(fo, "-", v, level+1, idx)
		}

	default:
		outputValue := value
		if reflect.TypeOf(outputValue).String() == "string" {
			outputValue = fmt.Sprintf("\"%s\"", value)
		}

		if key == "-" {
			fo.WriteString(fmt.Sprintf("%s- %v\n", indent, outputValue))
			return
		}

		fo.WriteString(fmt.Sprintf("%s%s: %v\n", indent, key, outputValue))
	}
}
