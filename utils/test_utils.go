package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/dtran421/json-wizard/types"
)

func Rootpath() types.Filepath {
	pathname, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("error getting project root directory: %v", err))
	}

	projectRootPathIdx := strings.Index(pathname, "json-wizard")
	rootpath := types.NewFilepathFromAbsPath(pathname[:projectRootPathIdx+len("json-wizard")])

	return rootpath
}

func TestPathname() types.Filepath {
	Rootpath := Rootpath()

	return Rootpath.Append("test")
}

func OutputAndExpectedScanners(
	t *testing.T,
	input string,
	outputFilepath types.Filepath,
	expectedFilepath types.Filepath,
) (*bufio.Scanner, *bufio.Scanner) {
	var outputFile, outputErr = os.Open(outputFilepath.String())
	if outputErr != nil {
		t.Errorf("Convert(%q): %s", input, outputErr)
	}
	defer closeTestFiles(t, input, outputFile)

	var expectedFile, expectedErr = os.Open(expectedFilepath.String())
	if expectedErr != nil {
		t.Errorf("Convert(%q): %s", input, expectedErr)
	}
	defer expectedFile.Close()

	return bufio.NewScanner(expectedFile), bufio.NewScanner(outputFile)
}

func closeTestFiles(t *testing.T, in string, openTestOutputFiles ...*os.File) {
	for _, openTestOutputFile := range openTestOutputFiles {
		if err := openTestOutputFile.Close(); err != nil {
			t.Errorf("Convert(%s): %s", in, err)
		}
	}
}
