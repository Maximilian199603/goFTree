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
- `--include-hidden`, `-i`: Wether to include hidden files and or directories. Options are:
	- all: include all hidden entries
	- files: only includes hidden files
	- dirs: only includes hidden directories
	- none: do not include hidden entries
- `--version`, `-v`: Prints the installed version

## Installation

### **General Installation**

Before installing **GoFTree**, verify that the file matches the provided **SHA256 checksum** to ensure its integrity.

1. Find the checksum on the [release page](#your-release-page-url).
2. Use the following commands to check the file:
3. Compare the result with the release page checksum. If they match, proceed with the installation.

**Windows**:
``` powershell
certutil -hashfile <path_to_installer> SHA256
```
**Linux/Mac**:
``` sh
sha256sum <path_to_installer>
```
### **Windows**

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
### **Installing GoFTree with Chocolatey**

You can easily install **GoFTree** using the **Chocolatey** package manager. Follow these steps:
#### Prerequisites

If you don't have Chocolatey installed, follow the installation instructions on their official website: [Chocolatey Installation Guide](https://chocolatey.org/install).
#### Installing GoFTree

Once Chocolatey is installed, you can install **GoFTree** by running the following command in an elevated PowerShell or Command Prompt window:
``` powershell
choco install goftree
```
#### Verifying the Installation

After installation, verify that **GoFTree** was successfully installed by running the following command:
``` powershell
goftree --version
```
### **Linux and MacOS**
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

## Build Process

### Compiling GoFTree Binaries
To build the GoFTree binaries using the provided `Dockerfile`, follow the steps below:
#### 1. Build the Docker Image

First, build the Docker image from the `Dockerfile`:

``` bash
docker build -t goftreebuilder .
```

This will create a Docker image named `goftreebuilder`, which contains the necessary environment to build the GoFTree binaries.

#### 2. Run the Docker Container

Once the image is built, you can run the Docker container to build the binaries. Use the following command:
``` bash
docker run --rm -e VERSION=<versionstring> -v "<yourLocalOutputDir>:/output" goftreebuilder
```
##### Explanation of Placeholders:

- **`<versionstring>`**:
    - This placeholder represents the **version** of your build and can be replaced with any UTF-8 string (e.g., `v1.0.0`, `latest`, `2025-02-22`, etc.).
    - Example: `-e VERSION=v1.0.0`
- **`<yourLocalOutputDir>`**:
    - This placeholder represents a **path on your local machine** where the build binaries will be saved.
    - The `-v` flag **mounts** this directory onto the `/output` directory inside the container, so after the build process completes, you can access the compiled binaries in your local output directory.
    - **Note**: Make sure the directory you specify exists on your local machine.
    - Example: `-v "/path/to/output:/output"`

### Building the Chocolatey Package (`buildchoco.ps1`)

Run the `buildchoco.ps1` script from the **root of the project** to create a Chocolatey package.

#### Usage

Execute the following command in PowerShell from the root of the project:
``` powershell
.\buildchoco.ps1 "<path_to_32bit_exe>" "<path_to_64bit_exe>" "<version>"
```
- `<path_to_32bit_exe>`: Path to the 32-bit GoFTree executable.
- `<path_to_64bit_exe>`: Path to the 64-bit GoFTree executable.
- `<version>`: The version for the package (e.g., `1.0.0`).

#### Process

- The script will create ZIP files for both executables in `choco/goFTree/tools/`.
- It will update the `goftree.nuspec` version and run `choco pack`.
- The output `.nupkg` file will be placed in `choco/out/`.