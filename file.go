package fs

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
)

// CreateFile creates a file.
func CreateFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()
	return nil
}

// RemoveFile removes a file.
func RemoveFile(path string) error {
	return os.Remove(path)
}

// CopyFile copies a file.
func CopyFile(srcPath string, dstPath string) error {
	srcStat, err := os.Stat(srcPath)
	if err != nil {
		return err
	}

	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE, srcStat.Mode())
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

// RenameFile renames a file.
func RenameFile(srcPath, dstPath string) error {
	return os.Rename(srcPath, dstPath)
}

// MoveFile moves a file.
func MoveFile(srcPath, dstPath string) error {
	return RenameFile(srcPath, dstPath)
}

// OpenFile opens a file.
func OpenFile(path string, flagAndPerm ...interface{}) (*os.File, error) {
	//  according os.Open method
	var flag int = os.O_RDONLY
	var perm fs.FileMode = 0

	if len(flagAndPerm) > 0 {
		if flagAndPerm[0] != 0 {
			flagX, ok := flagAndPerm[0].(int)
			if !ok {
				return nil, fmt.Errorf("flag must be int, but got: %s(value: %v)", reflect.TypeOf(flagAndPerm[0]), flagAndPerm[0])
			}
			flag = flagX
		}

		if len(flagAndPerm) > 1 && flagAndPerm[1] != 0 {
			permX, ok := flagAndPerm[1].(fs.FileMode)
			if !ok {
				return nil, fmt.Errorf("flag must be uint32, but got: %s(value: %v)", reflect.TypeOf(flagAndPerm[1]), flagAndPerm[0])
			}
			perm = permX
		}
	}

	return os.OpenFile(path, flag, perm)
}

// WriteFile writes a file.
func WriteFile(path string, data []byte) (err error) {
	var f *os.File
	if !IsExist(path) {
		if err := CreateFile(path); err != nil {
			return err
		}
	} else {
		f, err = os.OpenFile(path, os.O_WRONLY, 0644)
		if err != nil {
			return err
		}

		// clean origin text
		if err := f.Truncate(0); err != nil {
			return err
		}
	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return err
	}

	return nil
}

// AppendFile appends a file.
func AppendFile(path string, data []byte) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	if _, err := f.Write(data); err != nil {
		return err
	}

	return nil
}

// ReadFile reads a file.
func ReadFile(srcPath string) ([]byte, error) {
	return ioutil.ReadFile(srcPath)
}

// ReadFileAsString reads a file as string.
func ReadFileAsString(srcPath string) (string, error) {
	bytes, err := ReadFile(srcPath)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// ReadFileLines reads a file by line.
func ReadFileLines(srcPath string) ([]string, error) {
	f, err := os.Open(srcPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	res := make([]string, 0)
	buf := bufio.NewReader(f)

	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		res = append(res, string(line))
	}

	return res, nil
}

// Stat returns the FileInfo structure describing file.
func Stat(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

// FilePart represents a part of a file.
type FilePart struct {
	Path string
	//
	Index int
}

// Merge merges files into one file.
func Merge(filePath string, parts []*FilePart) error {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	sort.Slice(parts, func(i, j int) bool {
		return parts[i].Index < parts[j].Index
	})

	for _, part := range parts {
		fp, err := os.Open(part.Path)
		if err != nil {
			return err
		}

		if _, err := io.Copy(f, fp); err != nil {
			return err
		}
	}

	return nil
}

// Size returns the size of the file.
func Size(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		return 0
	}

	return info.Size()
}
