package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var data [][]int

func init() {
	start := time.Now()
	defer func() {
		fmt.Println("Init took: ", time.Since(start))
	}()

	file, err := os.Open("./input.txt")
	// file, err := os.Open("./example.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		report := []int{}
		for _, strNum := range strings.Split(line, " ") {

			intValue, err := strconv.Atoi(strNum)
			check(err)

			report = append(report, intValue)
		}
		data = append(data, report)
	}

}

func partOne() {
	safeCount := 0

dataLoop:
	for _, report := range data {

		if !isSorted(report) {
			continue dataLoop
		}

	reportLoop:
		for idx, num := range report {

			if idx == 0 {
				continue reportLoop
			}

			if diff(num, report[idx-1]) < 1 || diff(num, report[idx-1]) > 3 {
				continue dataLoop
			}

		}

		safeCount++

	}

	fmt.Printf("Safe Count: %d\n", safeCount)

}

func partTwo() {
	safeCount := 0

	for _, report := range data {
		if isSafe(report) {
			safeCount++
			continue
		}

		// Try removing each element to see if it becomes safe
		for i := 0; i < len(report); i++ {
			newReport := removeElement(report, i)
			if isSafe(newReport) {
				safeCount++
				break
			}
		}
	}

	fmt.Printf("Safe Count: %d\n", safeCount)
}

func main() {
	start := time.Now()
	partOne()
	fmt.Println("Part 1 took: ", time.Since(start))
	start = time.Now()
	partTwo()
	fmt.Println("Part 2 took: ", time.Since(start))
}

func isSafe(report []int) bool {
	if !isSorted(report) {
		return false
	}

	for i := 1; i < len(report); i++ {
		if diff(report[i], report[i-1]) < 1 || diff(report[i], report[i-1]) > 3 {
			return false
		}
	}

	return true
}

func isSorted(slice []int) bool {
	if len(slice) < 2 {
		return true // Empty or single-element slices are both ascending and descending
	}

	ascending := true
	descending := true

	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[i-1] {
			ascending = false
		} else if slice[i] > slice[i-1] {
			descending = false
		} else {
			// Equal elements are neither ascending nor descending
			ascending = false
			descending = false
		}
	}

	return ascending || descending
}

// need to add this since append modifies the slice in place
// which is dumb
func removeElement(slice []int, index int) []int {
	newSlice := make([]int, len(slice)-1)   // Create a new slice with the appropriate length
	copy(newSlice, slice[:index])           // Copy elements before the index
	copy(newSlice[index:], slice[index+1:]) // Copy elements after the index
	return newSlice
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
