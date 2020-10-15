package main

import (
	"os"
	"time"
)

type DirectoryFile struct {
	FileInfo os.FileInfo
}

func (f *DirectoryFile) Name() string {
	return f.FileInfo.Name()
}

func (f *DirectoryFile) Size() int64 {
	return f.FileInfo.Size()
}

func (f *DirectoryFile) Mode() os.FileMode {
	return f.FileInfo.Mode()
}

func (f *DirectoryFile) ModTime() time.Time {
	return f.FileInfo.ModTime()
}
