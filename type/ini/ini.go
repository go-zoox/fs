package ini

import (
	"fmt"

	"github.com/go-zoox/fs"
	goini "github.com/subpop/go-ini"
)

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
