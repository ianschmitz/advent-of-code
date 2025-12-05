package main

import (
	"fmt"
	"log"
	"slices"
	"strings"
	"time"

	"aoc/internal/utils"
)

const (
	part1Expected = 3
	part2Expected = 14
)

func main() {
	runTest()
	run()
}

func runTest() {
	part1 := solvePart1("day_05_test.txt")
	part2 := solvePart2("day_05_test.txt")
	if part1 != part1Expected {
		log.Fatalf("Part 1 - Expected: %d, receieved: %d", part1Expected, part1)
	}
	if part2 != part2Expected {
		log.Fatalf("Part 2 - Expected: %d, receieved: %d", part2Expected, part2)
	}
}

func run() {
	start := time.Now()
	part1 := solvePart1("day_05.txt")
	fmt.Println("Part 1 duration:", time.Since(start))
	start = time.Now()
	part2 := solvePart2("day_05.txt")
	fmt.Println("Part 2 duration:", time.Since(start))
	fmt.Println("Part 1 - Answer is:", part1)
	fmt.Println("Part 2 - Answer is:", part2)
}

func solvePart1(fileName string) int {
	scanner := utils.GetInputFileLineScanner(fileName)

	freshIngredientIds := [][2]int{}
	countAvailableFreshIngredients := 0
	reachedAvailableIngredients := false

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			reachedAvailableIngredients = true
			continue
		}

		if !reachedAvailableIngredients {
			freshIngredientIds = append(freshIngredientIds, getFreshIngredientRange(line))
			continue
		}

		ingredientId := utils.StringToInt(line)
		isFresh := slices.ContainsFunc(freshIngredientIds, func(idRange [2]int) bool {
			return ingredientId >= idRange[0] && ingredientId <= idRange[1]
		})

		if isFresh {
			countAvailableFreshIngredients++
		}
	}

	return countAvailableFreshIngredients
}

func solvePart2(fileName string) int {
	scanner := utils.GetInputFileLineScanner(fileName)

	freshIngredientIds := [][2]int{}
	for scanner.Scan() {
		line := scanner.Text()

		// We're done scanning
		if line == "" {
			break
		}

		freshIngredientIds = append(freshIngredientIds, getFreshIngredientRange(line))
	}

	countFreshIngredients := 0
	for i, idRange := range freshIngredientIds {
		min := idRange[0]
		max := idRange[1]

		// Check all past ranges to see if we've already covered some of this range
		for k := range i {
			otherMin := freshIngredientIds[k][0]
			otherMax := freshIngredientIds[k][1]

			if min >= otherMin && min <= otherMax {
				min = otherMax + 1
			}
			if max >= otherMin && max <= otherMax {
				max = otherMin - 1
			}
		}

		rangeCount := max - min + 1

		if rangeCount > 0 {
			countFreshIngredients += rangeCount
		}
	}

	return countFreshIngredients
}

func getFreshIngredientRange(line string) [2]int {
	freshRange := strings.Split(line, "-")
	min := utils.StringToInt(freshRange[0])
	max := utils.StringToInt(freshRange[1])

	return [2]int{min, max}
}
