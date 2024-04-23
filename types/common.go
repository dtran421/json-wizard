package types

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func HomeDirpath() Filepath {
	return NewFilepath("/")
}

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
	pathname = strings.TrimSuffix(pathname, "/")

	dirname, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("error getting user home directory: %v", err))
	}

	if pathname == dirname {
		return Filepath(dirname)
	}

	if strings.HasPrefix(pathname, dirname) {
		return Filepath(pathname)
	}

	return Filepath(filepath.Join(dirname, pathname))
}

func NewFilepathFromAbsPath(pathname string) Filepath {
	pathname = strings.TrimSuffix(pathname, "/")

	if pathname == "" {
		return Filepath("/")
	}

	if !strings.HasPrefix(pathname, "/") {
		return Filepath(fmt.Sprintf("/%s", pathname))
	}

	return Filepath(pathname)
}

func (f Filepath) WithPrefix(prefix Filepath) Filepath {
	if prefix.IsEmpty() {
		return f
	}

	if f.HasPrefix(prefix) {
		return f
	}

	pathname := strings.TrimPrefix(f.String(), HomeDirpath().String())

	return NewFilepathFromAbsPath(fmt.Sprintf("%s%s", prefix.String(), pathname))
}

func (f Filepath) WithExtension(format OutputFormat) Filepath {
	if f.Extension() == string(format.GetExtension()) {
		return f
	}

	return NewFilepath(fmt.Sprintf("%s%s", f.String(), string(format.GetExtension())))
}

func (f Filepath) Append(pathname string) Filepath {
	if f.IsEmpty() {
		return NewFilepath(pathname)
	}

	formattedFilepath := strings.TrimPrefix(pathname, HomeDirpath().String())

	return NewFilepath(fmt.Sprintf("%s%s", f.String(), NewFilepathFromAbsPath(formattedFilepath).String()))
}

func (f Filepath) Directory() string {
	return filepath.Dir(f.String())
}

func (f Filepath) Base() string {
	if f.IsEmpty() {
		return ""
	}

	base := filepath.Base(f.String())

	return base
}

func (f Filepath) Extension() string {
	return filepath.Ext(f.String())
}

func (f Filepath) HasPrefix(prefix Filepath) bool {
	return strings.HasPrefix(f.String(), prefix.String())
}

func (f Filepath) IsAtHomeDir() bool {
	return f.HasPrefix(HomeDirpath()) && f == HomeDirpath()
}

func (f Filepath) IsEmpty() bool {
	return f == "/"
}

func (f Filepath) String() string {
	return string(f)
}
