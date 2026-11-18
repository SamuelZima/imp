package main

import (
	"cmp"
	"errors"
	"fmt"
	"io/fs"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

const (
	// Compiled langs
	goCode     = "package main\n\nimport (\n\t// Your imports\n\t\"fmt\"\n)\n\nfunc main() {\n\t// Your logic\n\tfmt.Println(\">:3\")\n}\n"
	cCode      = "// Your headers\n#include <stdio.h>\n\nint main(int argc, char *argv[]) {\n\t// Your logic\n\tprintf(\">:3\\n\");\n}\n"
	cppCode    = "// Your headers\n#include <iostream>\n\nint main(int argc, char *argv[]) {\n\t// Your logic\n\tstd::cout << \">:3\" << std::endl;\n}\n"
	rustCode   = "fn main() {\n\t// Your logic\n\tprintln!(\">:3\");\n}\n"
	zigCode    = "// Your imports\nconst std = @import(\"std\");\n\npub fn main() void {\n\t// Your logic\n\tstd.debug.print(\">:3\\n\", .{});\n}\n"
	nimCode    = "proc main() =\n    # Your logic\n    echo \">:3\"\n\nwhen isMainModule:\n    main()\n"
	hsCode     = "main :: IO ()\nmain = putStrLn \">:3\"\n"
	kotlinCode = "// Your imports\n\nfun main() {\n\t// Your logic\n\tprintln(\">:3\")\n}\n"
	swiftCode  = "// Your imports\nimport Foundation\n\n// Your logic\nprint(\">:3\")\n"
	javaCode   = "void main() {\n\t// Your logic\n\tIO.println(\">:3\");\n}\n"
	csharpCode = "// Your logic\nConsole.WriteLine(\">:3\");\n"
	dartCode   = "void main() {\n\t// Your logic\n\tprint('>:3');\n}\n"

	// Interpreted langs
	pyCode   = "#!/usr/bin/env python3\n\ndef main():\n\t# Your logic\n\tprint(\">:3\")\n\nif __name__ == \"__main__\":\n\tmain()\n"
	jsTsCode = "#!/usr/bin/env node\n\n// Your logic\nconsole.log(\">:3\")\n"
	phpCode  = "#!/usr/bin/env php\n<?php\n// Your logic\necho \">:3\\n\";\n"
	rubyCode = "#!/usr/bin/env ruby\n\n# Your logic\nputs \">:3\"\n"
	perlCode = "#!/usr/bin/env perl\n\n# Your modules\nuse strict;\nuse warnings;\n\n# Your logic\nprint(\">:3\\n\")\n"
	rCode    = "#!/usr/bin/env Rscript\n\n# Your logic\ncat(\">:3\\n\")\n"
	luaCode  = "#!/usr/bin/env lua\n\n-- Your logic\nprint(\">:3\")\n"
	bashCode = "#!/usr/bin/env bash\n\n# Your logic\necho \">:3\"\n"
	zshCode  = "#!/usr/bin/env zsh\n\n# Your logic\necho \">:3\"\n"
	fishCode = "#!/usr/bin/env fish\n\n# Your logic\necho \">:3\"\n"
	ps1Code  = "#!/usr/bin/env pwsh\n\n# Your logic\nWrite-Output \">:3\"\n"

	// Other files
	mdCode           = "# Title\n\n>:3\n"
	htmlCode         = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n\t<meta charset=\"UTF-8\">\n\t<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n\t<title>Title</title>\n</head>\n<body>\n\t<!-- Your markup -->\n\t<p>>:3</p>\n</body>\n</html>\n"
	cssCode          = "*, *::before, *::after {\n\tbox-sizing: border-box;\n}\n\nbody {\n\tmargin: 0;\n}\n\n/* Your styles */\n"
	gitignoreCode    = "# OS\n.DS_Store\nThumbs.db\n\n# Editors\n.idea/\n.vscode/\n\n# Secrets\n.env\n\n# Logs\n*.log\n\n# Your ignores\n"
	editorconfigCode = "root = true\n\n[*]\ncharset = utf-8\nend_of_line = lf\ninsert_final_newline = true\ntrim_trailing_whitespace = true\n\n# Your rules\n"
)

var stubs = map[string]string{
	// Compiled langs
	".go":    goCode,
	".c":     cCode,
	".cpp":   cppCode,
	".rs":    rustCode,
	".zig":   zigCode,
	".nim":   nimCode,
	".hs":    hsCode,
	".kt":    kotlinCode,
	".swift": swiftCode,
	".java":  javaCode,
	".cs":    csharpCode,
	".dart":  dartCode,

	// Interpreted langs
	".py":   pyCode,
	".js":   jsTsCode,
	".ts":   jsTsCode,
	".php":  phpCode,
	".rb":   rubyCode,
	".pl":   perlCode,
	".lua":  luaCode,
	".r":    rCode,
	".sh":   bashCode,
	".zsh":  zshCode,
	".fish": fishCode,
	".ps1":  ps1Code,

	// Other files
	".md":           mdCode,
	".html":         htmlCode,
	".css":          cssCode,
	".gitignore":    gitignoreCode,
	".editorconfig": editorconfigCode,
}

var VERSION = "v0.7.0"

func usage() {
	fmt.Print("imp - a tiny familiar that conjures source file stubs\n\n")
	fmt.Print("Usage: imp [-h|--help] [-v|--version] [-l|--list] [-f|--force] <file_1> [file_2 ...]\n\n")
	fmt.Print("Options:\n")
	fmt.Print("\t-h, --help\tShow this help message\n")
	fmt.Print("\t-v, --version\tShow version\n")
	fmt.Print("\t-l, --list\tShow supported file extensions\n")
	fmt.Print("\t-f, --force\tOverwrite existing files\n")
}

func parseArgs() ([]string, bool) {
	var posArgs []string
	force := false

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-h", "--help":
			usage()
			os.Exit(0)
		case "-v", "--version":
			fmt.Printf("imp %s\n", VERSION)
			os.Exit(0)
		case "-l", "--list":
			fmt.Println(strings.Join(slices.Sorted(maps.Keys(stubs)), "\n"))
			os.Exit(0)
		case "-f", "--force":
			force = true
		default:
			if strings.HasPrefix(os.Args[i], "-") {
				fmt.Println("ERROR: Unknown flag", os.Args[i])
				os.Exit(2)
			}
			posArgs = append(posArgs, os.Args[i])
		}
	}

	if len(posArgs) == 0 {
		usage()
		os.Exit(2)
	}

	return posArgs, force
}

