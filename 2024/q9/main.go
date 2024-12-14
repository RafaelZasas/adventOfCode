package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

// var fileName = "input.txt"

var fileName = "example.txt"

var fileBuffer []byte

var data []int

type fileBlock struct {
	numChars int
	startIdx int
}

type emptySpace struct {
	numChars int
	startIdx int
}

var fileBlocks map[int]fileBlock
var emptySpaces []emptySpace

func init() {

	start := time.Now()
	defer func() {
		fmt.Println("Init took: ", time.Since(start))
	}()

	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	// read file into buffer and split by two new lines
	fileBuffer, err := io.ReadAll(file)
	check(err)

	blockIdx := 0
	diskIdx := 0 // to track where blocks of data or space start

	fileBlocks = make(map[int]fileBlock)
	emptySpaces = make([]emptySpace, 0)

	for i, b := range fileBuffer {
		if b == '\n' {
			continue
		}

		numChars, err := strconv.Atoi(string(b))
		check(err)

		// even means it is a block of data
		if i%2 == 0 {
			// fmt.Printf("adding %d of %d\n", numChars, byte(blockIdx))
			// add numCHars of blockIndex to data
			for j := 0; j < numChars; j++ {
				data = append(data, int(blockIdx))
			}

			// need to account for the fact that 1010 is 4 chars not 2
			// since block index can be greater than 9
			fileBlocks[blockIdx] = fileBlock{numChars, diskIdx}
			blockIdx++
		} else {
			// fmt.Printf("adding %d of .\n", numChars)
			// odd means its empty spaces
			// add numChars of . to data
			for j := 0; j < numChars; j++ {
				data = append(data, -1)
			}
			emptySpaces = append(emptySpaces, emptySpace{numChars, diskIdx})
		}

		diskIdx += numChars
	}

	// fmt.Printf("data: %v\n", data)

}

func partOne() {

	start := time.Now()
	defer func() {
		fmt.Println("Part one took: ", time.Since(start))
	}()

	// make a copy of data to not mutate the original
	newData := make([]int, len(data))
	copy(newData, data)

	frontIdx := 0
	backIdx := len(data) - 1

	for frontIdx < backIdx {

		// need to pop the first available number from the back and add it to the
		// front
		if newData[frontIdx] == -1 {

			if newData[backIdx] == -1 {
				backIdx--
				continue
			} else {
				// back index is a number
				newData[frontIdx] = newData[backIdx]
				newData[backIdx] = -1

				backIdx--
				frontIdx++
				continue
			}

		}

		frontIdx++

	}

	// fmt.Printf("newData: %v\n", newData)

	sum := 0

	for i, b := range newData {
		if b == -1 {
			break
		}

		// fmt.Printf("i: %d, b: %d\n", i, b)
		num := int(b)

		sum += num * (i)
	}

	fmt.Printf("sum: %d\n", sum)
}

// Need to move the entire blocks of data to the front
// where they fit, ie. free space > len block and it has to start with
// the block that has the hightst index
func partTwo() {
	start := time.Now()
	defer func() {
		fmt.Println("Part two took: ", time.Since(start))
	}()

	// make a copy of data to not mutate the original
	newData2 := make([]int, len(data))
	copy(newData2, data)

	for i := len(fileBlocks); i >= 0; i-- {
		b := fileBlocks[i]

		for j := 0; j < len(emptySpaces); j++ {
			e := emptySpaces[j]

			if e.startIdx < b.startIdx && e.numChars >= b.numChars {
				// move the block to the empty space
				// and update the data to have empty space where the block was
				for k := 0; k < b.numChars; k++ {
					newData2[e.startIdx+k] = i
					newData2[b.startIdx+k] = -1
				}

				// update the empty space count and index
				emptySpaces[j].numChars -= b.numChars
				emptySpaces[j].startIdx += b.numChars

				// Update the vacated space as a new empty space
				emptySpaces = append(emptySpaces, emptySpace{b.numChars, b.startIdx})

				break
			}

		}

	}

	// printData(newData2)
	sum := 0

	for i, b := range newData2 {
		if b == -1 {
			continue
		}

		// fmt.Printf("i: %d, b: %d\n", i, b)
		num := int(b)

		sum += num * (i)
	}

	fmt.Printf("sum: %d\n", sum)

}

func printData(data []int) {
	str := ""
	for _, b := range data {
		if b == -1 {
			str += "."
		} else {
			str += strconv.Itoa(b)
		}
	}

	fmt.Println(str)
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
