package fs

import (
	"fmt"
	"os"
	"os/user"
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

// ConfigDir returns the config dir by user
//
//	if user is root, return system config dir
//	if user is common user, return user home config dir
func ConfigDir() string {
	if user, err := user.Current(); err == nil && user.Uid == "0" {
		return SystemConfigDir()
	}

	return UserConfigDir()
}

// UserConfigDir returns the user config directory, which is $HOME/.config
func UserConfigDir() string {
	return JoinPath(HomeDir(), ".config")
}

// SystemConfigDir returns the system config directory, which is /etc
func SystemConfigDir() string {
	return "/etc"
}

// JoinCurrentDir returns the path which relative with current dir.
func JoinCurrentDir(relativePath string) string {
	return JoinPath(CurrentDir(), relativePath)
}

// JoinHomeDir returns the path which relative with homedir.
func JoinHomeDir(relativePath string) string {
	return JoinPath(HomeDir(), relativePath)
}

// JoinConfigDir returns the config of appName + configName.
// configName default is config.yml.
func JoinConfigDir(appName string, configName ...string) string {
	if len(configName) == 0 {
		configName = append(configName, "config.yml")
	}

	if configName[0] == "" {
		panic("config name cannot be empty")
	}

	return JoinPath(ConfigDir(), appName, configName[0])
}
