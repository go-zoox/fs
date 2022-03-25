package fs

import "os"

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func IsFile(path string) bool {
	fi, err := os.Stat(path)
	return err == nil && fi.Mode().IsRegular()
}

func IsDir(path string) bool {
	fi, err := os.Stat(path)
	return err == nil && fi.Mode().IsDir()
}

func IsLink(path string) bool {
	fi, err := os.Stat(path)
	return err == nil && fi.Mode()&os.ModeSymlink != 0
}
