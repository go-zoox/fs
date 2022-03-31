package json

import (
	"fmt"

	osjson "github.com/go-zoox/encoding/json"

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

	return osjson.Decode([]byte(str), data)
}

func Write(path string, data interface{}) error {
	if path == "" {
		return fmt.Errorf("path is empty")
	}

	str, err := osjson.Encode(data)
	if err != nil {
		return err
	}

	return fs.WriteFile(path, str)
}
