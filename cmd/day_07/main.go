package main

import (
	"fmt"
	"log"
	"time"

	"aoc/internal/utils"
)

const (
	part1Expected = 21
	part2Expected = 40
)

func main() {
	runTest()
	run()
}

func runTest() {
	part1 := solvePart1("day_07_test.txt")
	part2 := solvePart2("day_07_test.txt")
	if part1 != part1Expected {
		log.Fatalf("Part 1 - Expected: %d, receieved: %d", part1Expected, part1)
	}
	if part2 != part2Expected {
		log.Fatalf("Part 2 - Expected: %d, receieved: %d", part2Expected, part2)
	}
}

func run() {
	start := time.Now()
	part1 := solvePart1("day_07.txt")
	fmt.Println("Part 1 duration:", time.Since(start))
	start = time.Now()
	part2 := solvePart2("day_07.txt")
	fmt.Println("Part 2 duration:", time.Since(start))
	fmt.Println("Part 1 - Answer is:", part1)
	fmt.Println("Part 2 - Answer is:", part2)
}

func solvePart1(fileName string) int {
	scanner := utils.GetInputFileLineScanner(fileName)

	timesSplit := 0
	beamIndexes := map[int]bool{}

	for scanner.Scan() {
		line := scanner.Text()

		for i, char := range line {
			if char == 'S' {
				beamIndexes[i] = true

				// First line always has only one "S"
				break
			}

			_, beamExists := beamIndexes[i]

			// This is a splitter and a beam is currently above it
			if char == '^' && beamExists {
				timesSplit++

				delete(beamIndexes, i)
				// This is potentially suboptimal since we're doing two appends with a single item
				// instead of one append with two items
				if i > 0 {
					beamIndexes[i-1] = true
				}
				if i < len(line)-1 {
					beamIndexes[i+1] = true
				}
			}
		}
	}

	return timesSplit
}

func solvePart2(fileName string) int {
	scanner := utils.GetInputFileLineScanner(fileName)

	beamIndexes := map[int]int{}

	for scanner.Scan() {
		line := scanner.Text()

		for i, char := range line {
			if char == 'S' {
				beamIndexes[i] = 1

				// First line always has only one "S"
				break
			}

			value, ok := beamIndexes[i]
			// This is a splitter and a beam is currently above it
			if char == '^' && ok {
				delete(beamIndexes, i)

				// There's already a beam here. Add together the value of both beams
				if leftVal, ok := beamIndexes[i-1]; ok {
					beamIndexes[i-1] = value + leftVal
				} else {
					beamIndexes[i-1] = value
				}

				if rightVal, ok := beamIndexes[i+1]; ok {
					beamIndexes[i+1] = value + rightVal
				} else {
					beamIndexes[i+1] = value
				}
			}
		}
	}

	sum := 0

	for _, value := range beamIndexes {
		sum += value
	}

	return sum
}
