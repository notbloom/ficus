package watcher

import (
	"os"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/fsnotify/fsnotify"
	"github.com/notbloom/ficus/config"
)

type Watcher struct {
	fsWatcher *fsnotify.Watcher
	cfg       *config.Config
	OnChange  func(path string)
	done      chan struct{}
}

func New(cfg *config.Config) (*Watcher, error) {
	fsWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	return &Watcher{
		fsWatcher: fsWatcher,
		cfg:       cfg,
		done:      make(chan struct{}),
	}, nil
}

func (w *Watcher) Start() error {
	// Add watched paths
	paths := []string{w.cfg.PagesPath()}

	// Also watch components if directory exists
	if _, err := os.Stat(w.cfg.ComponentsPath()); err == nil {
		paths = append(paths, w.cfg.ComponentsPath())
	}

	for _, path := range paths {
		if err := w.fsWatcher.Add(path); err != nil {
			log.Warn("Could not watch path", "path", path, "error", err)
			continue
		}
		log.Info("Watching", "path", path)
	}

	go w.listen()
	return nil
}

func (w *Watcher) Stop() {
	close(w.done)
	w.fsWatcher.Close()
}

func (w *Watcher) listen() {
	for {
		select {
		case <-w.done:
			return
		case event, ok := <-w.fsWatcher.Events:
			if !ok {
				return
			}
			if event.Has(fsnotify.Write) {
				if w.shouldIgnore(event.Name) {
					log.Debug("Ignoring file", "path", event.Name)
					continue
				}
				log.Info("File modified", "path", event.Name)
				if w.OnChange != nil {
					w.OnChange(event.Name)
				}
			}
		case err, ok := <-w.fsWatcher.Errors:
			if !ok {
				return
			}
			log.Error("Watch error", "error", err)
		}
	}
}

func (w *Watcher) shouldIgnore(path string) bool {
	for _, suffix := range w.cfg.Watch.ExcludeSuffix {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}
	return false
}
