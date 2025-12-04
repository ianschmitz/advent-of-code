package main

import (
	"fmt"
	"log"
	"strings"

	"aoc/internal/utils"
)

const (
	part1Expected = 13
	part2Expected = 43
)

func main() {
	runTest()
	run()
}

func runTest() {
	part1 := solvePart1("day_04_test.txt")
	part2 := solvePart2("day_04_test.txt")
	if part1 != part1Expected {
		log.Fatalf("Part 1 - Expected: %d, receieved: %d", part1Expected, part1)
	}
	if part2 != part2Expected {
		log.Fatalf("Part 2 - Expected: %d, receieved: %d", part2Expected, part2)
	}
}

func run() {
	part1 := solvePart1("day_04.txt")
	part2 := solvePart2("day_04.txt")
	fmt.Println("Part 1 - Answer is:", part1)
	fmt.Println("Part 2 - Answer is:", part2)
}

type coordinate struct {
	col int
	row int
}

func solvePart1(fileName string) int {
	grid := getGridFromFileName(fileName)

	accessibleRollPositions := []coordinate{}

	for rowIndex, row := range grid {
		for colIndex, value := range row {
			if value != 1 {
				continue
			}

			if isSafeToRemove(grid, colIndex, rowIndex) {
				accessibleRollPositions = append(accessibleRollPositions, coordinate{col: colIndex, row: rowIndex})
			}
		}
	}

	return len(accessibleRollPositions)
}

func solvePart2(fileName string) int {
	grid := getGridFromFileName(fileName)

	removedRollPositions := []coordinate{}

	keepTrying := true
	for keepTrying {
		removedRollThisPass := false

		for rowIndex, row := range grid {
			for colIndex, value := range row {
				if value != 1 {
					continue
				}

				if isSafeToRemove(grid, colIndex, rowIndex) {
					removedRollThisPass = true
					removedRollPositions = append(removedRollPositions, coordinate{col: colIndex, row: rowIndex})
					grid[rowIndex][colIndex] = 0
				}
			}
		}

		// Reset so we can try a full scan again
		if !removedRollThisPass {
			keepTrying = false
		}
	}

	return len(removedRollPositions)
}

func getGridFromFileName(fileName string) [][]int {
	scanner := utils.GetInputFileLineScanner(fileName)

	grid := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")

		row := make([]int, len(chars))

		for i, char := range chars {
			if char == "@" {
				row[i] = 1
			} else {
				row[i] = 0
			}
		}

		grid = append(grid, row)
	}

	return grid
}

func safeGridAccess(grid [][]int, col int, row int) int {
	if col < 0 || col >= len(grid[0]) || row < 0 || row >= len(grid) {
		return 0
	}
	return grid[row][col]
}

func isSafeToRemove(grid [][]int, colIndex int, rowIndex int) bool {
	topL := safeGridAccess(grid, colIndex-1, rowIndex-1)
	top := safeGridAccess(grid, colIndex, rowIndex-1)
	topR := safeGridAccess(grid, colIndex+1, rowIndex-1)
	l := safeGridAccess(grid, colIndex-1, rowIndex)
	r := safeGridAccess(grid, colIndex+1, rowIndex)
	botL := safeGridAccess(grid, colIndex-1, rowIndex+1)
	bot := safeGridAccess(grid, colIndex, rowIndex+1)
	botR := safeGridAccess(grid, colIndex+1, rowIndex+1)

	adjOccupied := topL + top + topR + l + r + botL + bot + botR

	return adjOccupied < 4
}
