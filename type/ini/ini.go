package ini

import (
	"fmt"

	goini "github.com/go-ini/ini"
	"github.com/go-zoox/fs"
)

func Read(path string) (*goini.File, error) {
	if path == "" {
		return nil, fmt.Errorf("path is empty")
	}
	if !fs.IsExist(path) {
		return nil, fmt.Errorf("path is not exist")
	}

	// str, err := fs.ReadFileAsString(path)
	// if err != nil {
	// 	return err
	// }

	return goini.Load(path)
}

func Write(path string, data interface{}) error {
	panic("@TO_IMPLEMENT")
	// if path == "" {
	// 	return fmt.Errorf("path is empty")
	// }

	// str, err := goini.Marshal(data)
	// if err != nil {
	// 	return err
	// }

	// return fs.WriteFile(path, str)
}
