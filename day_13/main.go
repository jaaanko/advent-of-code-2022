package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("input_01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	packets := parseInput(file)

	fmt.Println("Solve Part 1:", solvePart1(packets))
	fmt.Println("Solve Part 1:", solvePart2(packets))
}

func parseInput(r io.Reader) [][]any {
	scanner := bufio.NewScanner(r)
	packets := [][]any{}
	pair := []any{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			packets = append(packets, pair)
			pair = []any{}
		} else {
			var p any
			err := json.Unmarshal([]byte(scanner.Text()), &p)
			if err != nil {
				log.Fatal(err)
			}
			pair = append(pair, p)
		}
	}

	return append(packets, pair)
}

func solvePart1(packets [][]any) int {
	result := 0
	for i, pair := range packets {
		if compare(pair[0], pair[1]) <= 0 {
			result += i + 1
		}
	}

	return result
}

func compare(p1 any, p2 any) int {
	_, ok1 := p1.(float64)
	_, ok2 := p2.(float64)
	if ok1 && ok2 {
		return int(p1.(float64)) - int(p2.(float64))
	}

	if ok1 {
		p1 = []any{p1}
	}
	if ok2 {
		p2 = []any{p2}
	}

	if len(p1.([]any)) == 0 || len(p2.([]any)) == 0 {
		return len(p1.([]any)) - len(p2.([]any))
	}

	result := compare(p1.([]any)[0], p2.([]any)[0])
	if result == 0 {
		next1 := p1.([]any)[1:]
		next2 := p2.([]any)[1:]

		if len(next1) == 0 || len(next2) == 0 {
			return len(next1) - len(next2)
		}
		return compare(next1, next2)
	}

	return result
}

func solvePart2(packets [][]any) int {
	new := []any{}
	for _, pair := range packets {
		new = append(new, pair...)
	}

	var divider1 any
	err := json.Unmarshal([]byte("[[2]]"), &divider1)
	if err != nil {
		log.Fatal(err)
	}

	var divider2 any
	err = json.Unmarshal([]byte("[[6]]"), &divider2)
	if err != nil {
		log.Fatal(err)
	}

	new = append(new, []any{divider1, divider2}...)
	sort.Slice(new, func(i, j int) bool {
		return compare(new[i], new[j]) <= 0
	})

	result := 1
	for i, packet := range new {
		packet, err := json.Marshal(packet)
		if err != nil {
			log.Fatal(err)
		}
		packetString := string(packet)
		if packetString == "[[2]]" || packetString == "[[6]]" {
			result *= i + 1
		}
	}

	return result
}
