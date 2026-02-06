package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
	"github.com/notbloom/ficus/config"
	"github.com/notbloom/ficus/npm"
	"github.com/notbloom/ficus/server"
	"github.com/notbloom/ficus/watcher"
	"github.com/spf13/cobra"
)

var verbose bool

func init() {
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logging")
}

var rootCmd = &cobra.Command{
	Use:   "ficus",
	Short: "Ficus is a fast frontend development server",
	Long:  `A lightweight Go development server for rapid frontend development with Go HTML templates, HTMX, and Tailwind CSS.`,
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			log.SetLevel(log.DebugLevel)
		}

		// Read config file or get default
		cfg := config.GetConfig()

		log.Info("Starting Ficus", "url", "http://"+cfg.ServerAddr())
		log.Info("Press Ctrl+C to stop")

		// Setup file watcher
		w, err := watcher.New(cfg)
		if err != nil {
			log.Fatal("Failed to create watcher", "error", err)
		}

		w.OnChange = func(path string) {
			log.Info("Building Tailwind...")
			npm.RunTailwind(cfg)
			if broker := server.GetBroker(); broker != nil {
				broker.Reload()
			}
		}

		if err := w.Start(); err != nil {
			log.Fatal("Failed to start watcher", "error", err)
		}

		// Setup signal handling for graceful shutdown
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		// Start server in goroutine
		go server.StartServer(cfg)

		// Wait for shutdown signal
		<-sigChan
		log.Info("Shutting down...")

		// Graceful shutdown with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		w.Stop()
		if err := server.Shutdown(ctx); err != nil {
			log.Error("Shutdown error", "error", err)
		}

		log.Info("Goodbye!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
