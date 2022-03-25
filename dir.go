package fs

import "os"

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
