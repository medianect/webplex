package main

import (
	"os"
	"time"
)

type UnknownFile struct {
	FileInfo os.FileInfo
}

func (f *UnknownFile) Name() string {
	return f.FileInfo.Name()
}

func (f *UnknownFile) Size() int64 {
	return f.FileInfo.Size()
}

func (f *UnknownFile) Mode() os.FileMode {
	return f.FileInfo.Mode()
}

func (f *UnknownFile) ModTime() time.Time {
	return f.FileInfo.ModTime()
}
