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
	part2Expected = 0
)

func main() {
	runTest()
	run()
}

func runTest() {
	part1, part2 := solve("day_02_test.txt")
	if part1 != part1Expected {
		log.Fatalf("Part 1 - Expected: %d, receieved: %d", part1Expected, part1)
	}
	if part2 != part2Expected {
		log.Fatalf("Part 2 - Expected: %d, receieved: %d", part2Expected, part2)
	}
}

func run() {
	part1, part2 := solve("day_02.txt")
	fmt.Println("Part 1 - Answer is:", part1)
	fmt.Println("Part 2 - Answer is:", part2)
}

func solve(fileName string) (int, int) {
	file := string(utils.ReadInputFile(fileName))
	file = strings.TrimSuffix(file, "\n")

	idRanges := strings.Split(string(file), ",")

	part1InvalidIds := []int{}

	for _, idRange := range idRanges {
		splitIds := strings.Split(idRange, "-")
		start, end := stringToInt(splitIds[0]), stringToInt(splitIds[1])

		for id := start; id <= end; id++ {
			if !isValidPart1Id(id) {
				part1InvalidIds = append(part1InvalidIds, id)
			}
		}
	}

	part1 := 0

	for _, id := range part1InvalidIds {
		part1 += id
	}

	part2 := 0

	return part1, part2
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

func stringToInt(stringNum string) int {
	num, err := strconv.Atoi(stringNum)
	if err != nil {
		log.Fatal("Error converting string to int:", err)
	}
	return num
}
