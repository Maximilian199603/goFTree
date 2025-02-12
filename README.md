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
    - `nerd`: Uses Nerd Font icons to enhance the tree visualization with symbols.
  
- `--file`, `-f`: The file to which the output will be written. If the file exists, you will be asked for confirmation before overwriting it.
- `--version`, `-v`: Prints the installed version

## Installation
### Windows

1. Download the `goftree.exe` binary from the [releases page](https://github.com/EdgeLordKirito/goFTree/releases).
2. Move the selected binary  to a directory included in your `PATH` (e.g., `C:\Program Files\goftree\`).
3. Optionally, set a persistent PowerShell alias (see below).
#### Setting a Persistent PowerShell Alias
To create an alias so `goftree` can be used from anywhere:
1\. Open PowerShell and run:

```powershell
notepad.exe $PROFILE
```

2\. Add the following line:

```powershell
Set-Alias -Name goftree -Value "C:\Program Files\goftree\goftree.exe"
```

3\. Save and close Notepad.
4\. Restart PowerShell or run `. $PROFILE` to apply changes.
### Linux and MacOS
1\. Download the appropriate `goftree` binary for your OS and architecture from the [releases page](https://github.com/EdgeLordKirito/goFTree/releases)  
2\. Move it to `/usr/local/bin` and give it execute permissions:

```sh
sudo mv goftree /usr/local/bin/
sudo chmod +x /usr/local/bin/goftree
```

3\. (Optional) Set an alias by adding this line to `~/.bashrc`, `~/.bash_profile`, or `~/.zshrc` (for Zsh users):

```sh
alias goftree="/usr/local/bin/goftree"
```

4\. Apply changes by running:

```sh
source ~/.bashrc
```
#### Nerd style
To use the `nerd` style in `goFTree`, you'll need to have a Nerd Font installed. Visit the [Nerd Fonts GitHub page](https://github.com/ryanoasis/nerd-fonts) to download a font and install it according to your operating system's method.

Once installed, set your terminal to use the Nerd Font, and the `nerd` style in `goFTree` will display enhanced icons.
