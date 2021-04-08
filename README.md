# CLI Utils

## Installation

### MacOS

```bash
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

### Windows (INSTALL INSTRUCTIONS TODO)

```powershell
mkdir "%USERPROFILE%\.bin"
cd "%USERPROFILE%\.bin"
wget https://github.com/alshdavid/alshx/releases/latest/download/windows-amd64.zip
compact /u "windows-amd64.zip" /i /Q
rm windows-amd64.zip
```