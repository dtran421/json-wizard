package formatter

import (
	"encoding/json"
	"os"

	"github.com/dtran421/json-wizard/types"
	"github.com/dtran421/json-wizard/utils"
)

type Formatter struct{}

func NewFormatter() *Formatter {
	return &Formatter{}
}

func (f Formatter) Format(cmd types.Cmd) error {
	fo, err := os.Create(cmd.OutputFile().String())
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
	if err := json.Unmarshal(cmd.Input(), &output); err != nil {
		return err
	}

	indent := utils.GetCustomIndent(1, cmd.IndentSize())

	formattedOutput, err := json.MarshalIndent(output, "", indent)
	if err != nil {
		return err
	}

	if _, err := fo.Write(formattedOutput); err != nil {
		return err
	}

	return nil
}
