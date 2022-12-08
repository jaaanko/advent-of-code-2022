package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/jaaanko/advent-of-code-2022/imath"
)

type file struct {
	name        string
	size        int
	parent      *file
	children    []*file
	isDirectory bool
}

func (f *file) addChild(fileToAdd *file) {
	f.children = append(f.children, fileToAdd)
}

func main() {
	f, err := os.Open("input_01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	currDirectory := &file{name: ""}
	root := currDirectory
	for scanner.Scan() {
		line := scanner.Text()

		if line[0] == '$' {
			command := strings.Split(line, " ")
			if command[1] == "cd" {
				arg := command[2]
				if arg == ".." {
					currDirectory = currDirectory.parent
				} else if arg != "/" {
					nextDirectory := &file{name: arg, parent: currDirectory, isDirectory: true}
					currDirectory.addChild(nextDirectory)
					currDirectory = nextDirectory
				}
			}
		} else {
			content := strings.Split(line, " ")
			if content[0] != "dir" {
				size, err := strconv.Atoi(content[0])
				if err != nil {
					log.Fatal(err)
				}
				currDirectory.addChild(&file{name: content[1], size: size, parent: currDirectory})
			}
		}
	}

	populateDirSizes(root)
	fmt.Println("Solve Part 1", solvePart1(root))
	fmt.Println("Solve Part 2", solvePart2(root))
}

func populateDirSizes(root *file) int {
	currTotal := 0
	for _, child := range root.children {
		if child.isDirectory {
			populateDirSizes(child)
		}
		currTotal += child.size
	}
	root.size = currTotal
	return currTotal
}

func solvePart1(root *file) int {
	total := 0
	if root.size <= 100000 {
		total += root.size
	}

	for _, child := range root.children {
		if child.isDirectory {
			total += solvePart1(child)
		}
	}

	return total
}

func solvePart2(root *file) int {
	var dfs func(*file)
	target := 30000000 - (70000000 - root.size)
	minSize := math.MaxInt64

	dfs = func(root *file) {
		if root.isDirectory && root.size >= target {
			minSize = imath.MinInt(minSize, root.size)
		}

		for _, child := range root.children {
			dfs(child)
		}
	}

	dfs(root)
	return minSize
}
