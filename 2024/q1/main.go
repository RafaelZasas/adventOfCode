package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var distanceTotal int
var leftList []int
var rightList []int

func init() {
	file, err := os.Open("./input.txt")
	// file, err := os.Open("./example.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		nums := strings.Split(line, "   ")

		leftNum, err := strconv.Atoi(nums[0])
		check(err)

		rightNum, err := strconv.Atoi(nums[1])
		check(err)

		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func partOne() {
	sort.IntSlice(leftList).Sort()
	sort.IntSlice(rightList).Sort()

	for i := 0; i < len(leftList); i++ {
		distanceTotal += diff(leftList[i], rightList[i])
	}

	fmt.Println("Part One: ", distanceTotal)

}

func partTwo() {
	similarityScore := 0

	countDict := make(map[int]int)

	for i := 0; i < len(rightList); i++ {
		countDict[rightList[i]]++
	}

	for i := 0; i < len(leftList); i++ {
		val, ok := countDict[leftList[i]]
		multiplier := 0
		if !ok {
			continue
		}

		multiplier = val

		similarityScore += leftList[i] * multiplier
	}

	fmt.Printf("Similarity Score %v\n", similarityScore)
}

func main() {
	startTime := time.Now()
	partOne()
	fmt.Printf("Go Part One took: %vµs\n", time.Since(startTime).Microseconds())

	startTime = time.Now()
	partTwo()
	fmt.Printf("Go Part Two took: %vµs\n", time.Since(startTime).Microseconds())
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
