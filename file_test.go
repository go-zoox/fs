package fs

import "testing"

func TestCreateFile(t *testing.T) {
	path := "/tmp/test_create_file"
	Remove(path)

	if IsExist(path) {
		t.Errorf("Expected path [%s] not exist", path)
	}

	if err := CreateFile(path); err != nil {
		t.Error(err)
	}

	if !IsExist(path) {
		t.Errorf("Expected path [%s] exist", path)
	}

	if !IsFile(path) {
		t.Errorf("Expected path [%s] is a file", path)
	}

	if err := Remove(path); err != nil {
		t.Error(err)
	}
}

func TestRemoveFile(t *testing.T) {
	path := "/tmp/test_remove_file"
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
		t.Error(err)
	}

	if IsExist(path) {
		t.Errorf("Expected path [%s] not exist", path)
	}
}

func TestCop(t *testing.T) {
	src := "/tmp/test_copy_file"
	dst := "/tmp/test_copy_file_dst"
	Remove(src)
	Remove(dst)

	if IsExist(src) {
		t.Errorf("Expected path [%s] not exist", src)
	}

	if IsExist(dst) {
		t.Errorf("Expected path [%s] not exist", dst)
	}

	if err := WriteFile(src, []byte("test")); err != nil {
		t.Error(err)
	}

	if !IsExist(src) {
		t.Errorf("Expected path [%s] exist", src)
	}

	if err := Copy(src, dst); err != nil {
		t.Error(err)
	}

	if !IsExist(dst) {
		t.Errorf("Expected path [%s] exist", dst)
	}

	if err := Remove(src); err != nil {
		t.Error(err)
	}

	if err := Remove(dst); err != nil {
		t.Error(err)
	}
}

func TestRenameFile(t *testing.T) {
	src := "/tmp/test_rename_file"
	dst := "/tmp/test_rename_file_dst"
	Remove(src)
	Remove(dst)

	if IsExist(src) {
		t.Errorf("Expected path [%s] not exist", src)
	}

	if IsExist(dst) {
		t.Errorf("Expected path [%s] not exist", dst)
	}

	if err := WriteFile(src, []byte("test")); err != nil {
		t.Error(err)
	}

	if !IsExist(src) {
		t.Errorf("Expected path [%s] exist", src)
	}

	if err := Rename(src, dst); err != nil {
		t.Error(err)
	}

	if IsExist(src) {
		t.Errorf("Expected path [%s] not exist", src)
	}

	if !IsExist(dst) {
		t.Errorf("Expected path [%s] exist", dst)
	}

	Remove(src)
	Remove(dst)
}

func TestMoveFi(t *testing.T) {
	src := "/tmp/test_move_file"
	dst := "/tmp/test_move_file_dst"
	Remove(src)
	Remove(dst)

	if IsExist(src) {
		t.Errorf("Expected path [%s] not exist", src)
	}

	if IsExist(dst) {
		t.Errorf("Expected path [%s] not exist", dst)
	}

	if err := WriteFile(src, []byte("test")); err != nil {
		t.Error(err)
	}

	if !IsExist(src) {
		t.Errorf("Expected path [%s] exist", src)
	}

	if err := Move(src, dst); err != nil {
		t.Error(err)
	}

	if IsExist(src) {
		t.Errorf("Expected path [%s] not exist", src)
	}

	if !IsExist(dst) {
		t.Errorf("Expected path [%s] exist", dst)
	}

	Remove(src)
	Remove(dst)
}

func TestOpenFile(t *testing.T) {
	path := "/tmp/test_open_file"
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

	if f, err := OpenFile(path); err != nil {
		t.Error(err)
	} else {
		defer f.Close()
	}

	if err := Remove(path); err != nil {
		t.Error(err)
	}
}

func TestWriteFile(t *testing.T) {
	path := "/tmp/test_write_file"
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
		t.Error(err)
	}
}

func TestReadFile(t *testing.T) {
	path := "/tmp/test_read_file"
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

	if b, err := ReadFile(path); err != nil {
		t.Error(err)
	} else {
		if string(b) != "test" {
			t.Errorf("Expected content [%s] not equal [%s]", "test", string(b))
		}
	}

	if err := Remove(path); err != nil {
		t.Error(err)
	}
}

func TestReadFileAsString(t *testing.T) {
	path := "/tmp/test_read_file_as_string"
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

	if s, err := ReadFileAsString(path); err != nil {
		t.Error(err)
	} else {
		if s != "test" {
			t.Errorf("Expected content [%s] not equal [%s]", "test", s)
		}
	}

	if err := Remove(path); err != nil {
		t.Error(err)
	}
}

func TestReadFileLines(t *testing.T) {
	path := "/tmp/test_read_file_by_line"
	Remove(path)

	if IsExist(path) {
		t.Errorf("Expected path [%s] not exist", path)
	}

	if err := WriteFile(path, []byte("test1\ntest2\ntest3")); err != nil {
		t.Error(err)
	}

	if !IsExist(path) {
		t.Errorf("Expected path [%s] exist", path)
	}

	if lines, err := ReadFileLines(path); err != nil {
		t.Error(err)
	} else {
		if len(lines) != 3 {
			t.Errorf("Expected lines [%d] not equal [%d]", 3, len(lines))
		}
	}

	if err := Remove(path); err != nil {
		t.Error(err)
	}
}
