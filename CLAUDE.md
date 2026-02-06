# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Ficus is a lightweight Go development server for rapid frontend development with Go HTML templates, HTMX, and Tailwind CSS. It provides live file watching, hot reloading, and automatic Tailwind CSS compilation.

## Build and Run Commands

```bash
# Build the project
go build .

# Install globally
go install github.com/notbloom/ficus

# Initialize a new project (creates folders and starter files)
ficus init [path]

# Run the server
ficus [-v | --verbose]

# Stop gracefully with Ctrl+C
```

**Prerequisites**: Node.js/npm required for Tailwind CSS compilation via `npx tailwindcss`.

## Architecture

### Package Structure

- **cmd/**: CLI commands using Cobra
  - `root.go`: Main command - sets up watcher, signal handling, starts server
  - `init.go`: Scaffolds new project with complete folder structure and starter files

- **config/**: Unified configuration system
  - `Config` struct with `FoldersConfig`, `ServerConfig`, `WatchConfig`
  - Helper methods: `PagesPath()`, `AssetsPath()`, `ServerAddr()`, `InputCSSPath()`, `OutputCSSPath()`
  - Defaults used when no config.toml present

- **server/**: HTTP server using Gorilla Mux
  - `server.go`: Routes, template rendering, static file serving, graceful shutdown
  - `broker.go`: Server-Sent Events broker for live reload (uses context-based connection handling)
  - `homelist.handler.go`: Lists all HTML pages

- **watcher/**: File system watcher
  - `New(cfg)` creates watcher, `Start()` begins watching, `Stop()` cleans up
  - `OnChange` callback triggered on file modifications

- **npm/**: Tailwind CSS compilation wrapper

### Live Reload Flow

1. Watcher detects file change in pages/components directories
2. `npm.RunTailwind(cfg)` recompiles CSS
3. `broker.Reload()` called directly (no HTTP roundtrip)
4. Broker pushes SSE message to all connected clients at `/reload/stream`
5. Client-side EventSource triggers `location.reload()`

### Page Rendering

- Pages live in `{views_folder}/{pages}/{name}.html`
- Optional JSON data in `{views_folder}/{pages}/{name}.json` gets passed to template
- Template functions: `N(start, end)` ascending range, `M(start, end)` descending range
- Routes: `/pages/{page}/` renders templates, `/` shows page listing

## Configuration

Optional `config.toml` in project root (all settings have sensible defaults):

```toml
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
```

## Key Design Patterns

- **Config-driven**: All paths and settings flow from `config.GetConfig()`
- **Direct broker calls**: Watcher calls `broker.Reload()` directly instead of HTTP POST
- **Graceful shutdown**: Signal handling (SIGINT/SIGTERM) with proper cleanup
- **Callback-based watcher**: `watcher.OnChange` allows flexible response to file changes
