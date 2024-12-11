package utils

import "os"

func GetIndent(level int) string {
	indent := ""
	for i := 0; i < level; i++ {
		indent += " "
	}

	return indent
}

func GetCustomIndent(level int, indentSize int) string {
	indent := ""
	for i := 0; i < level*indentSize; i++ {
		indent += " "
	}

	return indent
}

func CreateTempFile() (*os.File, error) {
	ft, err := os.CreateTemp("", "json-wizard_temp")
	if err != nil {
		return nil, err
	}

	return ft, nil
}

func RemoveTempFile(ft *os.File) error {
	if err := ft.Close(); err != nil {
		return err
	}

	if err := os.Remove(ft.Name()); err != nil {
		return err
	}

	return nil
}
