package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func writeToFile(file string, code string) error {
	err := os.WriteFile(file, []byte(code), 0744)
	if err != nil {
		return errors.New("ERROR: Couldn't write to file.")
	}
	return nil
}

func writeGoStub(file string) error {
	code := "package main\n\nimport (\n\t// Your imports\n\t\"fmt\"\n)\n\nfunc main() {\n\t// Your logic\n\tfmt.Println(\">:3\")\n}"
	err := writeToFile(file, code)
	if err != nil {
		return errors.New("ERROR: Couldn't write Go stub.")
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
	fmt.Println(file, ext)

	switch ext {
	case ".go":
		if err := writeGoStub(file); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Cannot create stub for " + ext + " extension.")
	}
}
