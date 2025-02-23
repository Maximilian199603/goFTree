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

# Create the tools directory if it doesn't exist
if (-not (Test-Path -Path $toolsDir)) {
    Write-Host "Creating tools directory at $toolsDir"
    New-Item -ItemType Directory -Force -Path $toolsDir
}

# Create the out directory if it doesn't exist
if (-not (Test-Path -Path $outDir)) {
    Write-Host "Creating output directory at $outDir"
    New-Item -ItemType Directory -Force -Path $outDir
}

# Create zip files for both the 32-bit and 64-bit executables
$exe32Zip = "$toolsDir\goFTree_32bit.zip"
$exe64Zip = "$toolsDir\goFTree_64bit.zip"

# Function to create zip files
function CreateZip {
    param (
        [string]$exePath,
        [string]$zipPath
    )

    Write-Host "Creating zip file for $exePath"
    if (Test-Path $exePath) {
        Compress-Archive -Path $exePath -DestinationPath $zipPath -Force
        Write-Host "Created zip file: $zipPath"
    } else {
        Write-Host "Error: File $exePath does not exist!"
        exit 1
    }
}

# Create the zip files for both the 32-bit and 64-bit executables
CreateZip $exe32Path $exe32Zip
CreateZip $exe64Path $exe64Zip

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

# Return to the original directory before exiting
Set-Location -Path $originalDirectory

Write-Host "Chocolatey package build process completed!"
