package utils_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/dtran421/json-wizard/types"
	"github.com/dtran421/json-wizard/utils"
)

func TestRootpath(t *testing.T) {
	Rootpath := utils.Rootpath()

	pathname, _ := os.Getwd()
	expected := types.NewFilepath(types.NewFilepath(pathname).Directory())

	if Rootpath != expected {
		t.Errorf("Rootpath() == %q, want %q", Rootpath, expected)
	}
}

func TestTestPathname(t *testing.T) {
	testpath := utils.TestPathname()

	pathname, _ := os.Getwd()
	expected := types.NewFilepath(fmt.Sprintf("%s/test", filepath.Dir(pathname)))

	if testpath != expected {
		t.Errorf("TestPathname() == %q, want %q", testpath, expected)
	}
}

func TestOutputAndExpectedScanners(t *testing.T) {
	input := `{"key": "value"}`
	outputFilepath := utils.TestPathname().Append("/output/json/output.json")
	expectedFilepath := utils.TestPathname().Append("/strategy/formatter/json/test1_expected.json")

	expectedScanner, outputScanner := utils.OutputAndExpectedScanners(t, input, outputFilepath, expectedFilepath)

	if expectedScanner == nil {
		t.Errorf("OutputAndExpectedScanners() == %v, want !nil", expectedScanner)
	}

	if outputScanner == nil {
		t.Errorf("OutputAndExpectedScanners() == %v, want !nil", outputScanner)
	}
}
