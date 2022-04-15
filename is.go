package fs

import (
	"os"
)

// IsExist checks whether a file or directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// IsFile checks whether the path is a file.
func IsFile(path string) bool {
	fi, err := os.Stat(path)
	return err == nil && fi.Mode().IsRegular()
}

// IsDir checks whether the path is a directory.
func IsDir(path string) bool {
	fi, err := os.Stat(path)
	return err == nil && fi.Mode().IsDir()
}

// IsLink checks whether the path is a symbolic link.
func IsLink(path string) bool {
	fi, err := os.Stat(path)
	return err == nil && fi.Mode()&os.ModeSymlink != 0
}

// IsEmpty checks whether the dir/file is empty.
func IsEmpty(path string) bool {
	if !IsFile(path) {
		if bytes, err := ReadFile(path); err != nil {
			panic(err)
		} else {
			return len(bytes) == 0
		}
	}

	files, err := ListDir(path)
	if err != nil {
		panic(err)
	}

	return len(files) == 0
}
