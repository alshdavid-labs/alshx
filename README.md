# CLI Utils

## Installation

### MacOS

```bash
rm -rf "${HOME}/.local/alshx"
mkdir -p "${HOME}/.local/alshx"
cd "${HOME}/.local/alshx"
wget https://github.com/alshdavid/alshx/releases/latest/download/darwin-amd64.zip
unzip darwin-amd64.zip .
rm -rf darwin-amd64.zip
```

Make sure to add `$HOME/.local/alshx` to your `$PATH`

```bash
export PATH="${PATH}:${HOME}/.local/alshx"
```

### Linux

```bash
rm -rf "${HOME}/.local/alshx"
mkdir -p "${HOME}/.local/alshx"
cd "${HOME}/.local/alshx"
wget https://github.com/alshdavid/alshx/releases/latest/download/linux-amd64.zip
unzip linux-amd64.zip .
rm -rf linux-amd64.zip
```

Make sure to add `$HOME/.local/alshx` to your `$PATH`

```bash
export PATH="${PATH}:${HOME}/.local/alshx"
```

### Windows

Run the following command in PowerShell:

```powershell
(New-Object System.Net.WebClient).DownloadString("https://raw.githubusercontent.com/alshdavid/alshx/master/install.ps1") | powershell -command -
```

Be sure to add `%USERPROFILE%\.local\alshx` to your `%PATH%` directory

## Updating

Once you have the utils installed, to update to the latest version just run

```bash
alshx update
```