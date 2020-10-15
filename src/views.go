package main

import (
	"fmt"
	"log"
	"html/template"
	"github.com/medianect/octicon"
)

var viewsFuncs template.FuncMap

func LoadView(name string, content []string) {
	tmpl := template.New(name).Funcs(viewsFuncs)
	for i, c := range(content) {
		var err error
		tmpl, err = tmpl.Parse(c)
		if err != nil {
			log.Fatalf("Failed to load template source `%s' (%d): %s", name, i, err.Error())
		}
	}
	views[name] = tmpl
}

var views map[string]*template.Template = make(map[string]*template.Template, 0)

const Kilo int64 = 1000
const Mega int64 = 1000 * Kilo
const Giga int64 = 1000 * Mega
const Tera int64 = 1000 * Giga
const Peta int64 = 1000 * Tera

func init() {
	viewsFuncs = template.FuncMap{
		"DashIcon": func() template.HTML {return template.HTML(octicon.Icon("dash", 24, 24))},
		"DeviceCameraVideoIcon": func() template.HTML {return template.HTML(octicon.Icon("device-camera-video", 24, 24))},
		"FileBinaryIcon": func() template.HTML {return template.HTML(octicon.Icon("file-binary", 24, 24))},
		"FileCodeIcon": func() template.HTML {return template.HTML(octicon.Icon("file-code", 24, 24))},
		"FileDirectoryIcon": func() template.HTML {return template.HTML(octicon.Icon("file-directory", 24, 24))},
		"FileMediaIcon": func() template.HTML {return template.HTML(octicon.Icon("file-media", 24, 24))},
		"ImageIcon": func() template.HTML {return template.HTML(octicon.Icon("image", 24, 24))},

		"HumanReadableSize": func(size int64) string {
			if size >= Peta {
				return fmt.Sprintf("%.2fP", float64(size) / float64(Peta))
			} else if size >= Tera {
				return fmt.Sprintf("%.2fT", float64(size) / float64(Tera))
			} else if size >= Giga {
				return fmt.Sprintf("%.2fG", float64(size) / float64(Giga))
			} else if size >= Mega {
				return fmt.Sprintf("%.2fM", float64(size) / float64(Mega))
			} else if size >= Kilo {
				return fmt.Sprintf("%.2fK", float64(size) / float64(Kilo))
			} else {
				return fmt.Sprintf("%d", size)
			}
		},
	}
	LoadView("directory", []string{page_html, directory_html})
	LoadView("text_file", []string{page_html, text_file_html})
}
