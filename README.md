# fs - A FileSystem Abstraction System for Go

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/fs)](https://pkg.go.dev/github.com/go-zoox/fs)
[![Build Status](https://github.com/go-zoox/fs/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/fs/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/fs)](https://goreportcard.com/report/github.com/go-zoox/fs)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/fs/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/fs?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/fs.svg)](https://github.com/go-zoox/fs/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/fs.svg?label=Release)](https://github.com/go-zoox/fs/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/fs
```

## Getting Started

```go
f := fs.OpenFile("/path/to/file")
```

## List of all available functions

```go
// go doc
func BaseName(path_ string) string
func Chmod(name string, mode os.FileMode) error
func Chown(name string, uid, gid int) error
func Copy(srcPath, dstPath string) error
func CopyDir(srcPath string, dstPath string) error
func CopyFile(srcPath string, dstPath string) error
func CreateDir(path string) error
func CreateFile(path string) error
func DirName(path_ string) string
func ExtName(path_ string) string
func IsAbsPath(path_ string) bool
func IsDir(path string) bool
func IsEmpty(path string) bool
func IsExist(path string) bool
func IsFile(path string) bool
func IsLink(path string) bool
func JoinPath(paths ...string) string
func ListDir(path string) ([]iofs.FileInfo, error)
func Mkdir(path string) error
func Mkdirp(path string) error
func Move(srcPath, dstPath string) error
func MoveDir(srcPath string, dstPath string) error
func MoveFile(srcPath, dstPath string) error
func Open(path string) (*os.File, error)
func OpenFile(path string) (*os.File, error)
func ReadFile(srcPath string) ([]byte, error)
func ReadFileAsString(srcPath string) (string, error)
func ReadFileByLine(srcPath string) ([]string, error)
func Remove(path string) error
func RemoveDir(path string) error
func RemoveFile(path string) error
func Rename(srcPath, dstPath string) error
func RenameDir(srcPath string, dstPath string) error
func RenameFile(srcPath, dstPath string) error
func Stat(path string) (os.FileInfo, error)
func TmpDirPath() string
func TmpFilePath() string
func Walk(path string, fn iofs.WalkDirFunc) error
func WalkDir(path string, fn iofs.WalkDirFunc) error
func WriteFile(path string, data []byte) error
```

## License
GoZoox is released under the [MIT License](./LICENSE).