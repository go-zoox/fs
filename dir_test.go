package fs

import (
	iofs "io/fs"
	"testing"
)

func TestListDir(t *testing.T) {
	files, err := ListDir(".")
	if err != nil {
		t.Error(err)
	}

	if len(files) == 0 {
		t.Error("Expected at least one file")
	}

	for _, file := range files {
		t.Logf("[%s] %v - %d", file.Name(), file.IsDir(), file.Size())
	}
}

func TestWalkDir(t *testing.T) {
	err := WalkDir(".", func(path string, info iofs.DirEntry, err error) error {
		if err != nil {
			t.Error(err)
		}

		t.Logf("[%s] %v - %s", path, info.IsDir(), info.Type())
		return nil
	})

	if err != nil {
		t.Error(err)
	}
}
