package main

import (
	"fmt"
	"log"
	"regexp"
	"slices"
	"sort"
	"strings"
	"time"

	"aoc/internal/utils"
)

const (
	part1Expected = 7
	part2Expected = 0
)

func main() {
	runTest()
	run()
}

func runTest() {
	part1 := solvePart1("day_10_test.txt")
	part2 := solvePart2("day_10_test.txt")
	if part1 != part1Expected {
		log.Fatalf("Part 1 - Expected: %d, receieved: %d", part1Expected, part1)
	}
	if part2 != part2Expected {
		log.Fatalf("Part 2 - Expected: %d, receieved: %d", part2Expected, part2)
	}
}

func run() {
	start := time.Now()
	part1 := solvePart1("day_10.txt")
	fmt.Println("Part 1 duration:", time.Since(start))
	start = time.Now()
	part2 := solvePart2("day_10.txt")
	fmt.Println("Part 2 duration:", time.Since(start))
	fmt.Println("Part 1 - Answer is:", part1)
	fmt.Println("Part 2 - Answer is:", part2)
}

func solvePart1(fileName string) int {
	scanner := utils.GetInputFileLineScanner(fileName)

	machines := []machine{}
	for scanner.Scan() {
		line := scanner.Text()

		machines = append(machines, parseMachine(line))
	}

	machinePresses := []int{}
	for _, machine := range machines {
		machinePresses = append(machinePresses, getNumberOfPresses(machine))
	}

	sum := 0
	for _, val := range machinePresses {
		sum += val
	}
	return sum
}

func solvePart2(fileName string) int {
	return 0
}

type machine struct {
	lightDiagram        int
	wiringSchematics    []int
	joltageRequirements []int
}

func parseMachine(machineStr string) machine {
	diagramRe := regexp.MustCompile(`\[([.#]+)\]`)
	diagramStr := diagramRe.FindString(machineStr)
	diagramStr = diagramStr[1 : len(diagramStr)-1]

	lightDiagram := 0

	chars := strings.Split(diagramStr, "")
	for i, char := range chars {
		switch char {
		case "#":
			lightDiagram |= (1 << i)
		}
	}

	schematicsRe := regexp.MustCompile(`\(((?:\d+,?)+)\)`)
	schematicsMatches := schematicsRe.FindAllStringSubmatch(machineStr, -1)
	wiringSchematics := []int{}
	for _, match := range schematicsMatches {
		schematicVal := 0
		for val := range strings.SplitSeq(match[1], ",") {
			schematicVal |= (1 << utils.StringToInt(val))
		}

		wiringSchematics = append(wiringSchematics, schematicVal)
	}

	joltageRe := regexp.MustCompile(`\{((?:\d,?)+)\}`)
	joltageMatches := joltageRe.FindAllStringSubmatch(machineStr, -1)
	joltageRequirements := []int{}
	for match := range strings.SplitSeq(joltageMatches[0][1], ",") {
		joltageRequirements = append(joltageRequirements, utils.StringToInt(match))
	}

	return machine{lightDiagram, wiringSchematics, joltageRequirements}
}

// getCombinations generates all unique combinations of length n from a given slice.
func getCombinations(elements []int, n int) [][]int {
	var result [][]int
	if n < 0 || n > len(elements) {
		return result // Invalid n or not enough elements
	}

	// Sort the input to handle duplicates and ensure unique combinations
	// (e.g., {1,2} and {2,1} are considered the same combination)
	sort.Ints(elements)

	// Recursive helper function
	var generate func(start int, currentCombination []int)
	generate = func(start int, currentCombination []int) {
		if len(currentCombination) == n {
			// Found a combination of the desired length
			// Create a copy to avoid modification issues
			temp := make([]int, n)
			copy(temp, currentCombination)
			result = append(result, temp)
			return
		}

		for i := start; i < len(elements); i++ {
			// Skip duplicate elements to ensure unique combinations
			if i > start && elements[i] == elements[i-1] {
				continue
			}

			// Include the current element
			currentCombination = append(currentCombination, elements[i])
			// Recurse with the next element and updated combination
			generate(i+1, currentCombination)
			// Backtrack: remove the current element for other combinations
			currentCombination = currentCombination[:len(currentCombination)-1]
		}
	}

	generate(0, []int{})
	return result
}

func getNumberOfPresses(machine machine) int {
	// Can do in one press, there's already a combo that meets our light diagram
	if slices.Contains(machine.wiringSchematics, machine.lightDiagram) {
		return 1
	}

	// Gather all N combinations of presses starting from the smallest possible number of presses
	for i := 2; i < len(machine.wiringSchematics); i++ {
		combinations := getCombinations(machine.wiringSchematics, i)

		for _, combination := range combinations {
			result := 0

			for _, num := range combination {
				result = result ^ num
				if result == machine.lightDiagram {
					return i
				}
			}
		}
	}

	log.Fatal("Couldn't find number of presses needed", machine)
	return -1
}
