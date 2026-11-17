package main

import (
	"cmp"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	javaCode   = "// Your imports\n\npublic class Main {\n\tpublic static void main(String[] args) {\n\t\t// Your logic\n\t\tSystem.out.println(\">:3\");\n\t}\n}\n"
	kotlinCode = "// Your imports\n\nfun main() {\n\t// Your logic\n\tprintln(\">:3\")\n}\n"
	swiftCode  = "// Your imports\nimport Foundation\n\n// Your logic\nprint(\">:3\")\n"
	vCode      = "fn main() {\n\t// Your logic\n\tprintln('>:3')\n}\n"

	// Interpreted langs
	pyCode   = "#!/usr/bin/env python3\n\ndef main():\n\t# Your logic\n\tprint(\">:3\")\n\nif __name__ == \"__main__\":\n\tmain()\n"
	jsTsCode = "// Your logic\nconsole.log(\">:3\")\n"
	phpCode  = "#!/usr/bin/env php\n<?php\n// Your logic\necho \">:3\\n\";\n"
	rubyCode = "#!/usr/bin/env ruby\n\n# Your logic\nputs \">:3\"\n"
	perlCode = "#!/usr/bin/env perl\n\n# Your modules\nuse strict;\nuse warnings;\n\n# Your logic\nprint(\">:3\\n\")\n"
	luaCode  = "#!/usr/bin/env lua\n\n-- Your logic\nprint(\">:3\")\n"
	bashCode = "#!/usr/bin/env bash\n\n# Your logic\necho \">:3\"\n"
	shCode   = "#!/bin/sh\n\n# Your logic\necho \">:3\"\n"
)

var VERSION = "v0.1.0"

func usage() {
	fmt.Println("imp")
}

func parseArgs() string {
	var posArgs []string

	if len(os.Args) == 1 {
		usage()
		os.Exit(2)
	}

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-h", "--help":
			usage()
			os.Exit(0)
		case "-v", "--version":
			fmt.Printf("imp %s\n", VERSION)
			os.Exit(0)
		default:
			posArgs = append(posArgs, os.Args[i])
		}
	}

	return posArgs[0]
}

func writeToFile(file string, code string) error {
	err := os.WriteFile(file, []byte(code), 0744)
	if err != nil {
		return errors.New("ERROR: Couldn't write to " + file)
	}
	return nil
}

func main() {
	file := parseArgs()
	ext := filepath.Ext(file)

	switch ext {
	case ".go":
		if err := writeToFile(file, goCode); err != nil {
			log.Fatal(err)
		}
	case ".c":
		if err := writeToFile(file, cCode); err != nil {
			log.Fatal(err)
		}
	case ".cpp", ".cc", ".cxx":
		if err := writeToFile(file, cppCode); err != nil {
			log.Fatal(err)
		}
	case ".rs":
		if err := writeToFile(file, rustCode); err != nil {
			log.Fatal(err)
		}
	case ".zig":
		if err := writeToFile(file, zigCode); err != nil {
			log.Fatal(err)
		}
	case ".nim":
		if err := writeToFile(file, nimCode); err != nil {
			log.Fatal(err)
		}
	case ".hs":
		if err := writeToFile(file, hsCode); err != nil {
			log.Fatal(err)
		}
	case ".java":
		if err := writeToFile(file, javaCode); err != nil {
			log.Fatal(err)
		}
	case ".kt":
		if err := writeToFile(file, kotlinCode); err != nil {
			log.Fatal(err)
		}

	case ".swift":
		if err := writeToFile(file, swiftCode); err != nil {
			log.Fatal(err)
		}
	case ".v":
		if err := writeToFile(file, vCode); err != nil {
			log.Fatal(err)
		}
	case ".py":
		if err := writeToFile(file, pyCode); err != nil {
			log.Fatal(err)
		}
	case ".js", ".ts":
		if err := writeToFile(file, jsTsCode); err != nil {
			log.Fatal(err)
		}
	case ".php":
		if err := writeToFile(file, phpCode); err != nil {
			log.Fatal(err)
		}
	case ".rb":
		if err := writeToFile(file, rubyCode); err != nil {
			log.Fatal(err)
		}
	case ".pl":
		if err := writeToFile(file, perlCode); err != nil {
			log.Fatal(err)
		}
	case ".lua":
		if err := writeToFile(file, luaCode); err != nil {
			log.Fatal(err)
		}
	case ".sh":
		if err := writeToFile(file, bashCode); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("ERROR: Cannot create stub for " + cmp.Or(ext, "EMPTY") + " extension.")
	}
}
