package converter

import (
	"fmt"

	"github.com/dtran421/json-wizard/strategy/converter/ts"
	"github.com/dtran421/json-wizard/strategy/converter/yaml"
	"github.com/dtran421/json-wizard/types"
	"github.com/dtran421/json-wizard/types/convert"
)

type ConverterFactoryInstance struct{}

var factorySingleton ConverterFactoryInstance

func ConverterFactory() *ConverterFactoryInstance {
	if factorySingleton != (ConverterFactoryInstance{}) {
		factorySingleton = ConverterFactoryInstance{}
	}

	return &factorySingleton
}

func (c ConverterFactoryInstance) BuildConverter(outputFormat types.OutputFormat) (convert.ConvertStrategy, error) {
	switch outputFormat {
	case types.YAML:
		return yaml.NewYAMLConverter(), nil
	case types.TS:
		return ts.NewTSConverter(), nil
	case types.GO:
		fmt.Println("Converting to Go struct")
	case types.RS:
		fmt.Println("Converting to Rust struct")
	default:
		return nil, fmt.Errorf("invalid output format specified: %s", outputFormat)
	}

	return nil, nil
}
