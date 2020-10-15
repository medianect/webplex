package main

import (
	"os"
	"time"
)

type ImageFile struct {
	FileInfo os.FileInfo
}

func (f *ImageFile) Name() string {
	return f.FileInfo.Name()
}

func (f *ImageFile) Size() int64 {
	return f.FileInfo.Size()
}

func (f *ImageFile) Mode() os.FileMode {
	return f.FileInfo.Mode()
}

func (f *ImageFile) ModTime() time.Time {
	return f.FileInfo.ModTime()
}
