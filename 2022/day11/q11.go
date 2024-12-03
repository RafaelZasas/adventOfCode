// monkey business problem from day 11 of the 2022 advent of code callenge
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

const numRounds = 20

type monkey struct {
	items          []int
	operationType  string // the operand +-*/ for which the monkey operates worry level
	operationNum   int
	testType       string
	testNumber     int
	truthyThrowIdx int // the monkey number if test is true
	falsyThrowIdx  int
	itemsSeen      int
}

var monkeys []monkey

func (m monkey) String() string {
	return fmt.Sprintf("{\nitems: %v\n, itemsSeen: %v\n}\n", m.items, m.itemsSeen)
}

// init reads the input file and adds each monkey isntance to the monkeys slcie
func init() {
	file, err := os.Open("./input.txt")

	check(err)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// remove will delete an element from a slice
// and return the resulting slice
func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func computeOperation(operand string, operationNum, oldNum int) int {
	switch operand {
	case "+":
		return oldNum + operationNum
	case "*":
		return oldNum * operationNum
	default:
		return 0
	}
}

func testMonkeyThrow(worryLvl, testNum int) bool {
	if worryLvl%testNum == 0 {
		return true
	} else {
		return false
	}
}

// computeMonkeyThrow will determine to which monkey the item must be thrown
func computeMonkeyThrow(m *monkey) {
	for i := 0; i < len(m.items); i++ {
		worryLvl := m.items[i]
		worryLvl = computeOperation(m.operationType, m.operationNum, worryLvl)
		worryLvl = int(math.Round(float64(worryLvl) / 3))

		// throw the item to the corresponding monkey based on the test
		if testMonkeyThrow(worryLvl, m.testNumber) {
			monkeys[m.truthyThrowIdx].items = append(monkeys[m.truthyThrowIdx].items, worryLvl)
			monkeys[m.truthyThrowIdx].itemsSeen += 1
		} else {
			monkeys[m.falsyThrowIdx].items = append(monkeys[m.falsyThrowIdx].items, worryLvl)
			monkeys[m.falsyThrowIdx].itemsSeen += 1
		}

		m.items = remove(m.items, i) // remove the item from the current monkey since its been thrown
	}
}

func computeMonkeyBusiness() {
	for i := 0; i < numRounds; i++ {
		for j := 0; j < len(monkeys); i++ {
		}
	}
}

func main() {
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].itemsSeen < monkeys[j].itemsSeen
	})
}
