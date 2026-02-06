package server

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gorilla/mux"
	"github.com/notbloom/ficus/config"
)

var funcMap = template.FuncMap{
	"N": N,
	"M": M,
}

// Package-level config reference
var cfg *config.Config

// Package-level broker reference for direct reload calls
var broker *Broker

// Package-level server reference for shutdown
var srv *http.Server

func StartServer(c *config.Config) {
	cfg = c

	r := mux.NewRouter()
	broker = NewBrokerServer()

	r.HandleFunc("/reload/messages", broker.BroadcastMessage).Methods("POST")
	r.HandleFunc("/reload/stream", broker.Stream).Methods("GET")

	r.HandleFunc("/pages/{page}/", PagesHandler)
	r.HandleFunc("/", HomeList)

	fs := http.FileServer(http.Dir(cfg.AssetsPath()))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	srv = &http.Server{
		Handler:      r,
		Addr:         cfg.ServerAddr(),
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server error", "error", err)
	}
}

// Shutdown gracefully shuts down the server
func Shutdown(ctx context.Context) error {
	if srv != nil {
		return srv.Shutdown(ctx)
	}
	return nil
}

// GetBroker returns the broker for direct reload calls
func GetBroker() *Broker {
	return broker
}

func N(start, end int) (stream chan int) {
	stream = make(chan int)
	go func() {
		for i := start; i <= end; i++ {
			stream <- i
		}
		close(stream)
	}()
	return
}

func M(start, end int) (stream chan int) {
	stream = make(chan int)
	go func() {
		for i := start; i >= end; i-- {
			stream <- i
		}
		close(stream)
	}()
	return
}

func GetPageData(name string) any {
	path := filepath.Join(cfg.PagesPath(), name+".json")

	if _, err := os.Stat(path); err != nil {
		log.Debug("page data not found", "path", path)
		return nil
	}
	input, err := os.ReadFile(path)
	if err != nil {
		log.Debug("error reading page data", "error", err)
		return nil
	}
	m := map[string]interface{}{}
	err = json.Unmarshal(input, &m)
	if err != nil {
		log.Debug("error parsing page data", "error", err)
		return nil
	}
	log.Debug("Loaded json data", "data", m)
	return m
}

func PagesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	path := filepath.Join(cfg.PagesPath(), vars["page"]+cfg.Folders.ViewsExtension)

	if _, err := os.Stat(path); err != nil {
		fmt.Fprintf(w, "template page not found: %v\n", path)
		return
	}

	tmpl := template.Must(template.New(vars["page"]).Funcs(funcMap).ParseFiles(path))
	err := tmpl.ExecuteTemplate(w, vars["page"]+cfg.Folders.ViewsExtension, GetPageData(vars["page"]))

	if err != nil {
		fmt.Fprintf(w, "exec template error path: %v\n", path)
		fmt.Fprintf(w, "error: %v\n", err)
	}
}
