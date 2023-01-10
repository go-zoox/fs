package fs

import "testing"

func TestIsExist(t *testing.T) {
	path := "/tmp/test_is_exist"
	Remove(path)

	if IsExist(path) {
		t.Errorf("Expected path [%s] not exist", path)
	}

	if err := WriteFile(path, []byte("test")); err != nil {
		t.Error(err)
	}

	if !IsExist(path) {
		t.Errorf("Expected path [%s] exist", path)
	}

	if err := Remove(path); err != nil {
		t.Errorf("Expected path [%s] exist", path)
	}

	if err := Mkdir(path); err != nil {
		t.Error(err)
	}

	if !IsExist(path) {
		t.Errorf("Expected path [%s] exist", path)
	}

	if err := Remove(path); err != nil {
		t.Error(err)
	}
}

func TestIsFile(t *testing.T) {
	if IsFile("/tmp") {
		t.Error("/tmp is not a file")
	}

	path := "/tmp/test_is_file"
	if err := WriteFile(path, []byte("test")); err != nil {
		t.Error(err)
	}

	if !IsFile(path) {
		t.Errorf("Expected path [%s] is a file", path)
	}

	if err := Remove(path); err != nil {
		t.Error(err)
	}
}

func TestIsDir(t *testing.T) {
	if !IsDir("/tmp") {
		t.Error("/tmp is not a file")
	}

	path := "/tmp/test_is_file"
	if err := WriteFile(path, []byte("test")); err != nil {
		t.Error(err)
	}

	if IsDir(path) {
		t.Errorf("Expected path [%s] is a dir", path)
	}

	if err := Remove(path); err != nil {
		t.Error(err)
	}
}

func TestIsSymbolicLink(t *testing.T) {
	if !IsSymbolicLink("/tmp") {
		t.Error("/tmp is not a symbolic link")
	}

	testpath := "/tmp/TestIsSymbolicLink"
	if err := Mkdirp(testpath); err != nil {
		t.Error(err)
	}

	originFilePath := testpath + "/origin.txt"
	symbolicLinkFilePath := testpath + "/symbolic_link.txt"
	if err := WriteFile(originFilePath, []byte("test")); err != nil {
		t.Error(err)
	}

	if !IsExist(symbolicLinkFilePath) {
		if err := CreateSymbolicLink(originFilePath, symbolicLinkFilePath); err != nil {
			t.Error(err)
		}
	}

	if !IsSymbolicLink(symbolicLinkFilePath) {
		t.Errorf("Expected path [%s] is a symbolic link", symbolicLinkFilePath)
	}

	if err := Remove(testpath); err != nil {
		t.Error(err)
	}
}

func TestIsEmpty(t *testing.T) {
	path := "/tmp/test_is_empty"
	if err := Mkdirp(path); err != nil {
		t.Error(err)
	}

	if !IsEmpty(path) {
		t.Errorf("Expected path [%s] is empty", path)
	}

	if err := WriteFile(path+"/test.txt", []byte("test")); err != nil {
		t.Error(err)
	}

	if IsEmpty(path) {
		t.Errorf("Expected path [%s] is not empty", path)
	}

	if err := Remove(path); err != nil {
		t.Error(err)
	}
}
