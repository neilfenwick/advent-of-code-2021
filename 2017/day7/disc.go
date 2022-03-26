package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Disc struct {
	Name     string
	Weight   int
	Children []string
}

func ParseDisc(str string) Disc {
	parts := strings.Split(str, " -> ")
	if len(parts) == 0 {
		panic(fmt.Sprintf("Invalid disc syntax: %s", str))
	}

	main := parts[0]
	weightIndex := strings.Index(main, " (")
	weightEnd := strings.LastIndex(str, ")")

	name := main[:weightIndex]
	weight, err := strconv.Atoi(main[weightIndex+2 : weightEnd])
	if err != nil {
		panic("Could not parse weight")
	}

	var children []string
	if len(parts) > 1 {
		children = strings.Split(parts[1], ", ")
	}

	return Disc{
		Name:     name,
		Weight:   weight,
		Children: children,
	}
}
