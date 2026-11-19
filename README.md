# Imp

<p align="center">
  <img src="./img/imp.png" alt="Imp" width="225">
</p>

**A tiny familiar that conjures source file stubs.**

`imp` creates new files pre-filled with a minimal, runnable starter for their language: a hello world for code, a skeleton for HTML/CSS, sensible defaults for config files. You can pass one or more file names. The stub is chosen by the file's extension. Scripts get a shebang and executable permissions, and existing files are left untouched unless you pass the `-f` or `--force` flag.

## Install

### Prebuilt binaries

Download the archive for your platform (Linux, macOS, Windows) from the [releases page](https://github.com/SamuelZima/imp/releases/latest), extract it and put `imp` somewhere on your `PATH`:

```sh
tar -xzf imp_*_darwin_arm64.tar.gz
sudo mv imp /usr/local/bin/
```

On macOS, if Gatekeeper blocks the binary, run `xattr -d com.apple.quarantine /usr/local/bin/imp`.

### With Go

```sh
go install github.com/SamuelZima/imp@latest
```

## Usage

```
imp [-h|--help] [-v|--version] [-l|--list] [-f|--force] <file_1> [file_2 ...]
```

| Flag | Description |
| --- | --- |
| `-h`, `--help` | Show help |
| `-v`, `--version` | Show version |
| `-l`, `--list` | List supported file extensions |
| `-f`, `--force` | Overwrite existing files |

The stub is picked by file extension. Stubs with a shebang (scripts) are created executable.

## Examples

```sh
imp main.go                  # Go hello world
imp main.c util.cpp app.py   # several files at once
imp .gitignore .editorconfig # config files
imp -f main.rs               # overwrite an existing file
imp -l                       # list supported extensions
```

```sh
$ imp hello.py && ./hello.py
Conjured: hello.py
>:3
```

```sh
$ imp index.html server_logic.php script.js .gitignore
Conjured: index.html
Conjured: server_logic.php
Conjured: script.js
Conjured: .gitignore
```

## Supported extensions

- **Compiled:** `.go` `.c` `.cpp` `.rs` `.zig` `.nim` `.hs` `.kt` `.swift` `.java` `.cs` `.dart`
- **Interpreted:** `.py` `.js` `.ts` `.php` `.rb` `.pl` `.lua` `.r` `.sh` `.zsh` `.fish` `.ps1`
- **Other:** `.md` `.html` `.css` `.gitignore` `.editorconfig`

## Notes

- Existing files are never overwritten without `-f`.
- Output is coloured with ANSI escape codes: successes in green, errors (unknown flags, unsupported extensions, write failures) in red.

## License

MIT. See [LICENSE](LICENSE)
