package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	num  int
	from int
	to   int
}

func (i instruction) String() string {
	return fmt.Sprintf("{\n\tnum: %d,\n\tfrom: %d,\n\tto: %d\n}\n", i.num, i.from, i.to)
}

var instructions []instruction

type crate string

type stack []crate

func (s stack) String() string {
	tmp := ""
	for _, c := range s {
		tmp += fmt.Sprintf("[%s] ", c)
	}
	tmp += "\n"
	return tmp
}

var crates []stack

// insert will add a crate to the bottom of the stack
// and is used when initializing the crates from the file
func (s *stack) insert(c crate) {
	*s = append([]crate{c}, *s...)
}

// push adds a crate to the top of the stack
func (s *stack) push(c crate) {
	*s = append(*s, c)
}

// pop removes a crate from the top of the stack and returns the removed crate
func (s *stack) pop() (crate, error) {
	if len(*s) > 0 {
		c := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return c, nil
	}
	fmt.Println(crates)
	return " ", fmt.Errorf("Attempting to pop from empty stack: %v", s)
}

// initCrates provides a clean setup for the crates and instructions
// which can be used to reset the configuration before part two
func initCrates() {
	crates = []stack{}
	instructions = []instruction{}
	file, err := os.ReadFile("./input.txt")
	check(err)

	input := strings.Split(string(file), "\n\n")
	cratesLines := strings.Split(input[0], "\n")
	cratesLines = cratesLines[:len(cratesLines)-1] // chop off the last line of numbers
	instructionsLines := strings.Split(input[1], "\n")

	// Build the crates grid
	for colIdx := 1; colIdx < len(cratesLines[0]); colIdx += 4 { // looping through the columns
		var s stack

		for _, line := range cratesLines {
			crateVal := line[colIdx]
			if string(crateVal) == " " {
				continue
			} else {
				s.insert(crate(crateVal))
			}
		}
		crates = append(crates, s)
	}

	// Build the instructions
	for _, i := range instructionsLines {

		if len(i) == 0 {
			break
		}

		moveIdx := strings.Index(i, "move")
		// First attempt to parse 2 digit move instruction
		num, err := strconv.Atoi(i[moveIdx+5 : moveIdx+7])
		// parse one digit move instruction if it fails
		if err != nil {
			num, _ = strconv.Atoi(string(i[moveIdx+5]))
		}

		from, _ := strconv.Atoi(string(i[strings.Index(i, "from")+5]))
		to, _ := strconv.Atoi(string(i[strings.Index(i, "to")+3]))

		instructions = append(instructions, instruction{
			from: from,
			to:   to,
			num:  num,
		})
	}
}

// partOne moves the crates one by one
func partOne() {
	for _, inst := range instructions {
		for i := 0; i < inst.num; i++ { // number of crates to move
			c, err := crates[inst.from-1].pop() // pop crate from required stack
			check(err)
			crates[inst.to-1].push(c)
		}
	}
}

// partTwo moves crates in chunks
func partTwo() {
	for _, inst := range instructions {
		var cratesToMove stack
		for i := 0; i < inst.num; i++ { // number of crates to move
			c, err := crates[inst.from-1].pop() // pop crate from required stack
			check(err)
			cratesToMove.insert(c)
		}

		// re stack the crates in the correct order
		for _, c := range cratesToMove {
			crates[inst.to-1].push(c)
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	initCrates()
	partOne()
	result := ""
	for _, row := range crates {
		result += string(row[len(row)-1])
	}
	fmt.Printf("Part One (Top Crates): %s\n", result)
	initCrates()
	partTwo()

	result = ""
	for _, row := range crates {
		result += string(row[len(row)-1])
	}
	fmt.Printf("Part Two (Top Crates 9001): %s\n", result)
}
