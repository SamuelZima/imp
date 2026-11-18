# Imp

A tiny familiar that conjures source file stubs.

`imp` creates new files pre-filled with a minimal, runnable starter for their language: a hello world for code, a skeleton for HTML/CSS, sensible defaults for config files. You can pass one or more file names. The stub is chosen by the file's extension. Scripts get a shebang and executable permissions, and existing files are left untouched unless you pass the `-f` or `--force` flag.

## Install

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
Conjured hello.py
>:3
```

## Supported extensions

- **Compiled:** `.go` `.c` `.cpp` `.rs` `.zig` `.nim` `.hs` `.kt` `.swift` `.java` `.cs` `.dart`
- **Interpreted:** `.py` `.js` `.ts` `.php` `.rb` `.pl` `.lua` `.r` `.sh` `.zsh` `.fish` `.ps1`
- **Other:** `.md` `.html` `.css` `.gitignore` `.editorconfig`

## Notes

- Existing files are never overwritten without `-f`.

## License

MIT. See [LICENSE](LICENSE)
