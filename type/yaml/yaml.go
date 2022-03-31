package yaml

import (
	"fmt"

	goyaml "github.com/go-zoox/encoding/yaml"
	"github.com/go-zoox/fs"
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

	return goyaml.Decode([]byte(str), data)
}

func Write(path string, data interface{}) error {
	if path == "" {
		return fmt.Errorf("path is empty")
	}

	str, err := goyaml.Encode(data)
	if err != nil {
		return err
	}

	return fs.WriteFile(path, str)
}
