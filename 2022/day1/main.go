// question one of the 2022 advent of code challenge
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)


type elf struct {
	food []int
	cal  int
}

func (e elf) String () string {

	return fmt.Sprintf("{food: %v, cal: %v}\n", e.food, e.cal)
}

var elves []elf

// read and process the text file
func init() {

	file, err := os.Open("./q1.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	index  := 0
	for scanner.Scan() {
		
		if scanner.Text() != "" {
			
			if len(elves) <= index{

				elves = append(elves, elf{food: []int{}, cal: 0})
			}
			
			cals, err := strconv.Atoi(scanner.Text())
			check(err)
			elves[index].cal += cals
			elves[index].food = append (elves[index].food, cals)
		} else {

			index+=1
		}
	}
}

// check panics if there is an error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].cal < elves[j].cal
	})

	// part one: max cals
	fmt.Printf("Elf with most cals: %v\n", elves[len(elves)-1])

	// part two: sum of top 3 cals
  sum := 0
	for i := 1; i <= 3; i ++ {
		sum += elves[len(elves)-i].cal
	}

	fmt.Printf("Sum of top 3 elves cals: %v", sum)
}
