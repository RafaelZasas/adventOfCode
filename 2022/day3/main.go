package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type ruckSack struct {
	comp1    string
	comp2    string
	allItems string
}

var (
	ruckSacks   []ruckSack
	prioritySum = 0
)

func init() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		c1 := line[:len(line)/2]
		c2 := line[len(line)/2:]

		rs := ruckSack{
			comp1:    c1,
			comp2:    c2,
			allItems: line,
		}

		ruckSacks = append(ruckSacks, rs)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getPriority(s rune) int {
	if unicode.IsUpper(s) {
		return int(s - 'A' + 27) // 26 letters in the alphabet + 1 for indexing
	}

	return int(s - 'a' + 1)
}

// part one finds the first common item in the two compartments of each ruckSack
// and adds the priority level of that item type to the total
func partOne() {
	for i := 0; i < len(ruckSacks); i++ {
		for _, letter := range ruckSacks[i].comp1 {
			if strings.ContainsRune(ruckSacks[i].comp2, letter) {
				prioritySum += getPriority(letter)
				break
			}
		}
	}
}

func partTwo() {
    prioritySum = 0 // reset after part one
	for i := 0; i < len(ruckSacks); i += 3 {
		for _, letter := range ruckSacks[i].allItems {
            rs2Contains :=  strings.ContainsRune(ruckSacks[i+1].allItems, letter) 
            rs3Contains := strings.ContainsRune(ruckSacks[i+2].allItems, letter)
			if rs2Contains && rs3Contains  {
                prioritySum += getPriority(letter)
				break
			}
		}
	}
}

func main() {
	partOne()
	fmt.Printf("Part One: %v\n", prioritySum)

    partTwo()
    fmt.Printf("Part Two: %v\n", prioritySum)
}
