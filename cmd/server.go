package cmd

/*
func main() {
	log.SetLevel(log.DebugLevel)
	config.GetConfig()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				//log.Println("event:", event)
				if event.Has(fsnotify.Write) {
					log.Print("modified file:", event.Name)
					server.SendRefresh()
					//b.Reload()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Print("error:", err)
			}
		}
	}()

	// Add a path.
	err = watcher.Add("./frontend/views/pages/")
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Watching: ./frontend/views/pages/")

	r := mux.NewRouter()
	b := server.NewBrokerServer()

	//b.Reload()

	r.HandleFunc("/reload/messages", b.BroadcastMessage).Methods("POST")
	r.HandleFunc("/reload/stream", b.Stream).Methods("GET")

	//r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/dev/pages/{page}", server.PagesHandler)
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
*/
