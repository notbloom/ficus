package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init [path]",
	Short: "Initialize a new Ficus project",
	Long:  `Initialize a new Ficus project with the standard folder structure and starter files.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		basePath := "."
		if len(args) > 0 {
			basePath = args[0]
		}

		if err := initProject(basePath); err != nil {
			log.Fatal("Failed to initialize project", "error", err)
		}

		log.Info("Project initialized successfully!")
		log.Info("Run 'ficus' to start the development server")
	},
}

func initProject(basePath string) error {
	// Create directory structure
	dirs := []string{
		"views",
		"views/assets",
		"views/assets/css",
		"views/assets/js",
		"views/pages",
		"views/components",
	}

	for _, dir := range dirs {
		path := filepath.Join(basePath, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create %s: %w", dir, err)
		}
		log.Info("Created directory", "path", path)
	}

	// Create files
	files := map[string]string{
		"tailwind.config.js":          tailwindConfigContent,
		"views/input.css":             inputCSSContent,
		"config.toml":                 configTOMLContent,
		"views/index.html":            indexHTMLContent,
		"views/layout.html":           layoutHTMLContent,
		"views/pages/example.html":    examplePageContent,
		"views/assets/css/style.css":  "", // Empty, Tailwind will populate
	}

	for filename, content := range files {
		path := filepath.Join(basePath, filename)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to create %s: %w", filename, err)
		}
		log.Info("Created file", "path", path)
	}

	return nil
}

const tailwindConfigContent = `/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.{html,js}"],
  theme: {
    extend: {},
  },
  plugins: [],
}
`

const inputCSSContent = `@tailwind base;
@tailwind components;
@tailwind utilities;
`

const configTOMLContent = `# Ficus Configuration
# All settings are optional - defaults shown below

[folders]
views_folder = "./views"
pages = "pages"
assets = "assets"
components = "components"
views_extension = ".html"

[server]
address = "127.0.0.1"
port = 8000
write_timeout = 15
read_timeout = 15

[watch]
include_suffix = [".html", ".css", ".js", ".json"]
exclude_suffix = ["~"]
`

const indexHTMLContent = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Ficus Dev Server</title>
    <link href="/assets/css/style.css" rel="stylesheet">
</head>
<body>
<script>
    const evtSource = new EventSource("/reload/stream");
    evtSource.onmessage = (event) => {
        console.log("Ficus: Reloading...");
        location.reload();
    };
</script>
<div class="container p-6">
    <h1 class="text-5xl py-2">Pages</h1>
    <ul>
    {{ range $index, $link := . }}
        <li>
            <a href="{{$link.Url}}" class="text-blue-600 hover:underline">{{$link.Title}}</a>
        </li>
    {{ end }}
    </ul>
</div>
</body>
</html>
`

const layoutHTMLContent = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <link href="/assets/css/style.css" rel="stylesheet">
</head>
<body>
<script>
    const evtSource = new EventSource("/reload/stream");
    evtSource.onmessage = (event) => {
        console.log("Ficus: Reloading...");
        location.reload();
    };
</script>
{{template "content" .}}
</body>
</html>
`

const examplePageContent = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Example Page</title>
    <link href="/assets/css/style.css" rel="stylesheet">
</head>
<body>
<script>
    const evtSource = new EventSource("/reload/stream");
    evtSource.onmessage = (event) => {
        console.log("Ficus: Reloading...");
        location.reload();
    };
</script>
<div class="container p-6">
    <h1 class="text-3xl font-bold">Hello Ficus!</h1>
    <p class="mt-4 text-gray-600">Edit this file and watch it reload automatically.</p>
    <a href="/" class="text-blue-600 hover:underline mt-4 inline-block">&larr; Back to pages</a>
</div>
</body>
</html>
`
