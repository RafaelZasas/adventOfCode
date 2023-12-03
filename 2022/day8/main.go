package q8

import "os"

var numVisibleTrees int
var treeGrid [][]int

func init() {
	// open the file input.txt for reading

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// read the file and store the values in a 2D array

	var row []int
	var grid [][]int

	for {
		var num int
		_, err := fmt.Fscanf(file, "%d\n", &num)
		if err != nil {
			break
		}
		row = append(row, num)
		if len(row) == 31 {
			grid = append(grid, row)
			row = nil
		}
	}

}
