package server

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Link struct {
	Url   string
	Title string
}

func HomeList(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	directory := "./views/pages/" // The current directory

	files, err := os.ReadDir(directory) //read the files from the directory
	if err != nil {
		fmt.Println("error reading directory:", err) //print error if directory is not read properly
		return
	}

	var htmlFiles []Link //declare a slice to store the HTML files

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".html") {
			fileName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			htmlFiles = append(htmlFiles, Link{Url: "/pages/" + fileName + "/", Title: file.Name()}) //append the HTML file names to the slice
		}
	}

	//	fmt.Println(htmlFiles)

	path := "./views/index.html"

	if _, err := os.Stat(path); err != nil {
		fmt.Fprintf(w, "template page not found: %v\n", path)
		return
	}

	tmpl := template.Must(template.ParseFiles(path))
	err = tmpl.ExecuteTemplate(w, "index.html", htmlFiles)

	if err != nil {
		fmt.Fprintf(w, "exec template error path: %v\n", path)
		fmt.Fprintf(w, "error: %v\n", err)
	}
}
