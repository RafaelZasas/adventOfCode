package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var fileName = "input.txt"

// var fileName = "example.txt"

var fileBuffer []byte

var solns []int
var inputs [][]int

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

	rows := strings.Split(string(fileBuffer), "\n")

	for _, row := range rows {
		if row == "" {
			continue
		}

		// split the row by :
		parts := strings.Split(row, ":")

		soln, err := strconv.Atoi(parts[0])
		check(err)
		solns = append(solns, soln)

		// split the second part by ' '
		var nums []int

		for _, num := range strings.Split(parts[1], " ") {
			if num == "" {
				continue
			}
			n, err := strconv.Atoi(num)
			check(err)
			nums = append(nums, n)
		}

		inputs = append(inputs, nums)

	}

	// for i, soln := range solns {
	// 	fmt.Printf("%d, %v\n", soln, inputs[i])
	// }

}

func partOne() {
	start := time.Now()
	defer func() {
		fmt.Println("Part one took: ", time.Since(start))
	}()

	ans := 0

solnLoop:
	for i, soln := range solns {

		nums := inputs[i]

		combinations := getCombinations(nums)

		for _, combination := range combinations {

			sum := 0

			for j, num := range nums {

				if j == 0 {
					sum += num
					continue
				}

				if combination[j-1] {
					sum += num
				} else {
					sum *= num
				}
			}

			if sum == soln {
				ans += soln
				continue solnLoop
			}

		}

	}

	fmt.Println("Part one: ", ans)
}

// well use true to represent + and false to represent *
func getCombinations(nums []int) [][]bool {

	numOperators := len(nums) - 1
	numCombinations := math.Pow(2, float64(numOperators))

	combinations := make([][]bool, int(numCombinations))

	for i := 0; i < int(numCombinations); i++ {
		combinations[i] = make([]bool, numOperators)

		for j := 0; j < numOperators; j++ {
			if i&(1<<uint(j)) != 0 {
				combinations[i][j] = true
			} else {
				combinations[i][j] = false
			}

		}

	}

	return combinations

}

func partTwo() {
	start := time.Now()
	defer func() {
		fmt.Println("Part two took: ", time.Since(start))
	}()

	ans := 0
solnLoop:
	for i, soln := range solns {

		nums := inputs[i]

		combinations := getCombinations2(nums)

		for _, combination := range combinations {

			sum := 0

			for j, num := range nums {

				if j == 0 {
					sum += num
					continue
				}

				if combination[j-1] == 0 {
					sum += num
				} else if combination[j-1] == 1 {
					sum *= num
				} else {
					// need to concatenate the two numbers
					// eg: 15 || 6 = 156
					strSum := strconv.Itoa(sum) + strconv.Itoa(num)
					tmp, err := strconv.Atoi(strSum)
					check(err)
					sum = tmp
				}
			}

			if sum == soln {
				ans += soln
				continue solnLoop
			}
		}
	}

	fmt.Println("Part two: ", ans)
}

func getCombinations2(nums []int) [][]uint8 {
	numOperators := len(nums) - 1
	numCombinations := int(math.Pow(3, float64(numOperators)))

	combinations := make([][]uint8, numCombinations)

	for i := 0; i < numCombinations; i++ {
		combinations[i] = make([]uint8, numOperators)
		value := i
		for j := 0; j < numOperators; j++ {
			combinations[i][j] = uint8(value % 3) // 0 = +, 1 = *, 2 = ||
			value /= 3
		}
	}

	return combinations
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
