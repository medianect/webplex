package main

import (
	"os"
	"time"
)

type VideoFile struct {
	FileInfo os.FileInfo
}

func (f *VideoFile) Name() string {
	return f.FileInfo.Name()
}

func (f *VideoFile) Size() int64 {
	return f.FileInfo.Size()
}

func (f *VideoFile) Mode() os.FileMode {
	return f.FileInfo.Mode()
}

func (f *VideoFile) ModTime() time.Time {
	return f.FileInfo.ModTime()
}
