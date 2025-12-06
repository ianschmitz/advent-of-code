package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"aoc/internal/utils"
)

const (
	part1Expected = 4277556
	part2Expected = 3263827
)

func main() {
	runTest()
	run()
}

func runTest() {
	part1 := solvePart1("day_06_test.txt")
	part2 := solvePart2("day_06_test.txt")
	if part1 != part1Expected {
		log.Fatalf("Part 1 - Expected: %d, receieved: %d", part1Expected, part1)
	}
	if part2 != part2Expected {
		log.Fatalf("Part 2 - Expected: %d, receieved: %d", part2Expected, part2)
	}
}

func run() {
	start := time.Now()
	part1 := solvePart1("day_06.txt")
	fmt.Println("Part 1 duration:", time.Since(start))
	start = time.Now()
	part2 := solvePart2("day_06.txt")
	fmt.Println("Part 2 duration:", time.Since(start))
	fmt.Println("Part 1 - Answer is:", part1)
	fmt.Println("Part 2 - Answer is:", part2)
}

func solvePart1(fileName string) int {
	scanner := utils.GetInputFileLineScanner(fileName)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		lines = append(lines, line)
	}

	rowNums := [][]int{}
	rowOperands := []string{}
	for i, line := range lines {
		splitResults := regexp.MustCompile(`\s+`).Split(line, -1)
		if i == len(lines)-1 {
			rowOperands = splitResults
		} else {
			numResults := []int{}

			for k := range splitResults {
				numResults = append(numResults, utils.StringToInt(splitResults[k]))
			}

			rowNums = append(rowNums, numResults)
		}
	}

	sum := 0
	for i, num := range rowNums[0] {
		result := num
		operand := rowOperands[i]

		for k := 1; k < len(rowNums); k++ {
			otherNum := rowNums[k][i]

			if operand == "*" {
				result *= otherNum
			} else {
				result += otherNum
			}
		}

		sum += result
	}

	return sum
}

func solvePart2(fileName string) int {
	scanner := utils.GetInputFileLineScanner(fileName)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	rowOperands := regexp.MustCompile(`\s+`).Split(lines[len(lines)-1], -1)
	rowOperandsIndex := 0

	sum := 0
	nums := []int{}
	for i := 0; i < len(lines[0]); i++ {
		num := ""

		// Skip over the operand row
		for k := 0; k < len(lines)-1; k++ {
			digit := lines[k][i]
			if digit == ' ' {
				continue
			} else {
				num += string(digit)
			}
		}

		if len(num) > 0 {
			nums = append(nums, utils.StringToInt(num))
		}

		// We've hit the end of this group, let's do the math
		if len(num) == 0 || i == len(lines[0])-1 {
			operand := rowOperands[rowOperandsIndex]
			rowOperandsIndex++

			result := nums[0]
			for k := 1; k < len(nums); k++ {
				otherNum := nums[k]
				if operand == "*" {
					result *= otherNum
				} else {
					result += otherNum
				}
			}

			sum += result
			// Reset for next group
			nums = []int{}
		}
	}

	return sum
}
