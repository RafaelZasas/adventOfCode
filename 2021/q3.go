package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var report []string
var numCols int

func init() {

	file, err := os.Open("./q3.txt")
	check(err)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		if numCols == 0 {
			numCols = len(scanner.Text())
		}
		report = append(report, scanner.Text())
	}
	file.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getGammaFromCol(report []string, colNum int) int {

	var countOnes int
	var countZeros int

	for _, line := range report {
		if line[colNum] == '1' {
			countOnes++
		} else {
			countZeros++
		}
	}
	if countOnes > countZeros {
		return 1
	} else {
		return 0
	}
}

func getEpsilonFromGamma(gamma string) string {
	e := ""
	for _, val := range gamma {
		if val == '0' {
			e += "1"
		} else {
			e += "0"
		}
	}
	return e
}

func part1() int64 {
	gamma := ""

	for i := 0; i < numCols; i++ {
		gamma += strconv.Itoa(getGammaFromCol(report, i))
	}

	epsilon := getEpsilonFromGamma(gamma)

	gammaVal, err := strconv.ParseInt(gamma, 2, 32)
	check(err)
	epsilonVal, err := strconv.ParseInt(epsilon, 2, 32)
	check(err)

	return gammaVal * epsilonVal
}

func getO2Rating(l []string, index int) string {
	if len(l) == 1 {
		return l[index]
	}

	var lZeros []string
	var lOnes []string

	for _, bitStr := range l {
		if bitStr[index] == '1' {
			lZeros = append(lZeros, bitStr)
		} else {
			lOnes = append(lOnes, bitStr)
		}
	}

	if len(lZeros) == len(lOnes) {
		for _, subL := range l {

			if subL[index] == '1' {
				return lZeros[index]
			} else {
				continue
			}
		}

	}

	if len(lZeros) > len(lOnes) {
		return getO2Rating(lZeros, index+1)
	} else {
		return getO2Rating(lOnes, index+1)
	}
}

func part2() {
	o2Rating := getO2Rating(report, 0)
	fmt.Println(o2Rating)
}

func main() {
	fmt.Printf("Part 1: %d", part1())
	part2()
}
