package yaml_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/dtran421/json-wizard/strategy/converter/yaml"
	"github.com/dtran421/json-wizard/types"
	"github.com/dtran421/json-wizard/utils"
)

var conv yaml.YAMLConverter

var testFilepath types.Filepath = types.NewFilepath("/strategy/converter/yaml/").WithPrefix(utils.TestPathname())
var outputFilepath types.Filepath = types.NewFilepath("/output/yaml/output.yaml").WithPrefix(utils.TestPathname())

func setupTest() {
	conv = *yaml.NewYAMLConverter()
	conv.SetOutputFile(outputFilepath)
	conv.SetIndentSize(2)
}

func TestNewYAMLConverter(t *testing.T) {
	var conv = yaml.NewYAMLConverter()

	if conv == nil {
		t.Errorf("NewYAMLConverter() == nil, want !nil")
	}
}

func TestConvert_Input_HappyPath(t *testing.T) {
	cases := []struct {
		in               types.Filepath
		expectedFilepath types.Filepath
	}{
		{
			in:               testFilepath.Append("test1_input.json"),
			expectedFilepath: testFilepath.Append("test1_expected.yaml"),
		},
		{
			in:               testFilepath.Append("test2_input.json"),
			expectedFilepath: testFilepath.Append("test2_expected.yaml"),
		},
	}

	for _, c := range cases {
		setupTest()

		var input, inputErr = os.Open(c.in.String())
		if inputErr != nil {
			t.Errorf("Convert(%q): %s", c.in, inputErr)
			continue
		}
		defer input.Close()

		var inputScanner = bufio.NewScanner(input)

		var inputString string
		for inputScanner.Scan() {
			inputString += inputScanner.Text()
		}

		conv.SetInput([]byte(inputString))

		if err := conv.Convert(); err != nil {
			t.Errorf("Convert(%q) == %q, want nil", c.in, err)
		}

		var expectedScanner, outputScanner = utils.OutputAndExpectedScanners(t, c.in.String(), outputFilepath, c.expectedFilepath)

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

func TestConvert_InputFile_HappyPath(t *testing.T) {
	t.Skip("TODO: need to implement this function")

	cases := []struct {
		inputFile        types.Filepath
		expectedFilepath types.Filepath
	}{
		{
			inputFile:        testFilepath.Append("test1_input.json"),
			expectedFilepath: testFilepath.Append("test1_expected.yaml"),
		},
		{
			inputFile:        testFilepath.Append("test2_input.json"),
			expectedFilepath: testFilepath.Append("test2_expected.yaml"),
		},
	}

	for _, c := range cases {
		setupTest()

		conv.SetInputFile(c.inputFile)

		if err := conv.Convert(); err != nil {
			t.Errorf("Convert(%q) == %q, want nil", c.inputFile, err)
		}

		var expectedScanner, outputScanner = utils.OutputAndExpectedScanners(t, c.inputFile.String(), outputFilepath, c.expectedFilepath)

		for {
			var expectedScan, outputScan = expectedScanner.Scan(), outputScanner.Scan()
			if !expectedScan || !outputScan {
				if expectedScan != outputScan {
					t.Errorf("Convert(%q): output and expected files have different number of lines", c.inputFile)
				}

				break
			}

			outputLine := outputScanner.Text()
			expectedLine := expectedScanner.Text()

			if outputLine != expectedLine {
				t.Errorf("Convert(%q): \noutput:\n\t%s\ndoes not match expected:\n\t%s",
					c.inputFile, outputLine, expectedLine)
				break
			}
		}
	}
}
