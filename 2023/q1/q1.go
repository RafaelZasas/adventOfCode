package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var document []string

func init() {
	file, err := os.Open("./q1.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		document = append(document, line)
	}
}

// q2 finds the sum of calibration values, taking into account that
// each line may have numeric or alphabetic numbers such as "one" or "2"
// calibration values are found in the document lines as the first and last number in each line
func q2() {

	var wordNums = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	var sum int

	for _, line := range document {

		// all digits in the line
		var digits []string

		// loop through each char

		for charIdx, c := range line {

			// if char is a digit, add to digits
			if unicode.IsDigit(c) {
				//fmt.Printf("Found digit: %s\n", string(c))
				digits = append(digits, string(c))
			}

			// check if the the char we are on is the start of a word number
			// if it is, add the number to digits
			for idx, wordNum := range wordNums {
				if strings.HasPrefix(line[charIdx:], wordNum) {
					//fmt.Printf("Found word number: %s, %v\n", wordNum, rune(idx+1))
					digits = append(digits, fmt.Sprint(idx+1))
				}
			}

		}

		// concat the first and last digits

		firstDigit := string(digits[0])
		lastDigit := string(digits[len(digits)-1])

		concat := fmt.Sprintf("%s%s", firstDigit, lastDigit)

		//fmt.Printf("Concat: %s (%s, %s)\n", concat, firstDigit, lastDigit)

		num, err := strconv.Atoi(concat)
		check(err)

		// add to sum
		sum += num
	}

	fmt.Printf("Sum of calibration values (part 2): %d\n", sum)

}

// q1 finds the sum of all calibration values
// calibration values are found in the document lines
// and are the first number in each line concatenated with the last number in each line
func q1() {

	var sum int

	for _, line := range document {
		// Get first number in line
		firstDigit := getFirstNumber(line)

		// reverse the line and get the first number
		// this is a hacky way to get the second number

		// reverse the line
		reversed := []rune(line)
		for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
			reversed[i], reversed[j] = reversed[j], reversed[i]
		}

		// get the first number
		secondDigit := getFirstNumber(string(reversed))

		// concat two digits and convert to int
		num, err := strconv.Atoi(string(firstDigit) + string(secondDigit))

		if err != nil {
			fmt.Println(err)
		}

		// add to sum
		sum += num

	}

	fmt.Printf("Sum of calibration values: %d\n", sum)
}

func main() {
	q1()
	q2()
}

// -------------------------------------------------------
// Helper functions
// -------------------------------------------------------

// getFirstNumber returns the first number in a string
func getFirstNumber(s string) rune {
	// Loop through chars in line, stop at first occurence of a number
	for _, c := range s {
		if unicode.IsDigit(c) {
			return c
		}
	}
	return 0
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
