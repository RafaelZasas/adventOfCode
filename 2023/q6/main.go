package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Race struct {
	time        int
	distance    int
	winningWays []int
}

var races []Race
var p2Race Race

func init() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := strings.Split(scanner.Text(), ":")[1]
	times := extractNumbers(line)
	p2Time := extractNumbersp2(line)

	scanner.Scan()

	line = strings.Split(scanner.Text(), ":")[1]
	distances := extractNumbers(line)
	p2Distance := extractNumbersp2(line)

	p2Race = Race{time: p2Time, distance: p2Distance}
	for i, t := range times {
		r :=
			Race{
				time:     t,
				distance: distances[i],
			}
		races = append(races, r)

	}
}

// Calculate the winning distances
// The boat increases by 1mm/ms for each ms the button is held
func q1() {

	product := 1
	for _, race := range races {

		// loop through n times where n = race time
		var distances []int
		for i := 0; i <= race.time; i++ {
			distances = append(distances, i*(race.time-i))
		}

		var winningDistances []int

		for _, d := range distances {
			if d > race.distance {
				winningDistances = append(winningDistances, d)
			}
		}

		race.winningWays = winningDistances

		product *= len(winningDistances)
	}

	fmt.Printf("Q1) Product of winning ways: %d\n", product)

}

func q2() {
	var distances []int
	for i := 0; i < p2Race.time; i++ {
		distances = append(distances, i*(p2Race.time-i))
	}

	var winningWays []int

	for _, d := range distances {
		if d > p2Race.distance {
			winningWays = append(winningWays, d)
		}
	}

	fmt.Printf("Q2) Num winning ways: %d\n", len(winningWays))
}

func main() {
	q1()
	q2()
}

// Helper functions

func extractNumbers(str string) (nums []int) {

	re, err := regexp.Compile("[0-9]+")

	if err != nil {
		panic(err)
	}

	// Find all the numbers
	numbers := re.FindAllString(str, -1)

	for _, number := range numbers {
		n, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}
	return
}

func extractNumbersp2(str string) (n int) {

	re, err := regexp.Compile("[0-9]+")

	if err != nil {
		panic(err)
	}

	// Find all the numbers
	numbers := re.FindAllString(str, -1)

	num := strings.Join(numbers, "")

	n, err = strconv.Atoi(num)

	if err != nil {
		panic(err)
	}

	return
}
