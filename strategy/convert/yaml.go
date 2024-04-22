package convert

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/dtran421/json-wizard/types/convert"
	"github.com/dtran421/json-wizard/utils"
)

type YAMLConverter struct {
	input     json.RawMessage
	inputFile string
}

func (c *YAMLConverter) SetInput(input json.RawMessage) {
	c.input = input
}

func (c *YAMLConverter) SetInputFile(inputFile string) {
	c.inputFile = inputFile
}

func (c *YAMLConverter) Convert() error {
	outputPath := filepath.Join(".", "output", string(convert.YAML))

	if err := os.MkdirAll(outputPath, os.ModePerm); err != nil {
		return err
	}

	fo, err := os.Create(fmt.Sprintf("output.%s", string(convert.YAML.GetExtension())))
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

	if err := convertToYAML(fo, output); err != nil {
		return err
	}

	return nil
}

func convertToYAML(fo *os.File, output interface{}) error {
	fo.WriteString("---\n")

	// TODO: handle array of objects

	for _, keyValuePair := range utils.SortedMap(output.(map[string]interface{})) {
		k, v := keyValuePair.Key, keyValuePair.Value
		if err := convertToYAMLHelper(fo, k, v, 0, -1); err != nil {
			return err
		}
	}

	return nil
}

func convertToYAMLHelper(fo *os.File, key string, value interface{}, level int, idx int) error {
	indent := getIndent(level)

	switch value := value.(type) {
	case map[string]interface{}:
		outputKey := key
		if key == "-" {
			outputKey = fmt.Sprintf("- %d", idx)
		}

		fo.WriteString(fmt.Sprintf("%s%s:\n", indent, outputKey))

		for _, keyValuePair := range utils.SortedMap(value) {
			k, v := keyValuePair.Key, keyValuePair.Value
			convertToYAMLHelper(fo, k, v, level+1, -1)
		}

	case []interface{}:
		fo.WriteString(fmt.Sprintf("%s%s:\n", indent, key))

		for idx, v := range value {
			convertToYAMLHelper(fo, "-", v, level+1, idx)
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

func getIndent(level int) string {
	indent := ""
	for i := 0; i < level; i++ {
		indent += " "
	}

	return indent
}
