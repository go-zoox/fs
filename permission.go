package fs

import "os"

// Chmod changes the mode of the named file to mode.
func Chmod(name string, mode os.FileMode) error {
	return os.Chmod(name, mode)
}

// Chown changes the numeric uid and gid of the named file.
func Chown(name string, uid, gid int) error {
	return os.Chown(name, uid, gid)
}
