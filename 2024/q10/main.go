package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

var fileName = "example.txt"

var fileBuffer []byte

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

var trailHeads []Point

var grid [][]int

func init() {

	start := time.Now()
	defer func() {
		fmt.Println("Init took: ", time.Since(start))
	}()

	grid = make([][]int, 0)

	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	// read file into buffer and split by two new lines
	fileBuffer, err := io.ReadAll(file)
	check(err)

	rows := strings.Split(string(fileBuffer), "\n")

	for i, row := range rows {
		if row == "" {
			continue
		}

		gridRow := make([]int, 0)

		for j, char := range row {

			num, err := strconv.Atoi(string(char))
			check(err)

			gridRow = append(gridRow, num)

			if num == 0 {
				trailHeads = append(trailHeads, Point{j, i})
			}
		}

		grid = append(grid, gridRow)

	}
}

func partOne() {

	start := time.Now()
	defer func() {
		fmt.Println("Part one took: ", time.Since(start))
	}()

	totalScore := 0
	for _, th := range trailHeads {
		score := bfs(th)
		totalScore += score
	}

	fmt.Println("Total trailhead score:", totalScore)
}

func partTwo() {
	start := time.Now()
	defer func() {
		fmt.Println("Part two took: ", time.Since(start))
	}()

	total := 0

	for _, th := range trailHeads {
		visited := make(map[Point]bool)
		total += dfsCountPaths(th, visited)

	}

	fmt.Println("Total number of paths to 9:", total)

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

func dfsCountPaths(current Point, visited map[Point]bool) int {
	if grid[current.y][current.x] == 9 {
		return 1 // Found a valid trail to a 9
	}

	visited[current] = true
	defer func() { visited[current] = false }() // Backtrack after recursive call

	count := 0
	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // Up, Right, Down, Left

	for _, dir := range directions {
		neighbor := Point{current.x + dir.x, current.y + dir.y}
		if isValidNeighbor(current, neighbor, visited) {
			count += dfsCountPaths(neighbor, visited) // Recursively explore paths
		}
	}

	return count
}

// Builds a graph of the grid to search for next possible paths
func bfs(start Point) int {
	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // Up, Right, Down, Left
	queue := []Point{start}
	visited := make(map[Point]bool)
	visited[start] = true

	reachableNines := make(map[Point]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			neighbor := Point{current.x + dir.x, current.y + dir.y}

			if isValidNeighbor(current, neighbor, visited) {
				visited[neighbor] = true
				queue = append(queue, neighbor)

				if grid[neighbor.y][neighbor.x] == 9 {
					reachableNines[neighbor] = true
				}
			}
		}
	}

	return len(reachableNines)
}

// Helper function to check if a neighbor is valid
func isValidNeighbor(current, neighbor Point, visited map[Point]bool) bool {
	if neighbor.y < 0 || neighbor.y >= len(grid) || neighbor.x < 0 || neighbor.x >= len(grid[0]) {
		return false // Out of bounds
	}

	if visited[neighbor] {
		return false // Already visited
	}

	// Ensure height increases by 1
	return grid[neighbor.y][neighbor.x] == grid[current.y][current.x]+1
}
