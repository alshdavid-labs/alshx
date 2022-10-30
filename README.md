# CLI Utils

## Installation

### MacOS or Linux

```bash
wget https://github.com/alshdavid/alshx/releases/latest/download/alshx-darwin-arm64.tar.gz
rm -rf /usr/local/alshx
mkdir /usr/local/alshx
tar -C /usr/local/alshx -xzf alshx-darwin-arm64.tar.gz

export PATH=/usr/local/alshx:$PATH
```

### Windows 10 / 11

Run the following command in PowerShell:

```powershell
(New-Object System.Net.WebClient).DownloadString("https://raw.githubusercontent.com/alshdavid/alshx/master/.github/installers/install.ps1") | powershell -command -
```

Be sure to add `%USERPROFILE%\.local\alshx` to your `%PATH%` directory

## Updating

Once you have the utils installed, to update to the latest version just run

```bash
alshx update
```
