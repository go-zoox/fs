package fs

import (
	"os"
	"path"
	"strings"

	"github.com/go-zoox/uuid"
)

// JoinPath joins paths into a path.
func JoinPath(paths ...string) string {
	return path.Join(paths...)
}

// BaseName returns the last element of path.
func BaseName(path_ string) string {
	return path.Base(path_)
}

// DirName returns all but the last element of path.
func DirName(path_ string) string {
	return path.Dir(path_)
}

// ExtName returns the file extension of path.
func ExtName(path_ string) string {
	return path.Ext(path_)
}

// IsAbsPath checks whether the path is absolute.
func IsAbsPath(path_ string) bool {
	return path.IsAbs(path_)
}

// TmpDirPath returns the path of the temporary directory.
func TmpDirPath() string {
	return os.TempDir()
}

// TmpFilePath returns the path of the temporary file.
func TmpFilePath() string {
	return strings.Join([]string{TmpDirPath(), uuid.V4()}, "/")
}
