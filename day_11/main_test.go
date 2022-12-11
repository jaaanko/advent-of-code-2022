package main

import (
	"strings"
	"testing"
)

var example = `Monkey 0:
Starting items: 79, 98
Operation: new = old * 19
Test: divisible by 23
  If true: throw to monkey 2
  If false: throw to monkey 3

Monkey 1:
Starting items: 54, 65, 75, 74
Operation: new = old + 6
Test: divisible by 19
  If true: throw to monkey 2
  If false: throw to monkey 0

Monkey 2:
Starting items: 79, 60, 97
Operation: new = old * old
Test: divisible by 13
  If true: throw to monkey 1
  If false: throw to monkey 3

Monkey 3:
Starting items: 74
Operation: new = old + 3
Test: divisible by 17
  If true: throw to monkey 0
  If false: throw to monkey 1`

func TestParseInput(t *testing.T) {
	monkeys := parseInput(strings.NewReader(example))
	expectedLen := 4
	if len(monkeys) != expectedLen {
		t.Errorf("Expected length: %v, actual: %v", expectedLen, len(monkeys))
		t.FailNow()
	}

	testMonkey := monkey{
		items:          []int{79, 98},
		changeWorryLvl: func(x int) int { return x * 19 },
		mod:            23,
		targetIfPass:   2,
		targetIfFail:   3,
	}

	for i, worryLvl := range testMonkey.items {
		if worryLvl != monkeys[0].items[i] {
			t.Errorf("Expected: %v, actual: %v", testMonkey.items, monkeys[0].items)
			break
		}
	}

	if testMonkey.changeWorryLvl(2) != monkeys[0].changeWorryLvl(2) {
		t.Errorf("Expected: %v, actual: %v", testMonkey.changeWorryLvl(2), monkeys[0].changeWorryLvl(2))
	}

	if testMonkey.mod != monkeys[0].mod {
		t.Errorf("Expected: %v, actual: %v", testMonkey.mod, monkeys[0].mod)
	}

	if testMonkey.targetIfPass != monkeys[0].targetIfPass {
		t.Errorf("Expected: %v, actual: %v", testMonkey.targetIfPass, monkeys[0].targetIfPass)
	}

	if testMonkey.targetIfFail != monkeys[0].targetIfFail {
		t.Errorf("Expected: %v, actual: %v", testMonkey.targetIfFail, monkeys[0].targetIfFail)
	}
}

func TestPart1(t *testing.T) {
	expected := 10605
	actual := solvePart1(parseInput(strings.NewReader(example)))

	if actual != expected {
		t.Errorf("Error in part 1, expected: %v, actual: %v", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 2713310158
	actual := solvePart2(parseInput(strings.NewReader(example)))

	if actual != expected {
		t.Errorf("Error in part 2, expected: %v, actual: %v", expected, actual)
	}
}