func writeToFile(file string, code string, force bool) error {
	// Set flags for file operations
	flags := os.O_WRONLY | os.O_CREATE | os.O_EXCL
	if force {
		// Only overwrite file when force flag is set
		flags = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	}

	// Set permissions, only stubs with a shebang are meant to be executed directly
	perm := fs.FileMode(0644)
	if strings.HasPrefix(code, "#!") {
		perm = 0755
	}

	// Create and open file, if it doesn't exist
	f, err := os.OpenFile(file, flags, perm)
	if errors.Is(err, fs.ErrExist) {
		return errors.New("ERROR: " + file + " already exists, use -f to overwrite.")
	}
	if err != nil {
		return fmt.Errorf("ERROR: Couldn't open file %s: %w", file, err)
	}
	defer f.Close()

	// Write code stub to file
	if _, err := f.WriteString(code); err != nil {
		return fmt.Errorf("ERROR: Couldn't write to %s: %w", file, err)
	}
	return f.Close()
}

func main() {
	files, force := parseArgs()

	for _, file := range files {
		// Get file extension
		ext := filepath.Ext(file)

		// Find code stub based on extension
		code, ok := stubs[ext]
		if !ok {
			fmt.Printf("Unsupported extension \"%s\" for file \"%s\"\n", cmp.Or(ext, "EMPTY"), file)
			continue
		}

		// Create or open file and write the code stub
		if err := writeToFile(file, code, force); err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Conjured", file)
	}
}
