package fs

import (
	iofs "io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// CreateDir creates a directory.
func CreateDir(path string) error {
	return os.MkdirAll(path, 0755)
}

// RemoveDir removes a directory.
func RemoveDir(path string) error {
	return os.RemoveAll(path)
}

// RenameDir renames a directory.
func RenameDir(srcPath string, dstPath string) error {
	return os.Rename(srcPath, dstPath)
}

// MoveDir moves a directory.
func MoveDir(srcPath string, dstPath string) error {
	return os.Rename(srcPath, dstPath)
}

// ListDir lists the files in a directory.
func ListDir(path string) ([]iofs.FileInfo, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	return files, nil
}

// WalkDir walks the files in a directory.
func WalkDir(path string, fn iofs.WalkDirFunc) error {
	return filepath.WalkDir(path, fn)
}

// CopyDir copies a directory.
func CopyDir(srcPath string, dstPath string) error {
	if !IsExist(dstPath) {
		Mkdirp(dstPath)
	}

	return WalkDir(srcPath, func(path string, dir os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if dir.IsDir() {
			dstPath := strings.Replace(path, srcPath, dstPath, 1)
			if IsExist(dstPath) {
				return nil
			}

			return CreateDir(dstPath)
		}

		return CopyFile(path, strings.Replace(path, srcPath, dstPath, 1))
	})
}
