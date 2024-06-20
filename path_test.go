package fs

import (
	"fmt"
	"testing"
)

func TestTmpFilePath(t *testing.T) {
	path := TmpFilePath()
	if path == "" {
		t.Fatal("Expected path not empty")
	}

	if IsExist(path) {
		t.Fatalf("Expected path [%s] exist", path)
	}

	fmt.Println("TmpFilePath:", path)
}
