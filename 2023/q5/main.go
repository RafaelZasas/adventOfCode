package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

var document []string
var seeds []int

type AlmanacMap struct {
	source      int
	destination int
	distance    int
}

// to string for almanac map
func (am AlmanacMap) String() string {
	return fmt.Sprintf("source: %d, destination: %d, distance: %d", am.source, am.destination, am.distance)
}

var fullMap [7][]AlmanacMap

func init() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		document = append(document, line)
	}

	line := strings.Split(document[0], ":")[1]

	line = strings.TrimSpace(line)

	// split numbers by space and convert to int

	tmpSeeds := strings.Split(line, " ")

	for _, seed := range tmpSeeds {
		seed, err := strconv.Atoi(seed)
		check(err)
		seeds = append(seeds, seed)
	}

	mapIndex := 0

	// looping through document after seeds to get almanac maps
	// starting at line 3 since this is the start of the seed-to-soil map
	for i := 3; i < len(document); i++ {
		line := document[i]

		if len(line) == 0 {
			mapIndex++ // At this point weve reached an empty line, which means we are at the end of the current map
			continue
		}

		// this means were on the descrtiption line
		if strings.Contains(line, "-") {
			continue
		}

		// split line by space
		parts := strings.Split(line, " ")

		// convert parts to int
		destination, err := strconv.Atoi(parts[0])
		check(err)

		source, err := strconv.Atoi(parts[1])
		check(err)

		distance, err := strconv.Atoi(parts[2])
		check(err)

		// create new almanac map
		am := AlmanacMap{source: source, destination: destination, distance: distance}

		// append to full map
		fullMap[mapIndex] = append(fullMap[mapIndex], am)

	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1(seeds []int) []int {
	defer timeTrack(time.Now(), "part1")

	locations := make([]int, len(seeds))

	for i, seed := range seeds {
		// this will be initialized to seed value, then soil, then fetrtilizer etc.
		// until we reach the end of the maps and get the location
		location := seed

		// loop through full map ie. the almanac maps array for each source to
		// destination map
		for _, mapArray := range fullMap {
			// fmt.Printf("destination-source map no: %d, value: %d\n", j, location)

			// loop through each almanac map in the source-desination map array
			for _, am := range mapArray {

				// fmt.Printf("almanac map number: %d\n", k)

				// determine is the initial value is in the range of the source and
				// destination

				if location >= am.source && location <= am.source+am.distance-1 {
					// if it is, then we set the initial value to the distance
					// between the source and destination
					// fmt.Println("Source does fit in range")
					// fmt.Println(am)
					location = location - am.source + am.destination

					break
				}

			}
		}

		locations[i] = location
	}

	return locations

}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func getNewSeeds(seeds []int) []int {

	defer timeTrack(time.Now(), "getNewSeeds")

	newSeeds := []int{}

	startRanges := []int{}
	endRanges := []int{}

	for i := 0; i < len(seeds); i++ {
		if i%2 == 0 {
			startRanges = append(startRanges, seeds[i])
		}
		if i%2 != 0 {
			endRanges = append(endRanges, seeds[i])
		}

	}

	fmt.Printf("Length of startRanges: %d\n", len(startRanges))

	for i := 0; i < len(startRanges); i++ {
		start := startRanges[i]
		end := endRanges[i]

		fmt.Printf("Checking range number: %d\n", i)

		// now we need to loop through the range and add each number to the
		// newSeeds array d to avoid adding it again
		for j := start; j <= start+end-1; j++ {
			newSeeds = append(newSeeds, j)
		}
	}

	fmt.Println("Sorting new seeds")
	slices.Sort(newSeeds)
	fmt.Println("Removing duplicates")
	newSeeds = slices.Compact(newSeeds)

	return newSeeds
}

// for this one the seeds are actually ranges, where the first number is the
// start of the range and the second number is the end of the range, and the
// pattern continues for each even and odd index
func part2() {

	newSeeds := getNewSeeds(seeds)

	defer timeTrack(time.Now(), "part2")

	fmt.Printf("Calculating locations for %d seeds\n\n", len(newSeeds))
	locations := part1(newSeeds)

	fmt.Println("Sorting locations")
	slices.Sort(locations)

	fmt.Printf("Smallest location: %d\n", locations[0])

}

func main() {
	locations := part1(seeds)

	// find the smallest value in the locations array

	slices.Sort(locations)

	fmt.Println(locations[0])
	part2()
}
