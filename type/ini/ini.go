package ini

import (
	"fmt"

	"github.com/go-zoox/fs"
	goini "github.com/subpop/go-ini"
)

// Read reads the file at the given path and parses it as JSON into data.
func Read(path string, data interface{}) error {
	if path == "" {
		return fmt.Errorf("path is empty")
	}
	if !fs.IsExist(path) {
		return fmt.Errorf("path is not exist")
	}

	bytes, err := fs.ReadFile(path)
	if err != nil {
		return err
	}

	return goini.Unmarshal(bytes, data)
}

// Write writes the given data to the file at the given path.
func Write(path string, data interface{}) error {
	if path == "" {
		return fmt.Errorf("path is empty")
	}

	str, err := goini.Marshal(data)
	if err != nil {
		return err
	}

	return fs.WriteFile(path, str)
}
