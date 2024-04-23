package utils

import (
	"fmt"
	"os"
	"strings"

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
