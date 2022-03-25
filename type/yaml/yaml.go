package yaml

import (
	"fmt"

	"github.com/go-zoox/fs"
	goyaml "github.com/goccy/go-yaml"
)

func Read(path string, data interface{}) error {
	if path == "" {
		return fmt.Errorf("path is empty")
	}
	if !fs.IsExist(path) {
		return fmt.Errorf("path is not exist")
	}

	str, err := fs.ReadFileAsString(path)
	if err != nil {
		return err
	}

	return goyaml.Unmarshal([]byte(str), data)
}

func Write(path string, data interface{}) error {
	if path == "" {
		return fmt.Errorf("path is empty")
	}

	str, err := goyaml.Marshal(data)
	if err != nil {
		return err
	}

	return fs.WriteFile(path, str)
}
