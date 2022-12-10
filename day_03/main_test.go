package main

import (
	"strings"
	"testing"
)

var example = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

func TestPart1(t *testing.T) {
	expected := 157
	actual := solvePart1(parseInput(strings.NewReader(example)))
	if actual != expected {
		t.Errorf("Error in part 1, expected: %v, actual: %v", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 70
	actual := solvePart2(parseInput(strings.NewReader(example)))
	if actual != expected {
		t.Errorf("Error in part 2, expected: %v, actual: %v", expected, actual)
	}
}
