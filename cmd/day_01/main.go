package main

import (
	"fmt"
	"log"
	"strconv"

	"aoc/internal/utils"
)

const (
	part1Expected = 3
	part2Expected = 6
	dialSize      = 100
)

func main() {
	runTest()
	run()
}

func runTest() {
	part1, part2 := solve("day_01_test.txt")
	if part1 != part1Expected {
		log.Fatalf("Part 1 - Expected: %d, receieved: %d", part1Expected, part1)
	}
	if part2 != part2Expected {
		log.Fatalf("Part 2 - Expected: %d, receieved: %d", part2Expected, part2)
	}
}

func run() {
	part1, part2 := solve("day_01.txt")
	fmt.Println("Part 1 - Answer is:", part1)
	fmt.Println("Part 2 - Answer is:", part2)
}

func solve(fileName string) (int, int) {
	scanner := utils.GetInputFileLineScanner(fileName)

	dial := 50
	part1 := 0
	part2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0:1]
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal("Error converting line to int:", err)
		}

		if direction == "L" {
			num = -num
		}

		newPosition := dial + num

		// Using `max()` instead of converting to/from float for `math.Abs`
		part2 += max(newPosition/dialSize, -(newPosition / dialSize))
		// Note: If dial was at 0, and we just moved 1 to the left, we wouldn't count that as having hit 0
		if newPosition <= 0 && dial != 0 {
			part2++
		}

		// `newPosition` might have passed over `0` one or more times.
		// By calculating the remainder we don't care how many times it passed 0 but
		// just the final position it landed on.
		dial = newPosition % dialSize
		if dial < 0 {
			dial = dialSize + dial
		}

		if dial == 0 {
			part1++
		}
	}

	return part1, part2
}
