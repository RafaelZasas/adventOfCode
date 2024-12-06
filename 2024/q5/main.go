package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

var fileName = "input.txt"

// var fileName = "example.txt"

var fileBuffer []byte
var updates []string

// this map will store the orders in the format of y -> [x1, x2, x3]
// where all the x's are the numbers that appear before y
var collectionDict = make(map[int][]int)

var incorrectUpdates [][]int

func init() {
	start := time.Now()
	defer func() {
		fmt.Println("Init took: ", time.Since(start))
	}()

	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	// read file into buffer and split by two new lines
	fileBuffer, err := io.ReadAll(file)
	check(err)

	parts := strings.Split(string(fileBuffer), "\n\n")

	ordersLines := strings.Split(parts[0], "\n")
	for _, orderLine := range ordersLines {
		nums := strings.Split(orderLine, "|")
		num1, err := strconv.Atoi(nums[0])
		check(err)
		num2, err := strconv.Atoi(nums[1])
		check(err)
		collectionDict[num2] = append(collectionDict[num2], num1)
	}

	updates = strings.Split(parts[1], "\n")

}

func partOne() {
	start := time.Now()
	defer func() {
		fmt.Println("Part one took: ", time.Since(start))
	}()

	sum := 0

	// for v, key := range collectionDict {
	// 	fmt.Println(v, key)
	// }

	for _, update := range updates {

		values := strings.Split(update, ",")

		intValues := make([]int, len(values))
		for i, v := range values {
			intValues[i], _ = strconv.Atoi(v)
		}

		ok := true

		for i, v := range intValues {
			for j, k := range intValues {
				if i < j && slices.Contains(collectionDict[v], k) {
					ok = false
					break
				}
			}
		}

		if ok {
			middle := len(intValues) / 2
			sum += intValues[middle]
		} else {
			incorrectUpdates = append(incorrectUpdates, intValues)
		}
	}

	fmt.Println("Sum: ", sum)

}

// Need to fix the incorrect updates
// using the rules and then sum the middle values
// eg [97,13,75,29,47] -> [97,75,47,29,13]
func partTwo() {
	start := time.Now()
	defer func() {
		fmt.Println("Part two took: ", time.Since(start))
	}()

	sum := 0

	for _, update := range incorrectUpdates {

		sort.Slice(update, func(i, j int) bool {
			left := update[i]
			right := update[j]
			rule, found := collectionDict[left]

			if found && slices.Contains(rule, right) {
				return true
			}

			return false
		})

		middle := len(update) / 2
		sum += update[middle]
	}

	// for _, v := range incorrectUpdates {
	// 	fmt.Println(v)
	// }

	fmt.Println("Sum: ", sum)

}

func main() {
	partOne()
	partTwo()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
