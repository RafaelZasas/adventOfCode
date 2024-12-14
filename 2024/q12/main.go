package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var fileName = "example.txt"

var fileBuffer []byte

var grid [][]rune

var regionPrices = map[rune]int{}

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

	for _, char := range fileBuffer {
		if char == '\n' {
			continue
		}
		regionPrices[rune(char)] = 0
	}

	rows := strings.Split(string(fileBuffer), "\n")

	for _, row := range rows {
		if row == "" {
			continue
		}

		gridRow := []rune(row)
		grid = append(grid, gridRow)

	}
}

func partOne() {
	start := time.Now()
	defer func() {
		fmt.Println("Part one took: ", time.Since(start))
	}()

	visited := make([][]bool, len(grid))
	for i := range grid {
		visited[i] = make([]bool, len(grid[i]))
	}

	totalPrice := 0

	for i, row := range grid {
		for j, plant := range row {
			if visited[i][j] {
				continue
			}

			// Flood-fill to find the region
			regionArea, regionPerimeter := exploreRegion(i, j, plant, visited)
			price := regionArea * regionPerimeter
			totalPrice += price
		}
	}

	fmt.Println("Total Price:", totalPrice)
}

func exploreRegion(x, y int, plant rune, visited [][]bool) (int, int) {
	// Directions for traversing up, down, left, right
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	stack := [][2]int{{x, y}}
	visited[x][y] = true
	area := 0
	perimeter := 0

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		cx, cy := current[0], current[1]
		area++

		// Check all 4 directions
		for _, dir := range directions {
			nx, ny := cx+dir[0], cy+dir[1]

			// If out of bounds, add to the perimeter
			if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid[0]) {
				perimeter++
				continue
			}

			// If the neighbor is not the same plant, add to the perimeter
			if grid[nx][ny] != plant {
				perimeter++
				continue
			}

			// If the neighbor is unvisited and part of the region, visit it
			if !visited[nx][ny] {
				visited[nx][ny] = true
				stack = append(stack, [2]int{nx, ny})
			}
		}
	}

	return area, perimeter
}

func partTwo() {
	start := time.Now()
	defer func() {
		fmt.Println("Part two took: ", time.Since(start))
	}()
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

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func printRegionPrices() {
	for k, v := range regionPrices {
		fmt.Println(string(k), v)
	}
}
