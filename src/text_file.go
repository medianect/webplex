package main

import (
	"fmt"
	"os"
)

type TextFile struct {
	Path     string
	FileInfo os.FileInfo
	MimeType string
	Content  string
}

func loadTextFile(path string) (*TextFile, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	return &TextFile{
		Path: path,
		FileInfo: info,
		MimeType: "",
		Content:"",
	}, nil
	//dir := Directory{
	//	Path: path,
	//	Directories : make([]File, 0),
	//	Videos: make([]File, 0),
	//	Audios: make([]File, 0),
	//	Images: make([]File, 0),
	//	Texts: make([]File, 0),
	//	Unknowns: make([]File, 0),
	//	//Files: make([]File, 0),
	//}

	//for _, i := range(info) {
	//	if i.IsDir() {
	//		dir.Directories = append(dir.Directories, &DirectoryFile{FileInfo: i})
	//	} else {
	//		category, _ := GetFileType(filepath.Ext(i.Name()))
	//		switch category {
	//		case VIDEO:
	//			dir.Videos = append(dir.Videos, &VideoFile{FileInfo: i})
	//		case AUDIO:
	//			dir.Audios = append(dir.Audios, &AudioFile{FileInfo: i})
	//		case IMAGE:
	//			dir.Images = append(dir.Images, &ImageFile{FileInfo: i})
	//		case TEXT:
	//			dir.Texts = append(dir.Texts, &TextFile{FileInfo: i})
	//		default:
	//			dir.Unknowns = append(dir.Unknowns, &UnknownFile{FileInfo: i})
	//		}
	//	}
	//}
	//return &dir, nil
}

func (f *TextFile) Print() {
	fmt.Println("Textfile")
	//fmt.Println(f.Path + ":")
	//for _, dp := range(d.Directories) {
	//	fmt.Printf("d %s\n", dp.Name())
	//}
	//for _, f := range(d.Videos) {
	//	fmt.Printf("v %s\n", f.Name())
	//}
	//for _, f := range(d.Audios) {
	//	fmt.Printf("a %s\n", f.Name())
	//}
	//for _, f := range(d.Images) {
	//	fmt.Printf("i %s\n", f.Name())
	//}
	//for _, f := range(d.Texts) {
	//	fmt.Printf("t %s\n", f.Name())
	//}
	//for _, f := range(d.Unknowns) {
	//	fmt.Printf("u %s\n", f.Name())
	//}
}
