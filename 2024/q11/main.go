package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var fileName = "example.txt"

var fileBuffer []byte

var initialStones []int64

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

	strStones := strings.Fields(string(fileBuffer))

	for _, s := range strStones {
		i, err := strconv.ParseInt(s, 10, 64)
		check(err)
		initialStones = append(initialStones, int64(i))
	}

}

func partOne() {
	start := time.Now()
	defer func() {
		fmt.Println("Part one took: ", time.Since(start))
	}()

	newStones := make([]int64, len(initialStones))
	copy(newStones, initialStones)

	for i := 0; i < 25; i++ {
		computedStones := blinkStones(&newStones)
		newStones = computedStones
		// fmt.Printf("Blink %d: %v\n", i+1, len(newStones))

	}

	fmt.Printf("len stones: %v\n", len(newStones))
}

// Rules:
// If stone is 0 it becomes 1
// If stone has even num of digits:
// - It gets split into two stones
// - Left stone gets left half
// - Right stone gets right half, except:
//   - Right half needs to be void of leding and trailing zeros
//
// # In all other cases, stone is multiplied by 2024
//
// Stones should be kept in order
func blinkStones(stones *[]int64) (computedStoes []int64) {
	for i := 0; i < len(*stones); i++ {

		stone := (*stones)[i]

		if stone == 0 {
			computedStoes = append(computedStoes, 1)
			continue
		}

		numDigits := int(math.Log10(float64(stone))) + 1

		if numDigits%2 == 0 {
			strStone := fmt.Sprintf("%d", stone)

			left := strStone[:numDigits/2]
			right := strings.TrimLeft(strStone[numDigits/2:], "0")

			if right == "" {
				right = "0"
			}

			intLeft, err := strconv.ParseInt(left, 10, 64)
			check(err)

			intRight, err := strconv.ParseInt(right, 10, 64)
			check(err)

			computedStoes = append(computedStoes, int64(intLeft), int64(intRight))

			continue
		}

		// multiply by 2024
		computedStoes = append(computedStoes, stone*2024)

	}

	return computedStoes
}

func partTwo() {
	start := time.Now()
	defer func() {
		fmt.Println("Part two took: ", time.Since(start))
	}()

	newStones := make([]int64, len(initialStones))
	copy(newStones, initialStones)

	for i := 0; i < 75; i++ {
		fmt.Printf("Blink %d: %v\n", i+1, len(newStones))
		computedStones := optimizedBlink(newStones)
		newStones = computedStones
	}

	fmt.Printf("len stones: %v\n", len(newStones))
}

func optimizedBlink(stones []int64) []int64 {
	var wg sync.WaitGroup
	resultChan := make(chan []int64, len(stones)/100+1) // Use buffered channel to avoid blocking

	// Split stones into chunks for parallel processing
	chunkSize := len(stones) / 8 // Divide work into 8 chunks
	if chunkSize == 0 {
		chunkSize = len(stones)
	}

	for start := 0; start < len(stones); start += chunkSize {
		end := start + chunkSize
		if end > len(stones) {
			end = len(stones)
		}

		wg.Add(1)
		go func(chunk []int64) {
			defer wg.Done()
			var localResult []int64
			for _, stone := range chunk {
				localResult = append(localResult, processStone(stone)...)
			}
			resultChan <- localResult
		}(stones[start:end])
	}

	// Collect results
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var result []int64
	for chunk := range resultChan {
		result = append(result, chunk...)
	}

	return result
}

func processStone(stone int64) []int64 {
	if stone == 0 {
		return []int64{1}
	}

	strStone := strconv.FormatInt(stone, 10)
	numDigits := len(strStone)

	if numDigits%2 == 0 {
		// Split into two stones
		left, _ := strconv.ParseInt(strStone[:numDigits/2], 10, 64)
		rightPart := strings.TrimLeft(strStone[numDigits/2:], "0")
		if rightPart == "" {
			rightPart = "0"
		}
		right, _ := strconv.ParseInt(rightPart, 10, 64)
		return []int64{left, right}
	}

	// Multiply by 2024
	return []int64{stone * 2024}
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
