package fs

import (
	"context"
	"fmt"
	iofs "io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

// CreateDir creates a directory.
func CreateDir(path string) error {
	return os.MkdirAll(path, 0755)
}

// RemoveDir removes a directory.
func RemoveDir(path string) error {
	return os.RemoveAll(path)
}

// RenameDir renames a directory.
func RenameDir(srcPath string, dstPath string) error {
	return os.Rename(srcPath, dstPath)
}

// MoveDir moves a directory.
func MoveDir(srcPath string, dstPath string) error {
	return os.Rename(srcPath, dstPath)
}

// ListDir lists the files in a directory.
func ListDir(path string) ([]iofs.FileInfo, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	return files, nil
}

// WalkDir walks the files in a directory.
func WalkDir(path string, fn iofs.WalkDirFunc) error {
	return filepath.WalkDir(path, fn)
}

// CopyDir copies a directory.
func CopyDir(srcPath string, dstPath string) error {
	if !IsExist(dstPath) {
		Mkdirp(dstPath)
	}

	return WalkDir(srcPath, func(path string, dir os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if dir.IsDir() {
			dstPath := strings.Replace(path, srcPath, dstPath, 1)
			if IsExist(dstPath) {
				return nil
			}

			return CreateDir(dstPath)
		}

		return CopyFile(path, strings.Replace(path, srcPath, dstPath, 1))
	})
}

// WatchDir watches the changes of files.
func WatchDir(ctx context.Context, paths []string, callback func(err error, event string, filepath string)) error {
	// @TODO cannot use go-zoox/debounce.New
	//	because go generic is bad
	debounce := func(fn func(err error, event string, filepath string), interval time.Duration) func(err error, event string, filepath string) error {
		var mu sync.Mutex
		var timer *time.Timer

		return func(err error, event string, filepath string) error {
			mu.Lock()
			defer mu.Unlock()

			if timer != nil {
				timer.Stop()
			}

			timer = time.AfterFunc(interval, func() {
				fn(err, event, filepath)
			})

			return nil
		}
	}

	debouncedCallback := debounce(callback, 300*time.Millisecond)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("failed to create watcher: %s", err)
	}

	for _, path := range paths {
		err := WalkDir(path, func(path string, d iofs.DirEntry, err error) error {
			if d.IsDir() {
				if err := watcher.Add(path); err != nil {
					return fmt.Errorf("failed to watch directory: %s (err: %s)", path, err)
				}
			}

			return nil
		})
		if err != nil {
			return fmt.Errorf("failed to walk dir: %v", err)
		}
	}

	for {
		select {
		case e := <-watcher.Events:
			if e.Op == fsnotify.Chmod {
				continue
			}

			op := ""
			switch e.Op {
			case fsnotify.Write:
				op = "WRITE"
			case fsnotify.Create:
				op = "CREATE"
			case fsnotify.Remove:
				op = "REMOVE"
			case fsnotify.Rename:
				op = "RENAME"
			default:
				op = fmt.Sprintf("UNKNOWN(%d)", e.Op)
			}

			debouncedCallback(nil, op, e.Name)
		case err := <-watcher.Errors:
			debouncedCallback(err, "ERROR", "")

		case <-ctx.Done():
			if err := watcher.Close(); err != nil {
				return fmt.Errorf("failed to close watcher: %s", err)
			}
			return nil
		}
	}
}
