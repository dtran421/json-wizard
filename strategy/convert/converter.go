package convert

import (
	"fmt"

	"github.com/dtran421/json-wizard/strategy/convert/xml"
	"github.com/dtran421/json-wizard/strategy/convert/yaml"
	"github.com/dtran421/json-wizard/types"
	"github.com/dtran421/json-wizard/types/convert"
)

type ConverterFactory struct{}

func NewConverterFactory() *ConverterFactory {
	return &ConverterFactory{}
}

func (c ConverterFactory) BuildConverter(outputFormat types.OutputFormat) convert.ConvertStrategy {
	switch outputFormat {
	case types.YAML:
		return &yaml.YAMLConverter{}
	case types.XML:
		return &xml.XMLConverter{}
	case types.TS:
		fmt.Println("Converting to TypeScript")
	case types.GO:
		fmt.Println("Converting to Go struct")
	case types.RS:
		fmt.Println("Converting to Rust struct")
	default:
		panic(fmt.Sprintf("invalid output format specified: %s", outputFormat))
	}

	return nil
}
