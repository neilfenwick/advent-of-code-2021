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
	file, err = os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening file: %s", os.Args[1])
	}
	defer file.Close()

}
