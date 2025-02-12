$ErrorActionPreference = "Stop"

$package = Get-Package -Name goftree -ErrorAction SilentlyContinue

if ($package) {
    Write-Host "Package is already installed"
} else {
    Write-Host "Installing goftree..."
    $binaryPath = Join-Path $env:ChocolateyPackageFolder "tools\goftree.exe"
    Write-Host "Copying binary to install location"
    Copy-Item $binaryPath -Destination "$env:ProgramFiles\goftree\goftree.exe"
    Write-Host "Installation completed."
}