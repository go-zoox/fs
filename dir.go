package fs

import (
	iofs "io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
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
	return RenameFile(srcPath, dstPath)
}

// MoveDir moves a directory.
func MoveDir(srcPath string, dstPath string) error {
	return MoveFile(srcPath, dstPath)
}

// Mkdir creates a directory.
func Mkdir(path string) error {
	return CreateDir(path)
}

// Mkdirp creates a deep directory.
func Mkdirp(path string) error {
	return CreateDir(path)
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
			if IsExist(path) {
				return nil
			}

			return CreateDir(path)
		}

		return CopyFile(path, JoinPath(dstPath, dir.Name()))
	})
}
