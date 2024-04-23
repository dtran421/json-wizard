package xml_test

import (
	"bufio"
	"os"
	"testing"

	convert "github.com/dtran421/json-wizard/strategy/convert/xml"
	"github.com/dtran421/json-wizard/types"
	"github.com/dtran421/json-wizard/utils"
)

var conv convert.XMLConverter

var testFilepath types.Filepath = types.NewFilepath("/strategy/convert/xml/").WithPrefix(utils.TestPathname())
var outputFilepath types.Filepath = types.NewFilepath("/output/xml/output.xml").WithPrefix(utils.TestPathname())

func setupTest() {
	conv = convert.XMLConverter{}
	conv.SetOutputFile(outputFilepath)
	conv.SetIndentSize(2)
}

func teardownTest(t *testing.T, in types.Filepath, openTestOutputFiles ...*os.File) {
	for _, openTestOutputFile := range openTestOutputFiles {
		if err := openTestOutputFile.Close(); err != nil {
			t.Errorf("Convert(%s): %s", in.String(), err)
		}
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

		var expectedScanner, outputScanner = getOutputAndExpectedScanners(t, c.in, outputFilepath, c.expectedFilepath)

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

		var input, inputErr = os.Open(c.inputFile.String())
		if inputErr != nil {
			t.Errorf("Convert(%q): %s", c.inputFile, inputErr)
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
			t.Errorf("Convert(%q) == %q, want nil", c.inputFile, err)
		}

		var expectedScanner, outputScanner = getOutputAndExpectedScanners(t, c.inputFile, outputFilepath, c.expectedFilepath)

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

func getOutputAndExpectedScanners(
	t *testing.T,
	input types.Filepath,
	outputFilepath types.Filepath,
	expectedFilepath types.Filepath,
) (*bufio.Scanner, *bufio.Scanner) {
	var outputFile, outputErr = os.Open(outputFilepath.String())
	if outputErr != nil {
		t.Errorf("Convert(%q): %s", input, outputErr)
	}
	defer teardownTest(t, input, outputFile)

	var expectedFile, expectedErr = os.Open(expectedFilepath.String())
	if expectedErr != nil {
		t.Errorf("Convert(%q): %s", input, expectedErr)
	}
	defer expectedFile.Close()

	return bufio.NewScanner(expectedFile), bufio.NewScanner(outputFile)
}
