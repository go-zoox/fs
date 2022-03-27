package fs

import (
	iofs "io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CreateDir(path string) error {
	return os.MkdirAll(path, 0755)
}

func RemoveDir(path string) error {
	return os.RemoveAll(path)
}

func RenameDir(srcPath string, dstPath string) error {
	return RenameFile(srcPath, dstPath)
}

func MoveDir(srcPath string, dstPath string) error {
	return MoveFile(srcPath, dstPath)
}

func Mkdir(path string) error {
	return CreateDir(path)
}

func Mkdirp(path string) error {
	return CreateDir(path)
}

func ListDir(path string) ([]iofs.FileInfo, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func WalkDir(path string, fn iofs.WalkDirFunc) error {
	return filepath.WalkDir(path, fn)
}
