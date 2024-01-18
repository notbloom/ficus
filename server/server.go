package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"
)

var funcMap = template.FuncMap{
	"N": N,
	"M": M,
}

func StartServer() {
	r := mux.NewRouter()
	b := NewBrokerServer()

	//b.Reload()

	r.HandleFunc("/reload/messages", b.BroadcastMessage).Methods("POST")
	r.HandleFunc("/reload/stream", b.Stream).Methods("GET")

	r.HandleFunc("/pages/{page}/", PagesHandler)
	r.HandleFunc("/", HomeList)
	//r.HandleFunc("/articles", ArticlesHandler)
	fs := http.FileServer(http.Dir("./views/assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info("Server Started at http://127.0.0.1:8000")
	log.Fatal(srv.ListenAndServe())
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
	path := "./views/pages/" + name + ".json"

	if _, err := os.Stat(path); err != nil {
		log.Debug("page data not found: %v\n", path)
		log.Debug(err)
		return nil
	}
	input, err := os.ReadFile(path)
	if err != nil {
		log.Debug(err)
		return nil
	}
	m := map[string]interface{}{}
	err = json.Unmarshal(input, &m)
	if err != nil {
		log.Debug(err)
		return nil
	}
	log.Info("Loaded json data")
	log.Info(m)
	return m
}

func PagesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	path := "./views/pages/" + vars["page"] + ".html"

	if _, err := os.Stat(path); err != nil {
		fmt.Fprintf(w, "template page not found: %v\n", path)
		return
	}

	tmpl := template.Must(template.New(vars["page"]).Funcs(funcMap).ParseFiles(path))
	err := tmpl.ExecuteTemplate(w, vars["page"]+".html", GetPageData(vars["page"]))

	if err != nil {
		fmt.Fprintf(w, "exec template error path: %v\n", path)
		fmt.Fprintf(w, "error: %v\n", err)
	}
}

func SendRefresh() {
	url := "http://127.0.0.1:8000/reload/messages"
	log.Info("Refresh page")

	var jsonStr = []byte(`{"name":"server", "msg":"refresh"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.ReadAll(resp.Body)

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := io.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
}
