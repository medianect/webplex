package main

import (
	"mime"
	"strings"
)

type FileType string

const (
	VIDEO = "video"
	AUDIO = "audio"
	IMAGE = "image"
	TEXT = "text"
	UNKNOWN = "unknown"
)

func init() {
	_ = mime.AddExtensionType(".go", "text/x-golang")
}

var textFilesTypes map[string]bool = map[string]bool{
	"application/json": true,
	"application/x-bsh": true,
	"application/x-httpd-php": true,
	"application/x-sh": true,
	"application/x-shar": true,
}

func GetFileType(ext string) (FileType, string) {
	ext = strings.ToLower(ext)
	typename := mime.TypeByExtension(ext)
	var category FileType

	if strings.HasPrefix(typename, "video/") {
		category = VIDEO
	} else if strings.HasPrefix(typename, "audio/") {
		category = AUDIO
	} else if strings.HasPrefix(typename, "image/") {
		category = IMAGE
	} else if strings.HasPrefix(typename, "text/") {
		category = TEXT
	} else if strings.HasSuffix(typename, "+json") {
		category = TEXT
	} else if strings.HasSuffix(typename, "+xml") {
		category = TEXT
	} else if textFilesTypes[typename] {
		category = TEXT
	} else {
		category = UNKNOWN
	}
	return category, typename
}
