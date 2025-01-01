package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	tree "github.com/neilfenwick/advent-of-code/data_structures"
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
		_, _ = strconv.Atoi(os.Args[2])
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

	/*
	   Because this problem looked like a file tree calculation, I just used a tree data
	   structure from previous years.  In hindsight, that old tree wasn't the best
	   implementation. And it could probably do with being re-written with generics.

	   This solution could probably have been achieved by just passing a hashtable around
	   and doing a recursive parse of the input.

	   The solution feels a bit scatter-brained and unelegant 😕
	*/

	fileTree := processTerminalOutput(file)

	fmt.Printf("*************\nPart 1:\n*************\n")
	results := searchDirectoriesMaxSize(fileTree, 100000)

	totalSize := 0
	for _, size := range results {
		totalSize += size
	}
	fmt.Printf("Total size: %d\n", totalSize)

	fmt.Printf("*************\nPart 2:\n*************\n")
	results = searchDirectoriesMaxSize(fileTree, math.MaxInt)

	rootSize, _ := results["$root"]
	spaceRemaining := 70_000_000 - rootSize
	requiredToFree := 30_000_000 - spaceRemaining

	largeDirectories := make([]directory, 0)
	for name, size := range results {
		if size >= requiredToFree {
			largeDirectories = append(largeDirectories, directory{name: name, size: size})
			fmt.Printf("%s: %d\n", name, size)
		}
	}
	sort.Slice(largeDirectories, func(i, j int) bool {
		return largeDirectories[i].size < largeDirectories[j].size
	})
	fmt.Printf("Directory to delete: %v\n", largeDirectories[0])
}

type termLine struct {
	prefix string
	suffix string
}

type file struct {
	name string
	size int
}

type directory struct {
	name        string
	files       []file
	directories []directory
	size        int
}

func processTerminalOutput(file *os.File) *tree.Tree {
	s := bufio.NewScanner(file)
	var t *tree.Tree
	var currentNode *tree.TreeNode

	for s.Scan() {
		if s.Text() == "$ cd /" {
			log.Println("$root")
			t = tree.NewTree(tree.TreeKey{Name: "$root", Value: directory{name: "$root"}})
			_, currentNode = t.GetRoot()
			continue
		}

		if strings.HasPrefix(s.Text(), "$") {
			processCommand(s, t, currentNode)
		}
	}

	return t
}

func processCommand(s *bufio.Scanner, t *tree.Tree, currentNode *tree.TreeNode) {
	cmd := termLine{}
	fmt.Sscanf(s.Text(), "$ %s %s", &cmd.prefix, &cmd.suffix)

	switch cmd.prefix {
	case "cd":
		log.Printf("Changing directory to '%s'\n", cmd.suffix)

		if cmd.suffix == ".." {
			currentNode = currentNode.Parent
		} else {

			dirNode, found := currentNode.GetChild(currentNode.Key.Name + "/" + cmd.suffix)
			if !found {
				log.Printf("Could not find directory '%s' under '%s'", cmd.suffix, currentNode.Key.Name)
				return
			}

			currentNode = dirNode
		}
		log.Printf("Current directory: '%s'\n", currentNode.Key.Name)

		// The next line after a cd command must either be another cd or ls
		s.Scan()
		processCommand(s, t, currentNode)
	case "ls":
		for s.Scan() {
			if s.Err() == io.EOF {
				return
			}
			if strings.HasPrefix(s.Text(), "$") {
				processCommand(s, t, currentNode)
				return
			}

			line := termLine{}
			fmt.Sscanf(s.Text(), "%s %s", &line.prefix, &line.suffix)

			if strings.HasPrefix(line.prefix, "dir") {
				log.Printf("Appending directory '%s'\n", line.suffix)
				directory := directory{name: line.suffix}
				t.AppendChild(
					currentNode.Key,
					tree.TreeKey{Name: currentNode.Key.Name + "/" + line.suffix, Value: directory},
				)
				continue
			}

			log.Printf("Parsing file: %v\n", line)
			size, err := strconv.Atoi(line.prefix)
			if err != nil {
				log.Printf("Error parsing size %s under directory '%s'\n", line.prefix, currentNode.Key.Name)
			}

			f := file{name: line.suffix, size: size}
			t.AppendChild(currentNode.Key, tree.TreeKey{Name: currentNode.Key.Name + "/" + line.suffix, Value: f})
		}
	}
}

func searchDirectoriesMaxSize(fileTree *tree.Tree, maxSizeThreshold int) map[string]int {
	allDirectorySizes := make(map[string]int, 0)
	_, root := fileTree.GetRoot()
	searchDirectoriesRecursive(root, allDirectorySizes)

	results := make(map[string]int, 0)
	for key, value := range allDirectorySizes {
		if value <= maxSizeThreshold {
			results[key] = value
		}

	}

	return results
}

func searchDirectoriesRecursive(treeNode *tree.TreeNode, directorySizeMap map[string]int) {
	switch treeNode.Key.Value.(type) {
	case directory:
		for _, child := range treeNode.Children {
			searchDirectoriesRecursive(child, directorySizeMap)
			if child.Parent != nil {
				size, _ := directorySizeMap[child.Key.Name]
				directorySizeMap[child.Parent.Key.Name] += size
			}
		}
	case file:
		size := treeNode.Key.Value.(file).size
		directorySizeMap[treeNode.Parent.Key.Name] += size
	}
}
