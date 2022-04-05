package fs

import (
	"os"
	"path"
	"strings"

	"github.com/go-zoox/uuid"
)

func CurrentDir() string {
	if dir, err := os.Getwd(); err == nil {
		return dir
	}

	return ""
}

func JoinPath(paths ...string) string {
	return path.Join(paths...)
}

func BaseName(path_ string) string {
	return path.Base(path_)
}

func DirName(path_ string) string {
	return path.Dir(path_)
}

func ExtName(path_ string) string {
	return path.Ext(path_)
}

func IsAbsPath(path_ string) bool {
	return path.IsAbs(path_)
}

func TmpDirPath() string {
	return os.TempDir()
}

func TmpFilePath() string {
	return strings.Join([]string{TmpDirPath(), uuid.V4()}, "/")
}
