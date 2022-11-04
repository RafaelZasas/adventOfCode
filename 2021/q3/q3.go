package q3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var report []string
var numCols int

// init reads the input file and formats it
// into a list of bit strings
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

// check panics if there is an error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// getGammaFromCol gets the resulting bit string
// from the majority of 0's or 1's from each column
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

// getEpsilonFromGamma flips all the bits of the gamma bit string
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

// getRating gets the remaining bit string's value
// which meets the criteria for either CO2 or O2
func getRating(l []string, index int, isO2 bool) int64 {

	// end of recursion, return the value of the last standing bit string
	if len(l) == 1 {
		res, err := strconv.ParseInt(l[0], 2, 32)
		check(err)
		return res
	}

	fmt.Printf("Col No: %d  |  len(l): %d\n", index, len(l))

	var lZeros []string
	var lOnes []string

	for _, bitStr := range l {
		if bitStr[index] == '1' {
			lOnes = append(lOnes, bitStr)
		} else {
			lZeros = append(lZeros, bitStr)
		}
	}

	if len(lZeros) == len(lOnes) {
		newL := []string{}
		fmt.Printf("Equal number of strings with zeros and ones\n BitStrings: %v\n", l)
		for i, subL := range l {
			bitString := subL[index]
			fmt.Printf("Sub List No: %d\nBit String (%[2]T): %[2]v\n", i, bitString)

			// only keep the bit strings with a 1 in the desired column if checking for O2
			if isO2 && bitString == '1' {
				fmt.Println("Was a 1")
				newL = append(newL, subL)
			}

			// only keep the bit strings with a 0 in the desired column if checking for CO2
			if !isO2 && bitString == '0' {
				fmt.Println("Was a 0")
				newL = append(newL, subL)
			}
		}

		fmt.Printf("Goin recursive again with:\n%v\n", newL)
		return getRating(newL, index+1, isO2)

	}

	if len(lZeros) > len(lOnes) {
		// Zeros are most common
		if isO2 {
			// if checking oxygen, choose most common (Zeros)
			return getRating(lZeros, index+1, isO2)
		} else {
			// if checking CO2, choose the one thats least common (Ones)
			return getRating(lOnes, index+1, isO2)
		}
	} else {
		if isO2 {
			// Checking O2, choose the most common (Ones)
			return getRating(lOnes, index+1, isO2)
		} else {
			// Checking CO2, choose the one thats least common (Zeros)
			return getRating(lZeros, index+1, isO2)

		}
	}
}

func part2() {
	fmt.Println("Getting O2")
	o2Rating := getRating(report, 0, true)
	fmt.Println("\nGetting CO2")
	cO2Rating := getRating(report, 0, false)
	fmt.Printf("O2 Rating: %v, CO2 Rating: %v, Result: %v\n", o2Rating, cO2Rating, o2Rating*cO2Rating)
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	part2()
}
