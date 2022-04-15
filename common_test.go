package fs

import "testing"

func TestMkdir(t *testing.T) {
	path := "/tmp/test_mkdir"
	Remove(path)

	if IsExist(path) {
		t.Errorf("Expected path [%s] not exist", path)
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

func TestMkdirp(t *testing.T) {
	path := "/tmp/test_mkdirp"
	nestPath := path + "/nest1/nest2/nest3"
	Remove(path)

	if IsExist(path) {
		t.Errorf("Expected path [%s] not exist", path)
	}

	if IsExist(nestPath) {
		t.Errorf("Expected path [%s] not exist", nestPath)
	}

	if err := Mkdirp(path); err != nil {
		t.Error(err)
	}

	if !IsExist(path) {
		t.Errorf("Expected path [%s] exist", path)
	}

	if err := Mkdirp(nestPath); err != nil {
		t.Error(err)
	}

	if !IsExist(nestPath) {
		t.Errorf("Expected path [%s] exist", nestPath)
	}

	if err := Remove(path); err != nil {
		t.Error(err)
	}
}
