package toml

import (
	"fmt"

	gotoml "github.com/go-zoox/encoding/toml"
	"github.com/go-zoox/fs"
)

// Read reads the file at the given path and parses it as JSON into data.
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

	return gotoml.Decode([]byte(str), data)
}

// Write writes the given data to the file at the given path.
func Write(path string, data interface{}) error {
	if path == "" {
		return fmt.Errorf("path is empty")
	}

	str, err := gotoml.Encode(data)
	if err != nil {
		return err
	}

	return fs.WriteFile(path, str)
}
