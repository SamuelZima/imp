package main

import (
	"cmp"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
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
	vCode      = "fn main() {\n\t// Your logic\n\tprintln('>:3')\n}\n"

	// Interpreted langs
	pyCode   = "#!/usr/bin/env python3\n\ndef main():\n\t# Your logic\n\tprint(\">:3\")\n\nif __name__ == \"__main__\":\n\tmain()\n"
	jsTsCode = "#!/usr/bin/env node\n\n// Your logic\nconsole.log(\">:3\")\n"
	phpCode  = "#!/usr/bin/env php\n<?php\n// Your logic\necho \">:3\\n\";\n"
	rubyCode = "#!/usr/bin/env ruby\n\n# Your logic\nputs \">:3\"\n"
	perlCode = "#!/usr/bin/env perl\n\n# Your modules\nuse strict;\nuse warnings;\n\n# Your logic\nprint(\">:3\\n\")\n"
	luaCode  = "#!/usr/bin/env lua\n\n-- Your logic\nprint(\">:3\")\n"
	bashCode = "#!/usr/bin/env bash\n\n# Your logic\necho \">:3\"\n"
)

var stubs = map[string]string{
	// Compiled langs
	".go":    goCode,
	".c":     cCode,
	".cpp":   cppCode,
	".cc":    cppCode,
	".cxx":   cppCode,
	".rs":    rustCode,
	".zig":   zigCode,
	".nim":   nimCode,
	".hs":    hsCode,
	".kt":    kotlinCode,
	".swift": swiftCode,
	".v":     vCode,

	// Interpreted langs
	".py":  pyCode,
	".js":  jsTsCode,
	".ts":  jsTsCode,
	".php": phpCode,
	".rb":  rubyCode,
	".pl":  perlCode,
	".lua": luaCode,
	".sh":  bashCode,
}

var VERSION = "v0.4.0"

func usage() {
	fmt.Print("imp - a tiny helper that conjures source file stubs >:3\n\n")
	fmt.Print("Usage: imp [-h|--help] [-v|--version] [-f|--force] <file_1> [file_2 ...]\n\n")
	fmt.Print("Options:\n")
	fmt.Print("\t-h, --help\tShow this help message\n")
	fmt.Print("\t-v, --version\tShow version\n")
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
		case "-f", "--force":
			force = true
		default:
			if strings.HasPrefix(os.Args[i], "-") {
				fmt.Println("ERROR: Unknown flag " + os.Args[i])
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
		fmt.Println(file)
		ext := filepath.Ext(file)

		code, ok := stubs[ext]
		if !ok {
			fmt.Println("ERROR: Cannot create stub for " + cmp.Or(ext, "EMPTY") + " extension.")
			os.Exit(1)
		}

		if err := writeToFile(file, code, force); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
