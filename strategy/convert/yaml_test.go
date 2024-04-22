package convert_test

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/dtran421/json-wizard/strategy/convert"
	"github.com/dtran421/json-wizard/types"
)

var conv convert.YAMLConverter

var outputFilepath types.Filepath = "output.yaml"

func setupTest() {
	conv = convert.YAMLConverter{}
}

func TestConvert_HappyPath(t *testing.T) {
	cases := []struct {
		in               string
		expectedFilepath types.Filepath
	}{
		{
			in:               "../../test/strategy/convert/test1_input.json",
			expectedFilepath: "../../test/strategy/convert/test1_expected.yaml",
		},
		{
			in:               "../../test/strategy/convert/test2_input.json",
			expectedFilepath: "../../test/strategy/convert/test2_expected.yaml",
		},
	}

	for _, c := range cases {
		setupTest()

		var input, inputErr = os.Open(string(c.in))
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

		fmt.Printf("inputString: %s\n", inputString)

		conv.SetInput([]byte(inputString))

		if err := conv.Convert(); err != nil {
			t.Errorf("Convert(%q) == %q, want nil", c.in, err)
		}

		var outputFile, outputErr = os.Open(string(outputFilepath))
		if outputErr != nil {
			t.Errorf("Convert(%q): %s", c.in, outputErr)
			continue
		}
		defer outputFile.Close()

		var expectedFile, expectedErr = os.Open(string(c.expectedFilepath))
		if expectedErr != nil {
			t.Errorf("Convert(%q): %s", c.in, expectedErr)
			continue
		}
		defer expectedFile.Close()

		expectedScanner := bufio.NewScanner(expectedFile)
		outputScanner := bufio.NewScanner(outputFile)

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
