# goFTree - A CLI Tool to Display Filesystem Structure

`goFTree` is a command-line tool that displays the structure of a given directory, generating a tree-like view of files and subdirectories. The output can be customized in various formats such as ASCII, Markdown, XML, and more. 

## Features

- **Filesystem Traversal**: Walk through directories and list their contents recursively.
- **Multiple Output Styles**: Render the directory structure in ASCII, Markdown, XML, JSON, or other custom formats.
- **File Output**: Save the generated output to a file or print it directly to the console.
- **Interactive Confirmation**: Prevent overwriting files unless confirmed by the user.
--- 
## Usage
```bash
goFTree [path] [flags]
```
---
## Flags

- `--style`, `-s` (default: `ascii`): The style in which the tree will be rendered. Options are:
    - `ascii`: Plain text with indentation.
    - `markdown`: Markdown format.
    - `xml`: XML format.
    - `json`: JSON format.
    - `line`: Line-based structure.
    - `dashed`: Line-based structure with dashed borders.
  
- `--file`, `-f`: The file to which the output will be written. If the file exists, you will be asked for confirmation before overwriting it.
