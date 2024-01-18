package watcher

import (
	"github.com/charmbracelet/log"
	"github.com/fsnotify/fsnotify"
	"github.com/notbloom/ficus/npm"
	"github.com/notbloom/ficus/server"
)

type Config struct {
	IncludeSuffix     []string
	ExcludeSuffix     []string
	Folders           []string
	IncludeSubFolders bool
}

func Start(config Config, OnChange func(event fsnotify.Event)) {

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}
	defer func(watcher *fsnotify.Watcher) {
		err := watcher.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(watcher)

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
					if event.Name[len(event.Name)-1:] == "~" {
						log.Debug("Ignoring temp~ file ")
					} else {
						log.Print("modified file:", event.Name)
						log.Info("Building Tailwind...")
						npm.RunTailwind("", "")
						server.SendRefresh()
						//b.Reload()
						OnChange(event)
					}
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
	for _, folder := range config.Folders {
		err = watcher.Add(folder)
		if err != nil {
			log.Fatal(err)
		}
		log.Info("Watching: " + folder)
	}
	/*err = watcher.Add("./views/pages/")
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Watching: ./views/pages/")*/
}
