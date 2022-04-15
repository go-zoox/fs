package fs

import (
	iofs "io/fs"
	"testing"
)

func TestCreateDir(t *testing.T) {
	path := "/tmp/test_create_dir"
	nestPath := path + "/nest1/nest2/nest3"
	Remove(path)

	if IsExist(path) {
		t.Errorf("Expected path [%s] not exist", path)
	}

	if IsExist(nestPath) {
		t.Errorf("Expected path [%s] not exist", nestPath)
	}

	if err := CreateDir(path); err != nil {
		t.Error(err)
	}

	if !IsExist(path) {
		t.Errorf("Expected path [%s] exist", path)
	}

	if err := CreateDir(nestPath); err != nil {
		t.Error(err)
	}

	if !IsExist(nestPath) {
		t.Errorf("Expected path [%s] exist", nestPath)
	}

	if err := Remove(path); err != nil {
		t.Error(err)
	}
}

func TestRemoveDir(t *testing.T) {
	path := "/tmp/test_remove_dir"
	nestPath := path + "/nest1/nest2/nest3"
	Remove(path)

	if IsExist(path) {
		t.Errorf("Expected path [%s] not exist", path)
	}

	if IsExist(nestPath) {
		t.Errorf("Expected path [%s] not exist", nestPath)
	}

	if err := CreateDir(path); err != nil {
		t.Error(err)
	}

	if !IsExist(path) {
		t.Errorf("Expected path [%s] exist", path)
	}

	if err := CreateDir(nestPath); err != nil {
		t.Error(err)
	}

	if !IsExist(nestPath) {
		t.Errorf("Expected path [%s] exist", nestPath)
	}

	if err := RemoveDir(path); err != nil {
		t.Error(err)
	}

	if IsExist(path) {
		t.Errorf("Expected path [%s] not exist", path)
	}

	if IsExist(nestPath) {
		t.Errorf("Expected path [%s] not exist", nestPath)
	}

	if err := Remove(path); err != nil {
		t.Error(err)
	}
}

func TestRenameDir(t *testing.T) {
	path := "/tmp/test_rename_dir"
	beforePath := path + "/before"
	afterPath := path + "/after"
	Remove(path)

	Mkdirp(beforePath)
	if !IsExist(beforePath) {
		t.Errorf("Expected path [%s] exist", beforePath)
	}

	if IsExist(afterPath) {
		t.Errorf("Expected path [%s] not exist", afterPath)
	}

	if err := RenameDir(beforePath, afterPath); err != nil {
		t.Error(err)
	}

	if IsExist(beforePath) {
		t.Errorf("Expected path [%s] not exist", beforePath)
	}

	if !IsExist(afterPath) {
		t.Errorf("Expected path [%s] exist", afterPath)
	}

	if err := Remove(afterPath); err != nil {
		t.Error(err)
	}
}

func TestMoveDir(t *testing.T) {
	path := "/tmp/test_move_dir"
	beforePath := path + "/before"
	beforeNestPath := beforePath + "/nest1/nest2/nest3"
	afterPath := path + "/after"
	afterNestPath := afterPath + "/nest1/nest2/nest3"
	Remove(path)

	Mkdirp(beforeNestPath)
	if !IsExist(beforeNestPath) {
		t.Errorf("Expected path [%s] exist", beforeNestPath)
	}

	if IsExist(afterNestPath) {
		t.Errorf("Expected path [%s] not exist", afterNestPath)
	}

	if err := MoveDir(beforePath, afterPath); err != nil {
		t.Error(err)
	}

	if IsExist(beforePath) {
		t.Errorf("Expected path [%s] not exist", beforePath)
	}

	if !IsExist(afterPath) {
		t.Errorf("Expected path [%s] exist", afterPath)
	}

	if !IsExist(afterNestPath) {
		t.Errorf("Expected path [%s] exist", afterNestPath)
	}

	if err := Remove(afterPath); err != nil {
		t.Error(err)
	}
}

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

func TestCopyDir(t *testing.T) {
	path := "/tmp/test_copy_dir"

	originPath := path + "/nest1"
	originNestPath := originPath + "/nest2/nest3"
	originFilePath := originPath + "/test.txt"

	copiedPath := path + "/nest1_copy"
	copiedNestPath := copiedPath + "/nest2/nest3"
	copiedFilePath := copiedPath + "/test.txt"
	Remove(path)
	Mkdirp(originNestPath)
	WriteFile(originFilePath, []byte("test"))

	if IsExist(copiedPath) {
		t.Errorf("Expected path [%s] not exist", copiedPath)
	}

	if IsExist(copiedNestPath) {
		t.Errorf("Expected path [%s] not exist", copiedNestPath)
	}

	if IsExist(copiedFilePath) {
		t.Errorf("Expected path [%s] not exist", copiedFilePath)
	}

	if err := Copy(originPath, copiedPath); err != nil {
		t.Error(err)
	}

	if !IsExist(copiedPath) {
		t.Errorf("Expected path [%s] exist", copiedPath)
	}

	if !IsExist(copiedNestPath) {
		t.Errorf("Expected path [%s] exist", copiedNestPath)
	}

	if !IsExist(copiedFilePath) {
		t.Errorf("Expected path [%s] exist", copiedFilePath)
	}

	if err := Remove(path); err != nil {
		t.Error(err)
	}
}
