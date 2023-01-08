package fs

import (
	"fmt"
	"os"
	ospath "path"
	"strings"

	"github.com/go-zoox/uuid"
)

// JoinPath joins any number of path elements into a single path, adding a
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

// CurrentDir returns the path of the current directory.
func CurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("cannot get current dir with os.Getwd(): %v", err))
	}

	return dir
}

// HomeDir returns the user home directory.
func HomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Sprintf("cannot get user home dir with os.UserHomeDir(): %v", err))
	}

	return homeDir
}

// ConfigDir returns the user config directory, which is $HOME/.config
func ConfigDir() string {
	return JoinPath(HomeDir(), ".config")
}

// JoinCurrentDir returns the path which relative with current dir.
func JoinCurrentDir(relativePath string) string {
	return JoinPath(CurrentDir(), relativePath)
}

// JoinHomeDir returns the path which relative with homedir.
func JoinHomeDir(relativePath string) string {
	return JoinPath(HomeDir(), relativePath)
}

// JoinConfigDir returns the path which relative with user home config dir.
func JoinConfigDir(relativePath string) string {
	return JoinPath(ConfigDir(), relativePath)
}
