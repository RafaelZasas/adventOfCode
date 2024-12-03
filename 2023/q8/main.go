package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

type Node struct {
	left    string
	right   string
	point   string
	nodeIdx int // used for part 2
}

func (n Node) String() string {
	return fmt.Sprintf("%s: (%s, %s) idx: %d", n.point, n.left, n.right, n.nodeIdx)
}

var nodes = make(map[string]Node)
var instructions string

func init() {

	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	instructions = scanner.Text()

	// move past the empty line
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		point := strings.Trim(parts[0], " ")

		directions := strings.Split(parts[1], ",")
		left := directions[0][2:]
		right := directions[1][1:4]

		n := Node{left: left, right: right, point: point}
		nodes[point] = n
	}

}

func main() {
	//q1()
	q2()
}

// Walk through the directions from point to point until you reach the point ZZZ
func q1() {
	currentNode := nodes["AAA"]

	steps := 0

	for i := 0; currentNode.point != "ZZZ"; i++ {

		steps++

		direction := string(instructions[i])

		if direction == "R" {
			currentNode = nodes[currentNode.right]
		}

		if direction == "L" {
			currentNode = nodes[currentNode.left]
		}

		// Ensure that the loop continues to the beginning of the instructions when the end is reached
		if i == len(instructions)-1 {
			i = -1
		}

	}

	fmt.Printf("Q1) num steps: %d\n", steps)

}

// starting nodes are all nodes with a point that ends in A
// walk through directions and update all of the current nodes to the node indicated by the
// direction.
// end when all nodes have points ending in Z
func q2() {

	currentNodes := make(map[string]Node)

	// this will map nodeIndexes to the number of steps it took to get to Z
	nodeCycles := make(map[int]int)

	nodeIdx := 0
	for _, node := range nodes {
		if strings.HasSuffix(node.point, "A") {
			// using this to ensure that we only add non entered node paths to the
			// list of cycles with which we will calculate the LCM
			node.nodeIdx = nodeIdx
			currentNodes[node.point] = node
			nodeIdx++
		}
	}

	steps := 0
	index := 0

	// this loop will never terminate in any reasonable time
	// but each node has a specific period in which the point will
	// end in Z.  So, we can just find the period for each node and
	// then find the least common multiple of all of the periods
	for !allNodesEndWithZ(currentNodes) {
		steps++

		// assumption: there are only 6 nodes that end in A
		// this is true for my input, but may not be true for all inputs
		if len(nodeCycles) == 6 {
			break
		}

		for point, node := range currentNodes {

			if strings.HasSuffix(node.point, "Z") {
				// ensure that we are not adding duplicates of the same node path
				// by checking if the node path has already been added to the map
				// ie. one path can reach a z twice before another path reaches it once,
				// in which case we would have 6 node path cycles in the map, but only
				// 5 unique node paths

				if _, ok := nodeCycles[node.nodeIdx]; !ok {
					// dont ask me about the -1, it just works
					nodeCycles[node.nodeIdx] = steps - 1
					fmt.Printf("node %v: %d\n", node, steps)
				}

			}

			direction := string(instructions[index])

			if direction == "R" {
				// have to destructure it this way since we are not allowed to do currentNodes[point].nodeIdx = node.nodeIdx
				currentNodes[point] = Node{nodes[node.right].left, nodes[node.right].right, nodes[node.right].point, node.nodeIdx}
			}

			if direction == "L" {
				currentNodes[point] = Node{nodes[node.left].left, nodes[node.left].right, nodes[node.left].point, node.nodeIdx}
			}

		}

		// Ensure that the loop continues to the beginning of the instructions when the end is reached
		if index == len(instructions)-1 {
			index = 0
		} else {
			index++
		}
	}

	res := lcm([]int{nodeCycles[0], nodeCycles[1], nodeCycles[2], nodeCycles[3], nodeCycles[4], nodeCycles[5]})

	fmt.Println(nodeCycles)
	fmt.Printf("Q2) num steps: %d\n", res)

}

// Helper Functions

func allNodesEndWithZ(n map[string]Node) bool {
	for _, node := range n {
		if !strings.HasSuffix(node.point, "Z") {
			return false
		}
	}
	return true
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := big.NewInt(int64(nums[0]))
	for _, num := range nums[1:] {
		bigNum := big.NewInt(int64(num))
		gcdVal := big.NewInt(int64(gcd(int(result.Int64()), num)))
		lcmVal := new(big.Int).Div(new(big.Int).Mul(result, bigNum), gcdVal)
		result.Set(lcmVal)
	}

	return int(result.Int64())
}
