package fs

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

func CreateFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()
	return nil
}

func RemoveFile(path string) error {
	return os.Remove(path)
}

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

func RenameFile(srcPath, dstPath string) error {
	return os.Rename(srcPath, dstPath)
}

func MoveFile(srcPath, dstPath string) error {
	return RenameFile(srcPath, dstPath)
}

func OpenFile(path string) (*os.File, error) {
	return os.Open(path)
}

func WriteFile(path string, data []byte) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write(data)
	return nil
}

func ReadFileAsString(srcPath string) (string, error) {
	bytes, err := ioutil.ReadFile(srcPath)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

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
