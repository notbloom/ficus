![ficus](ficus-banner.png)
# Ficus - Quick Go + HTMX + Tailwind Frontend-Dev

A lightweight Go development server for rapid frontend prototyping with Go HTML templates, HTMX, and Tailwind CSS.

## Features
- Local dev server with configurable address and port
- File watcher on pages and components directories
- Hot reloading via Server-Sent Events
- Automatic Tailwind CSS compilation
- Go HTML templates with optional JSON data
- Graceful shutdown (Ctrl+C)

## Install

```
go install github.com/notbloom/ficus
```

**Prerequisites**: Node.js/npm required for Tailwind CSS compilation (`npx tailwindcss`).

## Quick Start

```bash
# Create a new project
mkdir my-project && cd my-project
ficus init

# Start the dev server
ficus
```

Open http://127.0.0.1:8000 to see your pages. Edit any file in `views/pages/` and the browser reloads automatically.

## Commands

Start the server:
```
ficus [-v | --verbose]
```

Initialize a new project with folder structure and starter files:
```
ficus init [path]
```

## Project Structure

`ficus init` creates the following structure:

```
project-folder
├── views
│   ├── assets
│   │   ├── css
│   │   │   └── style.css           // Tailwind generated CSS
│   │   └── js
│   ├── pages
│   │   ├── [page].html             // Pages served at /pages/[page]/
│   │   └── [page].json             // Optional data passed to template
│   ├── components
│   ├── index.html                  // Page listing (served at /)
│   ├── layout.html                 // Base layout template
│   └── input.css                   // Tailwind source CSS
├── config.toml                     // Server configuration (optional)
└── tailwind.config.js              // Tailwind configuration
```

## Configuration

All settings are optional. Create a `config.toml` in your project root to override defaults:

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

## Template Functions

Templates have access to these helper functions:

- `N(start, end)` - generates an ascending range of integers
- `M(start, end)` - generates a descending range of integers

## How It Works

1. File watcher monitors `views/pages/` and `views/components/`
2. On file change, Tailwind CSS is recompiled automatically
3. The SSE broker notifies all connected browsers to reload
4. Pages are rendered using Go's `html/template` with optional JSON data
