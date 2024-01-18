package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var folders = []string{
	"views",
	"views/assets",
	"views/assets/css",
	"views/assets/js",
	"views/pages",
	"views/components",
}
var files = []string{""}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate template project",
	Long:  `Generate template project`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Writing tailwind.config.js")

		val := "/** @type {import('tailwindcss').Config} */\nmodule.exports = {\n  content: [\"./**/*.{html,js}\"],\n  theme: {\n    extend: {},\n  },\n  plugins: [],\n}\n"
		data := []byte(val)

		err := os.WriteFile("./tailwind.config.js", data, 0)

		if err != nil {
			log.Fatal(err)
		}

		log.Info("Writing ./views/input.css")

		val = "@tailwind base;\n@tailwind components;\n@tailwind utilities;"
		data = []byte(val)

		err = os.WriteFile("./views/input.css", data, 0)

		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println("done")
	},
}
