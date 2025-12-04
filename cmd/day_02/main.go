package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"aoc/internal/utils"
)

const (
	part1Expected = 1227775554
	part2Expected = 4174379265
)

func main() {
	runTest()
	run()
}

func runTest() {
	part1 := solvePart1("day_02_test.txt")
	part2 := solvePart2("day_02_test.txt")
	if part1 != part1Expected {
		log.Fatalf("Part 1 - Expected: %d, receieved: %d", part1Expected, part1)
	}
	if part2 != part2Expected {
		log.Fatalf("Part 2 - Expected: %d, receieved: %d", part2Expected, part2)
	}
}

func run() {
	part1 := solvePart1("day_02.txt")
	part2 := solvePart2("day_02.txt")
	fmt.Println("Part 1 - Answer is:", part1)
	fmt.Println("Part 2 - Answer is:", part2)
}

func solvePart1(fileName string) int {
	file := string(utils.ReadInputFile(fileName))
	file = strings.TrimSuffix(file, "\n")

	idRanges := strings.Split(string(file), ",")

	invalidIds := []int{}

	for _, idRange := range idRanges {
		splitIds := strings.Split(idRange, "-")
		start, end := utils.StringToInt(splitIds[0]), utils.StringToInt(splitIds[1])

		for id := start; id <= end; id++ {
			if !isValidPart1Id(id) {
				invalidIds = append(invalidIds, id)
			}
		}
	}

	answer := 0
	for _, id := range invalidIds {
		answer += id
	}

	return answer
}

func solvePart2(fileName string) int {
	file := string(utils.ReadInputFile(fileName))
	file = strings.TrimSuffix(file, "\n")

	idRanges := strings.Split(string(file), ",")

	invalidIds := []int{}

	for _, idRange := range idRanges {
		splitIds := strings.Split(idRange, "-")
		start, end := utils.StringToInt(splitIds[0]), utils.StringToInt(splitIds[1])

		for id := start; id <= end; id++ {
			if !isValidPart2Id(id) {
				invalidIds = append(invalidIds, id)
			}
		}
	}

	answer := 0
	for _, id := range invalidIds {
		answer += id
	}

	return answer
}

func isValidPart1Id(id int) bool {
	strId := strconv.Itoa(id)
	// Odd number of digits can never repeat as instructions say:
	// made **only** of some sequence of digits repeated **twice**
	if len(strId)%2 == 1 {
		return true
	}

	mid := len(strId) / 2
	firstHalf := strId[:mid]
	secondHalf := strId[mid:]

	return firstHalf != secondHalf
}

func isValidPart2Id(id int) bool {
	strId := strconv.Itoa(id)

	for i := 0; i < len(strId)/2; i++ {
		slices := splitIntoNCharSlices(strId, i+1)

		if len(slices) < 2 {
			continue
		}
		allEqual := allItemsInSliceEqual(slices)
		if allEqual {
			return false
		}
	}

	return true
}

func allItemsInSliceEqual(slice []string) bool {
	for _, v := range slice {
		if v != slice[0] {
			return false
		}
	}
	return true
}

func splitIntoNCharSlices(s string, sliceSize int) []string {
	var result []string
	for i := 0; i < len(s); i += sliceSize {
		// Handle the case where the last chunk might be less than 2 characters
		end := min(i+sliceSize, len(s))

		result = append(result, string(s[i:end]))
	}
	return result
}
