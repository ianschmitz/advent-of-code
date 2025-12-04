package main

import (
	"fmt"
	"log"
	"strings"

	"aoc/internal/utils"
)

const (
	part1Expected = 357
	part2Expected = 3121910778619
)

func main() {
	runTest()
	run()
}

func runTest() {
	part1 := solvePart1("day_03_test.txt")
	part2 := solvePart2("day_03_test.txt")
	if part1 != part1Expected {
		log.Fatalf("Part 1 - Expected: %d, receieved: %d", part1Expected, part1)
	}
	if part2 != part2Expected {
		log.Fatalf("Part 2 - Expected: %d, receieved: %d", part2Expected, part2)
	}
}

func run() {
	part1 := solvePart1("day_03.txt")
	part2 := solvePart2("day_03.txt")
	fmt.Println("Part 1 - Answer is:", part1)
	fmt.Println("Part 2 - Answer is:", part2)
}

func solvePart1(fileName string) int {
	scanner := utils.GetInputFileLineScanner(fileName)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		strNums := strings.Split(line, "")

		// Note: we could have converted the strings to integers, but it's unnecessary since
		// we're always comparing two strings with exactly two numbers within.

		largestSeen := "00"

		// Go through each number in the line and find the largest number by
		// combining it with numbers to its right.
		// Don't bother checking last number as there isn't anything to the right of it.
		for i := 0; i < len(strNums)-1; i++ {
			for k := i + 1; k < len(strNums); k++ {
				combined := strNums[i] + strNums[k]

				if combined > largestSeen {
					largestSeen = combined
				}
			}
		}

		sum += utils.StringToInt(largestSeen)
	}

	return sum
}

func solvePart2(fileName string) int {
	scanner := utils.GetInputFileLineScanner(fileName)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		strNums := strings.Split(line, "")

		lineStack := []string{strNums[0]}

		// Note: we could have converted the strings to integers, but it's unnecessary since
		// we're always comparing two single digit strings. We'll convert at the end

		// Note we initialized the lineStack above with the first number as a starting point
		for i := 1; i < len(strNums); i++ {
			currentNum := strNums[i]

			// Work our way back from the end of the stack and check if how far down the stack
			// we can insert this number (if at all).
			for k := len(lineStack) - 1; k >= -1; k-- {
				// Nothing to compare to, shove it at the start of a new list
				if k == -1 {
					lineStack = []string{currentNum}
					break
				}

				stackNum := lineStack[k]

				availableNumbersLeftToFill := len(strNums) - i
				numbersNeededToFillCurrentPosition := 12 - k

				// We've reached a point in the stack where we know we can no longer search for a
				// place to insert the current number.
				if currentNum <= stackNum || availableNumbersLeftToFill < numbersNeededToFillCurrentPosition {
					// Insert just after stack position.
					insertionPosition := k + 1

					// Can only have 12 numbers in the stack
					if insertionPosition > 11 {
						break
					}

					// Create new stack slice excluding all items after the point where we want to insert.
					lineStack = lineStack[:insertionPosition]
					lineStack = append(lineStack, currentNum)

					break
				}
			}
		}

		strNum := ""

		for _, stackItem := range lineStack {
			strNum += stackItem
		}

		sum += utils.StringToInt(strNum)
	}

	return sum
}
