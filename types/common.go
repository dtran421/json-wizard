package types

import (
	"fmt"
	"path/filepath"
	"strings"
)

type OutputFormat string

const (
	YAML OutputFormat = "yaml"
	XML  OutputFormat = "xml"
	TS   OutputFormat = "ts"
	GO   OutputFormat = "go"
	RS   OutputFormat = "rs"
)

type Extension string

func (f OutputFormat) GetExtension() Extension {
	switch f {
	case YAML:
		return ".yaml"
	case XML:
		return ".xml"
	case TS:
		return ".ts"
	case GO:
		return ".go"
	case RS:
		return ".rs"
	default:
		return ""
	}
}

type Filepath string

func NewFilepath(pathname string) Filepath {
	return Filepath(filepath.Join(".", pathname))
}

func (f Filepath) WithExtension(format OutputFormat) Filepath {
	// remove ending slash if it exists
	pathname := strings.TrimSuffix(string(f), "/")

	// remove the extension if it exists
	pathname = strings.TrimSuffix(pathname, string(format.GetExtension()))

	return Filepath(fmt.Sprintf("%s%s", pathname, string(format.GetExtension())))
}

func (f Filepath) Append(filepath Filepath) Filepath {
	pathname := strings.TrimSuffix(string(f), "/")

	return Filepath(fmt.Sprintf("%s/%s", pathname, filepath))
}

func (f Filepath) Directory() string {
	return filepath.Dir(string(f))
}

func (f Filepath) Base() string {
	base := filepath.Base(string(f))

	if base == "." {
		return ""
	}

	return base
}

func (f Filepath) Extension() string {
	return filepath.Ext(string(f))
}

func (f Filepath) IsEmpty() bool {
	return f == "."
}

func (f Filepath) String() string {
	return string(f)
}
