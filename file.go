package fs

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
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
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Open(dstPath)
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
func OpenFile(path string) (*os.File, error) {
	return os.Open(path)
}

// WriteFile writes a file.
func WriteFile(path string, data []byte) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write(data)
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

// ReadFileByLine reads a file by line.
func ReadFileByLine(srcPath string) ([]string, error) {
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
