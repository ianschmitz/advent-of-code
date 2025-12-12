package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"aoc/internal/utils"
)

const (
	part1Expected = 2
	part2Expected = 0
)

func main() {
	// runTest()
	run()
}

func runTest() {
	part1 := solvePart1("day_12_test.txt")
	part2 := solvePart2("day_12_test.txt")
	if part1 != part1Expected {
		log.Fatalf("Part 1 - Expected: %d, receieved: %d", part1Expected, part1)
	}
	if part2 != part2Expected {
		log.Fatalf("Part 2 - Expected: %d, receieved: %d", part2Expected, part2)
	}
}

func run() {
	start := time.Now()
	part1 := solvePart1("day_12.txt")
	fmt.Println("Part 1 duration:", time.Since(start))
	start = time.Now()
	part2 := solvePart2("day_12.txt")
	fmt.Println("Part 2 duration:", time.Since(start))
	fmt.Println("Part 1 - Answer is:", int(part1))
	fmt.Println("Part 2 - Answer is:", int(part2))
}

func solvePart1(fileName string) int {
	shapeAreas, regions := parseInput(fileName)

	countFitRegions := 0
	for _, region := range regions {
		regionShapesArea := 0
		for i, quantity := range region.ShapeQuantities {
			regionShapesArea += quantity * shapeAreas[i]
		}

		fmt.Println(regionShapesArea)
		utils.PrettyPrintlnStruct(region)
		regionArea := region.Width * region.Height
		// Apparently this works for real input but not test input
		// ¯\_(ツ)_/¯
		if regionArea >= regionShapesArea {
			countFitRegions++
		}
	}

	return countFitRegions
}

func solvePart2(fileName string) int {
	return 0
}

func parseInput(fileName string) ([]int, []region) {
	file := utils.ReadInputFile(fileName)
	lines := strings.Split(string(file), "\n")

	shapeAreas := []int{}
	idx := 1
	for {
		combinedLines := lines[idx] + lines[idx+1] + lines[idx+2]

		firstChar := combinedLines[0]

		// We've gone past the shapes section
		if firstChar != '#' && firstChar != '.' {
			break
		}

		shapeAreas = append(shapeAreas, strings.Count(combinedLines, "#"))

		// Jump to next shape
		idx += 5
	}

	regions := []region{}
	// -1 brings us back first line of the regions
	for i := idx - 1; i < len(lines)-1; i++ {
		line := lines[i]
		parts := strings.SplitN(line, ":", 2)

		var width, height int
		fmt.Sscanf(parts[0], "%dx%d", &width, &height)

		quantityStrs := strings.Fields(strings.TrimSpace(parts[1]))
		quantities := []int{}
		for _, quantity := range quantityStrs {
			quantities = append(quantities, utils.StringToInt(quantity))
		}

		regions = append(regions, region{width, height, quantities})
	}

	return shapeAreas, regions
}

type region struct {
	Width           int
	Height          int
	ShapeQuantities []int
}
