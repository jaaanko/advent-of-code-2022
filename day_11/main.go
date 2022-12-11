package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sort"
	
	"github.com/jaaanko/advent-of-code-2022/stack"
)

type monkey struct {
	items          stack.Stack[int]
	changeWorryLvl func(int) int
	mod            int
	targetIfPass   int
	targetIfFail   int
	itemsInspected int
}

func (m *monkey) addItem(item int) {
	m.items.Push(item)
}

func (m *monkey) popItem() (int,error) {
	return m.items.Pop()
}

func power(x int) int {
	return x * x
}

func multiply(y int) func(int) int {
	return func(x int) int {
		return x * y
	}
}

func add(y int) func(int) int {
	return func(x int) int {
		return x + y
	}
}

func main() {
	file1, err := os.Open("input_01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	file2, err := os.Open("input_01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	fmt.Println("Part 1:",solvePart1(parseInput(file1)))
	fmt.Println("Part 2:",solvePart2(parseInput(file2)))
}

func parseInput(r io.Reader) []monkey {
	scanner := bufio.NewScanner(r)
	monkeys := []monkey{}
	currMonkey := monkey{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			monkeys = append(monkeys, currMonkey)
			currMonkey = monkey{}
		}

		if strings.HasPrefix(line, "Starting items") {
			fields := strings.Fields(line)
			items := strings.Split(strings.Join(fields[2:], ""), ",")
			for _, worryLvl := range items {
				worryLvl, _ := strconv.Atoi(worryLvl)
				currMonkey.addItem(worryLvl)
			}
		} else if strings.HasPrefix(line, "Operation") {
			var op string
			var arg string
			fmt.Sscanf(line, "Operation: new = old %s %s", &op, &arg)

			if op == "*" {
				if arg == "old" {
					currMonkey.changeWorryLvl = power
				} else {
					intArg, err := strconv.Atoi(arg)
					if err != nil {
						log.Fatal(err)
					}
					currMonkey.changeWorryLvl = multiply(intArg)
				}
			} else {
				if arg == "old" {
					currMonkey.changeWorryLvl = multiply(2)
				} else {
					intArg, err := strconv.Atoi(arg)
					if err != nil {
						log.Fatal(err)
					}
					currMonkey.changeWorryLvl = add(intArg)
				}
			}
		} else if strings.HasPrefix(line, "Test") {
			var arg int
			fmt.Sscanf(line, "Test: divisible by %d", &arg)
			currMonkey.mod = arg
		} else if strings.HasPrefix(line, "If true") {
			var arg int
			fmt.Sscanf(line, "If true: throw to monkey %d", &arg)
			currMonkey.targetIfPass = arg
		} else if strings.HasPrefix(line, "If false") {
			var arg int
			fmt.Sscanf(line, "If false: throw to monkey %d", &arg)
			currMonkey.targetIfFail = arg
		}
	}

	return append(monkeys, currMonkey)
}

func solvePart1(monkeys []monkey) int {
	n := len(monkeys)
	itemsInspected := make([]int, n)

	for i:=0;i<20;i++ {
		for j:=0;j<n;j++ {
			for len(monkeys[j].items) > 0 {
				itemsInspected[j] += 1
				worryLvl,err := monkeys[j].popItem()

				if err != nil{
					log.Fatal(err)
				}
				worryLvl = monkeys[j].changeWorryLvl(worryLvl) / 3
				if worryLvl % monkeys[j].mod == 0 {
					monkeys[monkeys[j].targetIfPass].addItem(worryLvl)
				} else {
					monkeys[monkeys[j].targetIfFail].addItem(worryLvl)
				}
			}
		}
	}

	sort.Sort(sort.IntSlice(itemsInspected))
	return itemsInspected[n-1] * itemsInspected[n-2]
}

func solvePart2(monkeys []monkey) int {
	n := len(monkeys)
	mod := 1
	for i := range monkeys {
		mod *= monkeys[i].mod
	}
	itemsInspected := make([]int, n)

	for i:=0;i<10000;i++ {
		for j := range monkeys {
			for len(monkeys[j].items) > 0 {
				itemsInspected[j] += 1
				worryLvl,err := monkeys[j].popItem()

				if err != nil{
					log.Fatal(err)
				}
				worryLvl = monkeys[j].changeWorryLvl(worryLvl) % mod
				if worryLvl % monkeys[j].mod == 0 {
					monkeys[monkeys[j].targetIfPass].addItem(worryLvl)
				} else {
					monkeys[monkeys[j].targetIfFail].addItem(worryLvl)
				}
			}
		}
	}

	sort.Sort(sort.IntSlice(itemsInspected))
	return itemsInspected[n-1] * itemsInspected[n-2]
}
