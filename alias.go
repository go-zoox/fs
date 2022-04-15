package fs

import (
	iofs "io/fs"
	"os"
)

// Open opens a file.
func Open(path string) (*os.File, error) {
	return OpenFile(path)
}

// Copy copies a file or directory
func Copy(srcPath, dstPath string) error {
	if IsFile(srcPath) {
		return CopyFile(srcPath, dstPath)
	}

	return CopyDir(srcPath, dstPath)
}

// Move moves a file or directory
func Move(srcPath, dstPath string) error {
	if IsFile(srcPath) {
		return MoveFile(srcPath, dstPath)
	}

	return MoveDir(srcPath, dstPath)
}

// Remove removes a file or directory
func Remove(path string) error {
	if IsFile(path) {
		return RemoveFile(path)
	}

	return RemoveDir(path)
}

// Rename renames a file or directory
func Rename(srcPath, dstPath string) error {
	if IsFile(srcPath) {
		return RenameFile(srcPath, dstPath)
	}

	return RenameDir(srcPath, dstPath)
}

// Walk walks the files in a directory.
func Walk(path string, fn iofs.WalkDirFunc) error {
	return WalkDir(path, fn)
}
