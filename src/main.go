/* Copyright © 2018 Xvdfe <xvdfe@medianect.com>
   SPDX-License-Identifier: BSD-2-Clause */
package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	//"strings"
	"github.com/pborman/getopt/v2"
)

func editHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/view/"):]
	// TODO: If security mode is "clientGiven" then use filepath.Clean() instead
	path = filepath.Clean(path)
	//if strings.HasPrefix("../") {
	//	// TODO: Throw error
	//}
	// TODO: If security mode is not "withHiddenFiles"
	//if strings.HasPrefix(".") {
	//	// TODO: Throw error
	//}
	w.Header().Set("Content-Type", "text/html")

	// There's a race condition here on name changing type between
	// calls. Code should fail gracefully on these as we're web
	// browsing so people stay in front of a page rendered minutes
	// ago anyway.
	fi, err := os.Stat(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if fi.IsDir() {
		d, err := loadDirectory(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		views["directory"].Execute(w, d)
	} else {
		// TODO: Handle devices, fifos...
		category, _ := GetFileType(filepath.Ext(fi.Name()))
		switch category {
		case VIDEO, AUDIO, IMAGE:
			http.Error(w, "Not implemented yet.", http.StatusInternalServerError)
		case TEXT:
			t, err := loadTextFile(path)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tpl, found := views["text_file"]
			if !found {
				// TODO: Log more info
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			tpl.Execute(w, t)
		default:
			http.Error(w, "Not implemented yet.", http.StatusInternalServerError)
		}
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	views["login"].Execute(w, nil)
}

var titleFlag = getopt.StringLong("title", 't', "Documents", "Root directory title")

func main() {
	getopt.Parse()
	args := getopt.Args()
	if len(args) == 0 {
		args = append(args, "listen")
	}
	switch args[0] {
	case "listen":
		http.HandleFunc("/edit/", editHandler)
		http.HandleFunc("/view/", viewHandler)
		http.HandleFunc("/", homeHandler)
		fmt.Printf("Server will listen to \":8080\"\n")
		log.Fatal(http.ListenAndServe(":8080", nil))
	case "ls":
		if len(args) == 1 {
			args = append(args, ".")
		}
		for _, path := range(args[1:]) {
			d, e := loadDirectory(path)
			if e != nil {
				log.Fatal("Failed to load directory", e)
			}
			d.Print()
		}
	default:
		log.Fatal("Unknown command")
	}
}
