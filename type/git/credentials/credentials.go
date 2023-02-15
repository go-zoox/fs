package credentials

import (
	"fmt"

	gocredentials "github.com/go-zoox/encoding/git/credentials"
	"github.com/go-zoox/fs"
)

// Read reads the file at the given path and parses it as Git Credentials into data.
func Read(path string) (map[string]*gocredentials.Item, error) {
	if path == "" {
		return nil, fmt.Errorf("path is empty")
	}
	if !fs.IsExist(path) {
		return nil, fmt.Errorf("path is not exist")
	}

	bytes, err := fs.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return gocredentials.Decode(bytes)
}

// Write writes the given data to the file at the given path.
func Write(path string, data map[string]*gocredentials.Item) error {
	if path == "" {
		return fmt.Errorf("path is empty")
	}

	str, err := gocredentials.Encode(data)
	if err != nil {
		return err
	}

	return fs.WriteFile(path, str)
}
