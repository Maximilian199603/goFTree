param (
    [string]$exe32Path,    # Path to the 32-bit executable
    [string]$exe64Path,    # Path to the 64-bit executable
    [string]$version       # Version string
)

# Save the starting directory to return to later
$originalDirectory = Get-Location

# Ensure all parameters are provided
if (-not $exe32Path -or -not $exe64Path -or -not $version) {
    Write-Host "Usage: .\build-choco-package.ps1 -exe32Path <path_to_32bit_exe> -exe64Path <path_to_64bit_exe> -version <version>"
    exit 1
}

# Define the base directory for the zip files and the Chocolatey package files
$currentDir = Get-Location  # Get the current directory (absolute path)
$chocoDir = "$currentDir\choco\goFTree"
$toolsDir = "$chocoDir\tools"
$outDir = "$currentDir\choco\out"  # Absolute output directory for the .nupkg file
$tempDir = "$currentDir\choco\temp"  # Temporary directory for handling executables

# Create the necessary directories
$directories = @($toolsDir, $outDir, $tempDir)
foreach ($dir in $directories) {
    if (-not (Test-Path -Path $dir)) {
        Write-Host "Creating directory: $dir"
        New-Item -ItemType Directory -Force -Path $dir
    }
}

# Create 32bit and 64bit directories inside the temp directory
$exe32TempDir = "$tempDir\32bit"
$exe64TempDir = "$tempDir\64bit"
New-Item -ItemType Directory -Force -Path $exe32TempDir
New-Item -ItemType Directory -Force -Path $exe64TempDir

# Function to copy and rename the executable
function CopyAndRenameExe {
    param (
        [string]$exePath,
        [string]$destinationDir
    )
    
    if (Test-Path $exePath) {
        $destinationExePath = "$destinationDir\goftree.exe"
        Write-Host "Copying $exePath to $destinationExePath"
        Copy-Item -Path $exePath -Destination $destinationExePath -Force
    } else {
        Write-Host "Error: File $exePath does not exist!"
        exit 1
    }
}

# Copy and rename both the 32bit and 64bit executables
CopyAndRenameExe $exe32Path $exe32TempDir
CopyAndRenameExe $exe64Path $exe64TempDir

# Create zip files for the 32-bit and 64-bit executables
$exe32Zip = "$toolsDir\goFTree_32bit.zip"
$exe64Zip = "$toolsDir\goFTree_64bit.zip"

# Function to create zip files
function CreateZip {
    param (
        [string]$sourceExePath,
        [string]$zipPath
    )

    Write-Host "Creating zip file for $sourceExePath"
    Compress-Archive -Path $sourceExePath -DestinationPath $zipPath -Force
    Write-Host "Created zip file: $zipPath"
}

# Create the zip files for the executables only (not the entire directories)
CreateZip "$exe32TempDir\goftree.exe" $exe32Zip
CreateZip "$exe64TempDir\goftree.exe" $exe64Zip

# Modify the version in the goftree.nuspec file
$nuspecPath = "$chocoDir\goftree.nuspec"
if (Test-Path $nuspecPath) {
    Write-Host "Updating version in $nuspecPath to $version"
    (Get-Content $nuspecPath) -replace '(?<=<version>)(.*?)(?=</version>)', $version | Set-Content $nuspecPath
    Write-Host "Version updated in goftree.nuspec"
} else {
    Write-Host "Error: goftree.nuspec file not found!"
    exit 1
}

# Change directory to the choco\goFTree folder and run choc pack with the specified output directory
Set-Location -Path $chocoDir

Write-Host "Running choco pack..."
$chocoPackOutput = choco pack --output-directory $outDir
$chocoPackOutput | ForEach-Object { Write-Host $_ }

# Cleanup: Remove the temp directory and its contents
Write-Host "Cleaning up temporary files in $tempDir"
Remove-Item -Path $tempDir -Recurse -Force
Write-Host "Temporary files cleaned up."

# Return to the original directory before exiting
Set-Location -Path $originalDirectory

Write-Host "Chocolatey package build process completed!"
