package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	goCode = "package main\n\nimport (\n\t// Your imports\n\t\"fmt\"\n)\n\nfunc main() {\n\t// Your logic\n\tfmt.Println(\">:3\")\n}"
	cCode  = "// Your headers\n#include <stdio.h>\n\nint main(int argc, char *argv[]) {\n\t// Your logic\n\tprintf(\">:3\\n\");\n}"
)

func writeToFile(file string, code string) error {
	err := os.WriteFile(file, []byte(code), 0744)
	if err != nil {
		return errors.New("ERROR: Couldn't write to " + file)
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: Bad usage.")
		os.Exit(2)
	}

	file := os.Args[1]
	ext := filepath.Ext(file)
	//fmt.Println(file, ext) //DEBUG

	switch ext {
	case ".go":
		if err := writeToFile(file, goCode); err != nil {
			log.Fatal(err)
		}
	case ".c":
		if err := writeToFile(file, cCode); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Cannot create stub for " + ext + " extension.")
	}
}
