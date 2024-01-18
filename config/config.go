package config

import (
	"github.com/charmbracelet/log"
	"github.com/pelletier/go-toml/v2"
	"os"
)

type FoldersConfig struct {
	ViewsFolder string `toml:"views_folder"`
	Assets      string
	Components  string
	Pages       string

	ViewsRoute     string `toml:"views_route"`
	ViewsExtension string `toml:"views_extension"`
}

type ServerConfig struct {
	Address      string
	Port         int8
	WriteTimeout int8
	ReadTimeout  int8
}

func GetConfig() FoldersConfig {
	dat, err := os.ReadFile("./config.toml")
	if err != nil {
		//panic(err)
		log.Info("Default config ( use init to create toml .file) ")
	}

	//fmt.Print(string(dat))
	var cfg FoldersConfig

	err = toml.Unmarshal(dat, &cfg)
	if err != nil {
		panic(err)
	}
	/*log.Print("Port:", cfg.Port)
	log.Print("ViewsFolder:", cfg.ViewsFolder)
	log.Print("ViewsExtension:", cfg.ViewsExtension)*/
	return cfg
}
func defaultConfig() FoldersConfig {
	return FoldersConfig{
		ViewsFolder:    "",
		ViewsExtension: ".html",
		ViewsRoute:     "frontend/views/layout",
	}
}
