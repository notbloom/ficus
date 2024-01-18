# Quick Go + Htmx + Tailwind  Frontend-Dev

Starts a local server to test htmx pages + tailwind designs.
Install:
```
go install github.com/notbloom/frontend-dev
```
# Features
- Quick local server 
- Filewatcher 
- Hotreloading
- Auto recopile Tailwind
- Go Html Templates

# Commands

Start the server with defaults:
```
fontend-dev [ -v | --verbose ]
```
Create a project template:
```
frontend-dev init [ path ]
```

## Usage: 

# Directories

```
project-folder
├── views 
│   ├── assets
│   │   ├── css
│   │   │    └── style.css        // Tailwind generated css
│   │   ├── js
│   │   │   └── reload.js        // Reload script 
│   ├── pages
│   │   ├── [page].html          // Pages served
│   │   └── [page].json          // Optional data passed to template 
│   ├── components
│   │   ├── [component].html     // Components (no layout) served
│   │   └── [component].json     // Optional data passed to template 
│   ├── index.html               // Page + components listing
│   └── layout.html              // Layout for pages
├── config.toml                  // Server configuration
└── tailwind.config.js           // Tailwind configuration
```
# Todo
- [ ] Init doesnt create assets folder or style.css
- [ ] Correct folder structure
- [ ] Include layout in files
- [ ] Inject reload code in pages