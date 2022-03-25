package toml

import (
	"fmt"

	"github.com/go-zoox/fs"
	gotoml "github.com/pelletier/go-toml"
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

	return gotoml.Unmarshal([]byte(str), data)
}

func Write(path string, data interface{}) error {
	if path == "" {
		return fmt.Errorf("path is empty")
	}

	str, err := gotoml.Marshal(data)
	if err != nil {
		return err
	}

	return fs.WriteFile(path, str)
}
