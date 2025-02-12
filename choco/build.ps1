$ErrorActionPreference = "Stop"

# Define the paths
$packageVersion = "0.1.0"  # Define the version of the package
$packageDir = "choco/myapp"  # Path to the folder containing the .nuspec file
$toolsDir = "$packageDir/tools"  # Folder where the binary will be placed
$binaryFile = "$toolsDir/goftree.exe"  # Path to your prebuilt binary
$nuspecFile = "$packageDir/myapp.nuspec"  # Path to the .nuspec file

# Check if the binary exists
if (-Not (Test-Path $binaryFile)) {
    Write-Error "Prebuilt binary not found at $binaryFile"
    exit 1
}

# Ensure the tools directory exists
if (-Not (Test-Path $toolsDir)) {
    New-Item -ItemType Directory -Force -Path $toolsDir
}

# Copy the binary to the tools folder
Write-Host "Copying prebuilt binary to $toolsDir"
Copy-Item -Path $binaryFile -Destination $toolsDir -Force

# Run choco pack to create the .nupkg package
Write-Host "Packing Chocolatey package..."
choco pack $nuspecFile

Write-Host "Package created successfully!"