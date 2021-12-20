package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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

	diagnostics := initDiagnosticRegisters(file)
	power := diagnostics.PowerConsumption()
	fmt.Printf("Power consumption: %d\n", power.epsilonRate*power.gammaRate)

	lifeSupport := diagnostics.LifeSupport()
	fmt.Printf("Life support: %d\n", lifeSupport.oxygenGenerator*lifeSupport.co2Scrubber)
}

func initDiagnosticRegisters(r io.Reader) *DiagnosticRegisters {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	s.Scan()
	d := NewDiagnosticRegisters(len(s.Text()))
	d.AddReading(textToInt32(s.Text()))
	for s.Scan() {
		d.AddReading(textToInt32(s.Text()))
	}
	return d
}

func textToInt32(text string) int {
	i, err := strconv.ParseInt(text, 2, 32)
	if err != nil {
		log.Fatalf("Count not convert '%s' to int", text)
	}
	return int(i)
}
