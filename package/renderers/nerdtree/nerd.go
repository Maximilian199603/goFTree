package nerdtree

import (
	"strings"

	"github.com/EdgeLordKirito/goFTree/package/filetree"
	"github.com/EdgeLordKirito/goFTree/package/renderers/generaltree"
)

// fileIcons maps file extensions to their respective Nerd Font icons.
var fileIcons = map[string]string{
	//Programming filetypes
	".go":     " ", // Go
	".py":     " ", // Python
	".js":     " ", // JavaScript
	".ts":     " ", // TypeScript
	".html":   " ", // Html
	".css":    " ", // CSS
	".php":    "󰌟 ", // Php
	".java":   " ", // Java
	".jar":    " ", // Java archive
	".rb":     " ", // Ruby
	".rs":     " ", // Rust
	".cpp":    " ", // C++
	".c":      " ", // C
	".h":      " ", // C Header
	".sln":    "󰘐 ", // Solution
	".cs":     " ", // C Sharp file
	".md":     " ", // Markdown
	".json":   " ", // JSON
	".yaml":   " ", // YAML
	".toml":   " ", // TOML
	".xml":    "󰗀 ", // XML
	".log":    "󱂅 ", // Log file
	".bash":   " ", // Bash script
	".sh":     " ", // Shell script
	".zsh":    " ", // Z Shell script
	".fish":   " ", // Fish Shell script
	".ksh":    " ", // K Shell script
	".elv":    " ", // Elv script
	".txt":    " ", // Text file
	".csv":    " ", // CSV file
	".sql":    " ", // SQL file
	".db":     " ", // Database file
	".mdb":    " ", // MS Access file
	".sqlite": " ", // SQLite file
	".nix":    " ", // Nix file

	//Powershell filetypes
	".ps1":    " ", // Powershell
	".psm1":   " ", // Powershell
	".psd1":   " ", // Powershell
	".ps1xml": " ", // Powershell
	".psc1":   " ", // Powershell

	//Batch filetypes
	".bat": " ", // Batch
	".cmd": " ", // Command
	".btk": " ", // Batch

	//Archive filetypes
	".zip": " ", // ZIP Archive
	".tar": " ", // TAR Archive
	".gz":  " ", // GZ Archive
	".rar": " ", // Rar Archive
	".7z":  " ", // 7z Archive

	// Other common file types
	".exe": " ", // Executable (Windows EXE)
	".msi": " ", // Windows Installer
	".elf": " ", // Elf
	".bin": " ", // Bin
	".app": " ", // macOS application
	".iso": " ", // ISO disk image

	// Image files
	".jpg":  " ", // JPG image (opaque)
	".jpeg": " ", // JPEG image (opaque)
	".png":  " ", // PNG image (with transparency)
	".gif":  " ", // GIF image (with transparency)
	".webp": " ", // Webp image (with transparency)
	".bmp":  " ", // BMP image (opaque)
	".tiff": " ", // TIFF image (opaque)
	".svg":  " ", // SVG image (opaque)

	// Audio files
	".mp3":  " ", // MP3 audio
	".m4a":  " ", // MP3 audio
	".wav":  " ", // WAV audio
	".flac": " ", // FLAC audio
	".aac":  " ", // AAC audio
	".ogg":  " ", // AAC audio

	// Video files
	".mp4":  " ", // MP4 video
	".avi":  " ", // AVI video
	".mkv":  " ", // MKV video
	".mov":  " ", // MOV video
	".webm": " ", // MOV video

	//Office filetypes
	".pdf":  " ", // PDF file
	".doc":  " ", // Microsoft Word document
	".docx": " ", // Microsoft Word document
	".xls":  "󱎏 ", // Microsoft Excel spreadsheet
	".xlsx": "󱎏 ", // Microsoft Excel spreadsheet
	".ppt":  "󱎐 ", // PowerPoint presentation
	".pptx": "󱎐 ", // PowerPoint presentation
	".rtf":  "󰊄 ", // Rich Text Format
	".odt":  " ", // Open Document Text
	".ods":  "󰧷 ", // Open Document Spreadsheet
	".odp":  "󰐨 ", // Open Document Presentation

	//Fonts
	".ttf":  " ", //True Type Font
	".otf":  " ", //Open Type Font
	".woff": " ", //Web Open Font Format
}

const (
	// directoryIcon is the Nerd Font icon for directories.
	directoryIcon = " "
	// genericFileIcon is used when no specific file icon is found.
	genericFileIcon = " "
)

type Engine struct{}

// Render generates the Nerd Font file tree as a string.
func (e Engine) Render(tree *filetree.FileTree) (string, error) {

	set := &generaltree.RenderSettings{
		DirTJunction:  "├── ",
		DirLJunction:  "└── ",
		FileTJunction: "├── ",
		FileLJunction: "└── ",
		NoJunction:    "│   ",
		Empty:         "    ",

		DirPrepender:  func(s string, n *filetree.Node) (string, *filetree.Node) { return getIcon(n) + s, n },
		DirAppender:   func(s string, n *filetree.Node) (string, *filetree.Node) { return s + "/", n },
		FilePrepender: func(s string, n *filetree.Node) (string, *filetree.Node) { return getIcon(n) + s, n },
		FileAppender:  generaltree.Noop,
	}
	return generaltree.Render(tree, set)
}

// getIcon returns the appropriate Nerd Font icon for the given node.
func getIcon(node *filetree.Node) string {
	if node.IsDir {
		return directoryIcon
	}
	for ext, icon := range fileIcons {
		if strings.HasSuffix(node.Name, ext) {
			return icon
		}
	}
	return genericFileIcon
}
