package main

import (
	"log"
	"os"
)

func main() {
	var (
		file *os.File
		err  error
	)

	switch len(os.Args) {
	case 1:
		file = os.Stdin
	case 2:
		file, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatalf("Error opening file: %s", os.Args[1])
		}
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
}
