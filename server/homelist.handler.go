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
	w.WriteHeader(http.StatusOK)

	directory := cfg.PagesPath()

	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Fprintf(w, "error reading directory: %v\n", err)
		return
	}

	var htmlFiles []Link

	for _, file := range files {
		if strings.HasSuffix(file.Name(), cfg.Folders.ViewsExtension) {
			fileName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			htmlFiles = append(htmlFiles, Link{
				Url:   "/pages/" + fileName + "/",
				Title: file.Name(),
			})
		}
	}

	path := filepath.Join(cfg.Folders.ViewsFolder, "index.html")

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
