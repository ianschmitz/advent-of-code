package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"aoc/internal/utils"
)

const (
	part1Expected = 5
	part2Expected = 2
)

func main() {
	runTest()
	run()
}

func runTest() {
	part1 := solvePart1("day_11_test.txt")
	part2 := solvePart2("day_11_test_2.txt")
	if part1 != part1Expected {
		log.Fatalf("Part 1 - Expected: %d, receieved: %d", part1Expected, part1)
	}
	if part2 != part2Expected {
		log.Fatalf("Part 2 - Expected: %d, receieved: %d", part2Expected, part2)
	}
}

func run() {
	start := time.Now()
	part1 := solvePart1("day_11.txt")
	fmt.Println("Part 1 duration:", time.Since(start))
	start = time.Now()
	part2 := solvePart2("day_11.txt")
	fmt.Println("Part 2 duration:", time.Since(start))
	fmt.Println("Part 1 - Answer is:", part1)
	fmt.Println("Part 2 - Answer is:", part2)
}

func solvePart1(fileName string) int {
	scanner := utils.GetInputFileLineScanner(fileName)

	deviceMap := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ":")
		key := split[0]
		vals := strings.Split(strings.TrimSpace(split[1]), " ")
		deviceMap[key] = vals
	}

	return walkDagPart1(deviceMap)
}

func solvePart2(fileName string) int {
	scanner := utils.GetInputFileLineScanner(fileName)

	deviceMap := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ":")
		key := split[0]
		vals := strings.Split(strings.TrimSpace(split[1]), " ")
		deviceMap[key] = vals
	}

	return walkDagPart2(deviceMap)
}

func walkDagPart1(deviceMap map[string][]string) int {
	var walkDag func(key string) int

	walkDag = func(key string) int {
		if deviceMap[key][0] == "out" {
			return 1
		}

		total := 0
		for _, otherKey := range deviceMap[key] {
			total += walkDag(otherKey)
		}

		return total
	}

	return walkDag("you")
}

func walkDagPart2(deviceMap map[string][]string) int {
	memo := map[string]int{}

	var walkDag func(key string, dacSeen bool, fftSeen bool) int

	walkDag = func(key string, dacSeen bool, fftSeen bool) int {
		if key == "out" {
			if dacSeen && fftSeen {
				return 1
			} else {
				return 0
			}
		}

		// This is a weird looking key, but we're saying that for anyone that arrives with the same
		// values of 'dacSeen' and 'fftSeen' and there's an item in cache it represents the total
		// possible paths from this point down, so there's no need to recompute.
		memoKey := key + strconv.FormatBool(dacSeen) + strconv.FormatBool(fftSeen)
		if val, ok := memo[memoKey]; ok {
			return val
		}

		dacSeen = dacSeen || key == "dac"
		fftSeen = fftSeen || key == "fft"

		total := 0
		for _, otherKey := range deviceMap[key] {
			total += walkDag(otherKey, dacSeen, fftSeen)
		}

		memo[memoKey] = total

		return total
	}

	return walkDag("svr", false, false)
}
