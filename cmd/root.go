package cmd

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/fsnotify/fsnotify"
	"github.com/notbloom/ficus/config"
	"github.com/notbloom/ficus/npm"
	"github.com/notbloom/ficus/server"
	"github.com/spf13/cobra"
	"os"
)

var VerboseDebug bool

func init() {
	rootCmd.Flags().BoolVarP(&VerboseDebug, "verbose", "v", false, "verbose output")
}

var rootCmd = &cobra.Command{
	Use:   "front-dev",
	Short: "Front-Dev is server static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		//fmt.Println("ejecutar cosa", args)

		// Read config file or get default
		cfg := config.GetConfig()
		log.Info(cfg.ViewsRoute)

		log.Info("Server Started at http://127.0.0.1:8000")
		log.Info("Serving /layout/:id for /layout/:id.html")
		log.Info("Serving /pages/:id for /pages/:id.html")
		log.Info("File watching")

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
						if event.Name[len(event.Name)-1:] == "~" {
							log.Debug("Ignoring temp~ file ")
						} else {
							log.Print("modified file:", event.Name)
							log.Info("Building Tailwind...")
							npm.RunTailwind("", "")
							server.SendRefresh()
							//b.Reload()
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
		err = watcher.Add("./views/pages/")
		if err != nil {
			log.Fatal(err)
		}
		log.Info("Watching: ./views/pages/")

		server.StartServer()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
