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

var testFilepath types.Filepath = types.NewFilepath("../../test/strategy/convert/yaml/")
var outputFilepath types.Filepath = types.NewFilepath("output.yaml")

func setupTest() {
	conv = convert.YAMLConverter{}
	conv.SetOutputFile(testFilepath.Append(outputFilepath))
}

func teardownTest(t *testing.T, in types.Filepath, openTestOutputFiles ...*os.File) {
	for _, openTestOutputFile := range openTestOutputFiles {
		if err := openTestOutputFile.Close(); err != nil {
			t.Errorf("Convert(%s): %s", in.String(), err)
		}

		// if err := os.Remove(openTestOutputFile.Name()); err != nil {
		// 	t.Errorf("Convert(%q): %s", in, err)
		// }
	}
}

func TestConvert_HappyPath(t *testing.T) {
	cases := []struct {
		in               types.Filepath
		expectedFilepath types.Filepath
	}{
		{
			in:               types.NewFilepath("test1_input.json"),
			expectedFilepath: types.NewFilepath("test1_expected.yaml"),
		},
		{
			in:               types.NewFilepath("test2_input.json"),
			expectedFilepath: types.NewFilepath("test2_expected.yaml"),
		},
	}

	for _, c := range cases {
		setupTest()

		var input, inputErr = os.Open(testFilepath.Append(c.in).String())
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

		var outputFile, outputErr = os.Create(testFilepath.Append(outputFilepath).String())
		if outputErr != nil {
			t.Errorf("Convert(%q): %s", c.in, outputErr)
			continue
		}
		defer teardownTest(t, c.in, outputFile)

		var expectedFile, expectedErr = os.Open(testFilepath.Append(c.expectedFilepath).String())
		if expectedErr != nil {
			t.Errorf("Convert(%q): %s", c.in, expectedErr)
			continue
		}
		defer expectedFile.Close()

		expectedScanner := bufio.NewScanner(expectedFile)
		outputScanner := bufio.NewScanner(outputFile)

		fmt.Println("outputFile.Name():", outputFile.Name())

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
