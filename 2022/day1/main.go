package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	var (
		file *os.File
		err  error
	)

	switch len(os.Args) {
	case 1:
		file = os.Stdin
	case 3:
		fallthrough
	case 2:
		file, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatalf("Error opening file: %s", os.Args[1])
		}
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	printMaxCalories(buildFoodPacks(file))
}

type foodPack struct {
	calories int
}

func (f *foodPack) addFood(calories int) {
	f.calories += calories
}

func buildFoodPacks(reader io.Reader) []foodPack {
	pack := foodPack{}
	packs := []foodPack{pack}
	s := bufio.NewScanner(reader)
	for s.Scan() {
		if s.Err() == io.EOF {
			packs = append(packs, pack)
			pack = foodPack{}
			break
		}
		if len(s.Text()) == 0 {
			packs = append(packs, pack)
			pack = foodPack{}
			continue
		}
		calories, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Printf("Error parsing %s. %s", s.Text(), err)
		}
		pack.addFood(calories)
	}
	return packs
}

func printMaxCalories(foodPacks []foodPack) {
	sort.Sort(byCalories(foodPacks))
	fmt.Println(foodPacks)
}

type byCalories []foodPack

func (c byCalories) Len() int           { return len(c) }
func (c byCalories) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c byCalories) Less(i, j int) bool { return c[i].calories < c[j].calories }
