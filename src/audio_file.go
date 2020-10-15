package main

import (
	"os"
	"time"
)

type AudioFile struct {
	FileInfo os.FileInfo
}

func (f *AudioFile) Name() string {
	return f.FileInfo.Name()
}

func (f *AudioFile) Size() int64 {
	return f.FileInfo.Size()
}

func (f *AudioFile) Mode() os.FileMode {
	return f.FileInfo.Mode()
}

func (f *AudioFile) ModTime() time.Time {
	return f.FileInfo.ModTime()
}
