package main

import (
	"fmt"
	"io/ioutil"
	"os"
	_ "path"
	"path/filepath"
	_ "sort"
)

type Directory struct {
	Path        string
	Directories []os.FileInfo
	Videos      []os.FileInfo
	Audios      []os.FileInfo
	Images      []os.FileInfo
	Files       []os.FileInfo
	Texts       []os.FileInfo
	Unknowns    []os.FileInfo
}

// TODO: Use a multisort:
// - by type (default) as below
// - by size, keep separating directories from files

func loadDirectory(path string) (*Directory, error) {
	info, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	dir := Directory{
		Path: path,
		Directories : make([]os.FileInfo, 0),
		Videos: make([]os.FileInfo, 0),
		Audios: make([]os.FileInfo, 0),
		Images: make([]os.FileInfo, 0),
		Texts: make([]os.FileInfo, 0),
		Unknowns: make([]os.FileInfo, 0),
		//Files: make([]os.FileInfo, 0),
	}

	for _, i := range(info) {
		if i.IsDir() {
			dir.Directories = append(dir.Directories, i)
		} else {
			category, _ := GetFileType(filepath.Ext(i.Name()))
			switch category {
			case VIDEO:
				dir.Videos = append(dir.Videos, i)
			case AUDIO:
				dir.Audios = append(dir.Audios, i)
			case IMAGE:
				dir.Images = append(dir.Images, i)
			case TEXT:
				dir.Texts = append(dir.Texts, i)
			default:
				dir.Unknowns = append(dir.Unknowns, i)
			}
		}
	}
	return &dir, nil
}

func (d *Directory) Print() {
	fmt.Println(d.Path + ":")
	for _, i := range(d.Directories) {
		fmt.Printf("d %s\n", i.Name())
	}
	for _, i := range(d.Videos) {
		fmt.Printf("v %s\n", i.Name())
	}
	for _, i := range(d.Audios) {
		fmt.Printf("a %s\n", i.Name())
	}
	for _, i := range(d.Images) {
		fmt.Printf("i %s\n", i.Name())
	}
	for _, i := range(d.Texts) {
		fmt.Printf("t %s\n", i.Name())
	}
	for _, i := range(d.Unknowns) {
		fmt.Printf("u %s\n", i.Name())
	}
}
