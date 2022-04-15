package fs

import (
	"os"
	ospath "path"
	"strings"

	"github.com/go-zoox/uuid"
)

// JoinPath joins paths into a path.
func JoinPath(paths ...string) string {
	return ospath.Join(paths...)
}

// BaseName returns the last element of path.
func BaseName(path string) string {
	return ospath.Base(path)
}

// DirName returns all but the last element of path.
func DirName(path string) string {
	return ospath.Dir(path)
}

// ExtName returns the file extension of path.
func ExtName(path string) string {
	return ospath.Ext(path)
}

// IsAbsPath checks whether the path is absolute.
func IsAbsPath(path string) bool {
	return ospath.IsAbs(path)
}

// TmpDirPath returns the path of the temporary directory.
func TmpDirPath() string {
	return os.TempDir()
}

// TmpFilePath returns the path of the temporary file.
func TmpFilePath() string {
	return strings.Join([]string{TmpDirPath(), uuid.V4()}, "/")
}
