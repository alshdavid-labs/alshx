clear 

Write-Output "Alshx Installer for Windows"
Write-Output ""

$link = "https://github.com/alshdavid/alshx/releases/latest/download"
$filename = $null
$outdir = "$env:Programfiles\alshdavid\alshx"

if ($env:PROCESSOR_ARCHITECTURE -eq "AMD64") {
  $script:filename = "alshx-windows-amd64.tar.gz"
}
if ($env:PROCESSOR_ARCHITECTURE -eq "ARM64") {
  $script:filename = "alshx-windows-arm64.tar.gz"
}

if ($script:filename -eq $null) {
  Write-Output "Platform $env:PROCESSOR_ARCHITECTURE Not Supported"
  exit 1
}

Write-Output "Target Directory: $script:outdir"
Write-Output ""


if (Test-Path -Path "$script:outdir") {
  Write-Output "Tools Detected, Upgrading"
  Remove-Item -Path $script:outdir -Recurse -Force | Out-Null
  New-Item -Path $script:outdir -ItemType "directory" | Out-Null
} else {
  Write-Output "Installing Tools"
  New-Item -Path $script:outdir -ItemType "directory" | Out-Null
}

try {
  Write-Output "Downloading Tools"
  (New-Object System.Net.WebClient).DownloadFile("$script:link/$script:filename", "$env:temp\$script:filename")
  tar -C "$script:outdir" -xzf "$env:temp\$script:filename"
} finally {
  Remove-Item -Path "$env:temp\$script:filename" -Recurse -Force | Out-Null
}

if (Test-Path -Path "$script:outdir") {
} else {
  Write-Output "Error: Tools failed to install"
  exit 1
}

Write-Output ""

$oldPath = $([Environment]::GetEnvironmentVariable("PATH", "Machine"))

if ("$script:oldPath" -like "*$script:outdir*") {
  Write-Output "System PATH: Detected - Skipping"
  Write-Output ""
} else {
  Write-Output "System PATH: Detected - Installing"
  Write-Output ""
  $newPath = "$script:oldPath" + [IO.Path]::PathSeparator + "$script:outdir"
  [Environment]::SetEnvironmentVariable( "Path", $script:newPath, "Machine")
} 

Write-Output "Installation Complete"
Write-Output "Please restart terminal to see changes"
Write-Output ""
