# Dependencies

- [Go](https://go.dev/doc/install)
- Make
  - [Make for Windows](https://gnuwin32.sourceforge.net/packages/make.htm)
  - [Make for MacOS](https://formulae.brew.sh/formula/make)
- GNU CoreUtils
  - [CoreUtils for Windows](https://gnuwin32.sourceforge.net/packages/coreutils.htm)
  - [CoreUtils for MacOS](https://formulae.brew.sh/formula/coreutils)

_Tested on PowerShell Core, Bash and zsh (MacOS and Linux)_

## Building

Compiling all the tools:

```bash
make

# Optimized build 
env PROD=true make
```

For a specific project

```bash
make build tool=alshx

# Or
cd tools/alshx && make build
```
