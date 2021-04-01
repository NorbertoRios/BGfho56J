package types

import (
	"os"
	"path/filepath"
)

//IFile file interface
type IFile interface {
	Path() string
}

//NewFileWithDir ...
func NewFileWithDir(_dir, _path string) *File {
	return &File{
		filePath: _path,
		dir:      _dir,
	}
}

//NewFile ...
func NewFile(_path string) *File {
	return NewFileWithDir(filepath.Dir(os.Args[0]), _path)
}

//File utils for files (configs)
type File struct {
	filePath string
	dir      string
}

//Path returns absolute file path
func (file File) Path() string {
	absPath, _ := filepath.Abs(file.dir + file.filePath)
	return absPath
}
