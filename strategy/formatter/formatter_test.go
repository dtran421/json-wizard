package formatter_test

import (
	"testing"

	"github.com/dtran421/json-wizard/providers/format"
	"github.com/dtran421/json-wizard/providers/iofile_validator"
	formatStrategy "github.com/dtran421/json-wizard/strategy/formatter"
	"github.com/dtran421/json-wizard/types"
	"github.com/dtran421/json-wizard/utils"
)

var outputFilepath = utils.TestPathname().Append("/output/json/output.json")

var formatCmd format.FormatCmd
var formatter formatStrategy.Formatter

func setupTest() {
	ioFileValidator := iofile_validator.IOFileValidator{}
	formatCmd = *format.New(ioFileValidator)

	formatter = *formatStrategy.NewFormatter()
}

func TestNewFormatter(t *testing.T) {
	if formatStrategy.NewFormatter() == nil {
		t.Error("NewFormatter() == nil, want &Formatter{}")
	}
}

func TestFormat_WithInput_HappyPath(t *testing.T) {
	cases := []struct {
		in               string
		expectedFilepath types.Filepath
	}{
		{
			in:               `{"key": "value"}`,
			expectedFilepath: utils.TestPathname().Append("/strategy/formatter/json/one_level.json"),
		},
		{
			in:               `{"key": "value", "key2": "value2"}`,
			expectedFilepath: utils.TestPathname().Append("/strategy/formatter/json/multiple_keys.json"),
		},
		{
			in:               `{"key": {"key2": "value"}}`,
			expectedFilepath: utils.TestPathname().Append("/strategy/formatter/json/two_levels.json"),
		},
	}

	for _, c := range cases {
		setupTest()

		formatCmd.SetInput([]byte(c.in))
		formatCmd.SetOutputFile(outputFilepath.String())
		formatCmd.SetIndentSize(2)

		if err := formatter.Format(&formatCmd); err != nil {
			t.Errorf("Format() == %v, want nil", err)
		}

		var expectedScanner, outputScanner = utils.OutputAndExpectedScanners(t, c.in, outputFilepath, c.expectedFilepath)

		for {
			var expectedScan, outputScan = expectedScanner.Scan(), outputScanner.Scan()
			if !expectedScan || !outputScan {
				if expectedScan != outputScan {
					t.Errorf("Convert(%q): output and expected files have different number of lines", c.in)
				}

				break
			}

			outputLine := outputScanner.Text()
			expectedLine := expectedScanner.Text()

			if outputLine != expectedLine {
				t.Errorf("Convert(%q): \noutput:\n\t%s\ndoes not match expected:\n\t%s",
					c.in, outputLine, expectedLine)
				break
			}
		}
	}
}
