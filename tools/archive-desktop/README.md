# Archive Desktop Utility
## Keeping your desktop clean

### Usage

The following command will move all the files on your desktop to a predesignated folder with a timestamp.

e.g. The default behaviour will be to move the contents of `$HOME/Desktop` to `$HOME/Archive/YY-MM-DD.zip`

```bash
archive-desktop
```

#### No Compression

To avoid compressing the files, use the following arguments

```bash
archive-desktop --no-compress
archive-desktop --nc
```

#### Env Variables

To adjust the locations, set the following environment variables. The defaults will be automatically detected based on your OS.

```bash
ALSHX_DESKTOP_PATH="$HOME/Desktop"
ALSHX_DESKTOP_ARCHIVE_PATH="$HOME/Archive"
```

#### Pre Command

If you want to run a command on your desktop before archiving, use the following

```bash
archive-desktop --pre "rrm --find node_modules ."
```

Or set a default pre command with the following environment variable

```bash
ALSHX_DESKTOP_ARCHIVE_PRE_COMMAND="rrm --find node_modules ."
```