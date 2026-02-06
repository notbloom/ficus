package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Folders FoldersConfig `toml:"folders"`
	Server  ServerConfig  `toml:"server"`
	Watch   WatchConfig   `toml:"watch"`
}

type FoldersConfig struct {
	ViewsFolder    string `toml:"views_folder"`
	Assets         string `toml:"assets"`
	Components     string `toml:"components"`
	Pages          string `toml:"pages"`
	ViewsExtension string `toml:"views_extension"`
}

type ServerConfig struct {
	Address      string `toml:"address"`
	Port         int    `toml:"port"`
	WriteTimeout int    `toml:"write_timeout"`
	ReadTimeout  int    `toml:"read_timeout"`
}

type WatchConfig struct {
	IncludeSuffix []string `toml:"include_suffix"`
	ExcludeSuffix []string `toml:"exclude_suffix"`
}

// Singleton config instance
var cfg *Config

func GetConfig() *Config {
	if cfg != nil {
		return cfg
	}
	cfg = loadConfig()
	return cfg
}

func loadConfig() *Config {
	defaults := defaultConfig()

	dat, err := os.ReadFile("./config.toml")
	if err != nil {
		log.Debug("No config.toml found, using defaults")
		return defaults
	}

	var fileCfg Config
	if err := toml.Unmarshal(dat, &fileCfg); err != nil {
		log.Warn("Error parsing config.toml, using defaults", "error", err)
		return defaults
	}

	// Merge file config over defaults
	return mergeConfig(defaults, &fileCfg)
}

func defaultConfig() *Config {
	return &Config{
		Folders: FoldersConfig{
			ViewsFolder:    "./views",
			Assets:         "assets",
			Components:     "components",
			Pages:          "pages",
			ViewsExtension: ".html",
		},
		Server: ServerConfig{
			Address:      "127.0.0.1",
			Port:         8000,
			WriteTimeout: 15,
			ReadTimeout:  15,
		},
		Watch: WatchConfig{
			IncludeSuffix: []string{".html", ".css", ".js", ".json"},
			ExcludeSuffix: []string{"~"},
		},
	}
}

func mergeConfig(defaults, file *Config) *Config {
	result := *defaults

	// Merge Folders
	if file.Folders.ViewsFolder != "" {
		result.Folders.ViewsFolder = file.Folders.ViewsFolder
	}
	if file.Folders.Assets != "" {
		result.Folders.Assets = file.Folders.Assets
	}
	if file.Folders.Components != "" {
		result.Folders.Components = file.Folders.Components
	}
	if file.Folders.Pages != "" {
		result.Folders.Pages = file.Folders.Pages
	}
	if file.Folders.ViewsExtension != "" {
		result.Folders.ViewsExtension = file.Folders.ViewsExtension
	}

	// Merge Server
	if file.Server.Address != "" {
		result.Server.Address = file.Server.Address
	}
	if file.Server.Port != 0 {
		result.Server.Port = file.Server.Port
	}
	if file.Server.WriteTimeout != 0 {
		result.Server.WriteTimeout = file.Server.WriteTimeout
	}
	if file.Server.ReadTimeout != 0 {
		result.Server.ReadTimeout = file.Server.ReadTimeout
	}

	// Merge Watch
	if len(file.Watch.IncludeSuffix) > 0 {
		result.Watch.IncludeSuffix = file.Watch.IncludeSuffix
	}
	if len(file.Watch.ExcludeSuffix) > 0 {
		result.Watch.ExcludeSuffix = file.Watch.ExcludeSuffix
	}

	return &result
}

// Helper methods for path construction

func (c *Config) PagesPath() string {
	return filepath.Join(c.Folders.ViewsFolder, c.Folders.Pages)
}

func (c *Config) AssetsPath() string {
	return filepath.Join(c.Folders.ViewsFolder, c.Folders.Assets)
}

func (c *Config) ComponentsPath() string {
	return filepath.Join(c.Folders.ViewsFolder, c.Folders.Components)
}

func (c *Config) ServerAddr() string {
	return fmt.Sprintf("%s:%d", c.Server.Address, c.Server.Port)
}

func (c *Config) InputCSSPath() string {
	return filepath.Join(c.Folders.ViewsFolder, "input.css")
}

func (c *Config) OutputCSSPath() string {
	return filepath.Join(c.AssetsPath(), "css", "style.css")
}
