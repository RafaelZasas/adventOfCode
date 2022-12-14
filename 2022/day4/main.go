package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type elfPair struct {
	section1    []int
	section2    []int
	doesOverlap bool
    overlapsAny bool
}

var elves []elfPair

func (ep elfPair) String() string {
	return fmt.Sprintf("{\n\tsection1: %v,\n\tsection2: %v,\n\tdoesOverlap: %v\n},\n", ep.section1, ep.section2, ep.doesOverlap)
}

// init reads the file and builds the list of elf pairs
// with whether or not they overlap entirely or any at all
func init() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() { // scanner for the lines of the file
		line := scanner.Text()

		// scanner to split the line into each elves cleaning section
		pairsScanner := bufio.NewScanner(strings.NewReader(line))
		pairsScanner.Split(splitFunc(","))

		var pair elfPair

		for pairsScanner.Scan() { // scan through both pairs

			elf1Range := pairsScanner.Text() // the cleaning range for the first elf
			pairsScanner.Scan()
			elf2Range := pairsScanner.Text() // the cleaning range for the second elf

			// Ignoring errors from ascii to int convert

			elf1Start, _ := strconv.Atoi(elf1Range[:strings.IndexRune(elf1Range, '-')])
			elf1End, _ := strconv.Atoi(elf1Range[strings.IndexRune(elf1Range, '-')+1:])

			elf2Start, _ := strconv.Atoi(elf2Range[:strings.IndexRune(elf2Range, '-')])
			elf2End, _ := strconv.Atoi(elf2Range[strings.IndexRune(elf2Range, '-')+1:])

			pair.section1 = []int{elf1Start, elf1End}
			pair.section2 = []int{elf2Start, elf2End}
		}

		// checking to see if section 2 is entirely covered by section 1
		if pair.section2[0] >= pair.section1[0] && pair.section2[1] <= pair.section1[1] {
			pair.doesOverlap = true
		} else if pair.section1[0] >= pair.section2[0] && pair.section1[1] <= pair.section2[1] {
			// checking to see if section 1 is entirely covered by section 2
			pair.doesOverlap = true
		}

        // checking to see if an of the sections ovrelap at all
		if pair.section2[0] >= pair.section1[0] && pair.section2[0] <= pair.section1[1] {
			pair.overlapsAny = true
		} else if pair.section1[0] >= pair.section2[0] && pair.section1[0] <= pair.section2[1] {
			// checking to see if section 1 is entirely covered by section 2
			pair.overlapsAny = true
		}

		// if neither of the sections ovelaps the other the default will be false
		elves = append(elves, pair)
	}
}

// splitFunc implements a bufio.SplitFunc interface. It takes a
// string delimeter argument and returns a bufio.SplitFunc implementation
func splitFunc(delim string) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if i := strings.Index(string(data), fmt.Sprintf("%s", delim)); i >= 0 {
			return i + 1, data[0:i], nil
		}

		if atEOF {
			return len(data), data, nil
		}

		return
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	totalOverlap := 0
    overlapsAny := 0
    for _, elf := range elves {
		if elf.doesOverlap {
			totalOverlap += 1
		}
        if elf.overlapsAny {
            overlapsAny += 1
        }
	}
	fmt.Printf("Part One (total overlapping sections): %v\n", totalOverlap)
    fmt.Printf("Part Two (any overlapping sections): %v\n", overlapsAny)
}
