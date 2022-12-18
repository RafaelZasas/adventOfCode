package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type Dir struct {
	Size   int
	Name   string
	Files  []*File
	Dirs   []*Dir
	parent *Dir
}

func (d *Dir) String() string {
	return fmt.Sprintf("Dir: %v (%v)\nFiles:\n\t%v\n Dirs:\n\t%v", d.Name, d.Size, d.Files, d.Dirs)
}

type File struct {
	Size   int
	Name   string
	parent *Dir
}

func (f *File) String() string {
	return fmt.Sprintf("\n\tFile: %v %d", f.Name, f.Size)
}

type Instruction struct {
	cmd      string
	location string
}

type Dirs []*Dir

func (d Dirs) root() *Dir {
	var rootDir *Dir
	for _, dir := range d {
		if dir.Name == "/" {
			rootDir = dir
		}
	}
	return rootDir
}

var (
	dirs       Dirs
	currentDir *Dir
)

const (
	totalDiskSpace, unusedSpaceNeeded int = 70000000, 30000000
)

// isCmd returns whether a line is a command (prefixed with $) or not
func isCmd(line string) bool {
	return string(line[0]) == "$"
}

// isDir returns whether a line is a directory (starting with dir) or not
func isDir(line string) bool {
	return line[:3] == "dir"
}

// readLines recursively moves through lines in the scanner
// and calls itself when new files and directories are listed (ls line)
func readLines(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		if isCmd(line) {
			if len(line) == 4 { // only ls lines are length 4
				readLines(scanner)
			} else { // the only other cmd is cd <location>
				inst := Instruction{
					cmd:      line[2:4],
					location: line[5:],
				}
				if inst.location == ".." { // handle move up dir
					// at this point we know that all of the files in the dir have been acconted for
					currentDir.parent.Size += currentDir.Size // so we can add the size of the current dir to its parent
					// keeping in mind that the last dir hasnt been accounted for
					currentDir = currentDir.parent
				} else if inst.location == "/" { // handle move root
					currentDir = dirs.root()
				} else { // handle move to child dir

					for _, dir := range currentDir.Dirs {
						// change the current dir as per instruction
						if dir.Name == inst.location {
							currentDir = dir
						}
					}
				}
			}
		} else { // here we know that the line is either a file or a dir
			if isDir(line) {
				d := Dir{Size: 0, Name: line[4:], parent: currentDir}
				d.Name = line[4:]
				currentDir.Dirs = append(currentDir.Dirs, &d)
			} else { // here we know that the line is a file
				f := File{}
				f.parent = currentDir
				re := regexp.MustCompile("[0-9]+") // keep only numbers for the file size
				f.Size, _ = strconv.Atoi(re.FindAllString(line, -1)[0])

				re = regexp.MustCompile("[^a-zA-Z.]+") // keep all letters and period for the file name
				f.Name = re.ReplaceAllString(line, "")

				currentDir.Files = append(currentDir.Files, &f)
				currentDir.Size += f.Size // increase size of the dir

			}
		}
	}
}

func init() {
	// need to have at least one dir (root) to start
	rootDir := Dir{
		parent: nil,
		Name:   "/",
	}
	currentDir = &rootDir // start off by making current dir the root

	dirs = append(dirs, &rootDir)

	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	readLines(scanner)

	// now need to take into account the last dir
	// and add its size to the size of the total
	dirs.root().Size += dirs.root().Dirs[len(dirs.root().Dirs)-1].Size
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// countDirSizes recursively searches the file system tree
// and returns a list of the directory sizes in the fs
func countDirSizes(d *Dir) []int {
	var sizes []int
	sizes = append(sizes, d.Size)
	for _, dir := range d.Dirs {
		sizes = append(sizes, countDirSizes(dir)...)
	}

	return sizes
}

// findDirectoriesOfAtMost will find and return all the directories which have
// an upper bound of the given size
func findDirectoriesOfAtMost(dir *Dir, size int) []*Dir {
	var dirs []*Dir
	if dir.Size <= size {
		dirs = append(dirs, dir)
	}

	for _, childDir := range dir.Dirs {
		dirs = append(dirs, findDirectoriesOfAtMost(childDir, size)...)
	}

	return dirs
}

func main() {
	// Pretty Print in json and dump to file for debugging
	jsonDirs, err := json.MarshalIndent(dirs, "", "  ")
	check(err)
	_ = ioutil.WriteFile("output.json", jsonDirs, 0644)

	// Part one - Sum dirs with the upper bound
	atMostDirs := findDirectoriesOfAtMost(dirs.root(), 100000)
	totalSize := 0
	for _, dir := range atMostDirs {
		totalSize += dir.Size
	}
	fmt.Printf("Part One : %v\n", totalSize)

	// part two - finding the smallest dir to free space

	sizes := countDirSizes(dirs.root())
	sort.Ints(sizes)

	sizeNeeded := unusedSpaceNeeded - (totalDiskSpace - dirs.root().Size)
	var smallestPossible int
	for _, size := range sizes {
		if size >= sizeNeeded {
			smallestPossible = size
			break
		}
	}
	fmt.Printf("Part Two (Smallest Possible Dir):\nSize Needed: %v\nSmallest dir: %v", sizeNeeded, smallestPossible)
}
