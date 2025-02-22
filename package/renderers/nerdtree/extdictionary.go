package nerdtree

import ()

var fileExtensionToIcon = map[string]string{
	//Programming filetypes
	// Web
	".html": " ", // Html
	".css":  " ", // CSS

	// OOP Languages
	".java":  " ", // Java
	".cs":    " ", // C Sharp file
	".cpp":   " ", // C++
	".swift": "󰛥 ", // Swift
	".kt":    "󱈙 ", // Kotlin
	".dart":  " ", // Dart
	".ts":    " ", // TypeScript
	".php":   "󰌟 ", // Php
	".rb":    " ", // Ruby

	// Imperative Languages
	".go": " ", // Go
	".c":  " ", // C
	".h":  " ", // C Header
	".js": " ", // JavaScript
	".py": " ", // Python
	//".m":   " ", // Matlab
	//".m":  " ", // Objective-C
	//".p":  "",   // Pascal

	// Functional Languages
	".nix":   " ", // Nix file
	".rs":    " ", // Rust
	".clj":   " ", // Clojure
	".scala": " ", // Scala
	".sc":    " ", // Scala
	".elm":   " ", // Elm
	".hs":    " ", // Haskell
	".ml":    " ", // OCaml
	".scm":   " ", // Scheme
	".ex":    " ", // Elixir
	".exs":   " ", // Elixir
	".erl":   " ", // Erlang
	".hrl":   " ", // Erlang
	".r":     " ", // R
	//".gleam": "", // Gleam
	//".rkt": "",   // Racket

	// Config Filetypes
	".json":   " ", // JSON
	".xml":    "󰗀 ", // XML
	".sln":    "󰘐 ", // Solution
	".csproj": " ", // C Sharp file
	".yaml":   " ", // YAML
	".yml":    " ", // YAML
	".toml":   " ", // TOML
	".ini":    " ", // INI

	//Shell scripts
	".bash": " ", // Bash script
	".sh":   " ", // Shell script
	".zsh":  " ", // Z Shell script
	".fish": " ", // Fish Shell script
	".ksh":  " ", // K Shell script
	".elv":  " ", // Elvish script

	//Windows Shell scripts
	".ps1":    " ", // Powershell
	".psm1":   " ", // Powershell
	".psd1":   " ", // Powershell
	".ps1xml": " ", // Powershell
	".psc1":   " ", // Powershell
	".bat":    " ", // Batch
	".cmd":    " ", // Command
	".btk":    " ", // Batch

	//Database filetypes
	".sql":    " ", // SQL file
	".db":     " ", // Database file
	".mdb":    " ", // MS Access file
	".sqlite": " ", // SQLite file

	// Programming adjacent filetypes
	".csv": " ", // CSV file
	".log": "󱂅 ", // Log file
	".md":  " ", // Markdown

	// Executable filetypes
	".exe": " ", // Executable (Windows EXE)
	".msi": " ", // Windows Installer
	".elf": " ", // Elf
	".bin": " ", // Bin
	".app": " ", // macOS application
	".iso": " ", // ISO disk image

	//Archive filetypes
	".zip": " ", // ZIP Archive
	".tar": " ", // TAR Archive
	".gz":  " ", // GZ Archive
	".rar": " ", // Rar Archive
	".7z":  " ", // 7z Archive
	".jar": " ", // Java archive

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

	//Common filetypes
	".txt": " ", // Text file
	".rtf": " ", // Rich Text Format
	".pdf": " ", // PDF file

	//MS Office
	".doc":  " ", // Microsoft Word document
	".docx": " ", // Microsoft Word document
	".xls":  "󱎏 ", // Microsoft Excel spreadsheet
	".xlsx": "󱎏 ", // Microsoft Excel spreadsheet
	".ppt":  "󱎐 ", // PowerPoint presentation
	".pptx": "󱎐 ", // PowerPoint presentation

	//Open Office
	".odt": " ", // Open Document Text
	".ods": "󰧷 ", // Open Document Spreadsheet
	".odp": "󰐨 ", // Open Document Presentation

	//Fonts
	".ttf":   " ", //True Type Font
	".otf":   " ", //Open Type Font
	".woff":  " ", //Web Open Font Format
	".woff2": " ", //Web Open Font Format
}
