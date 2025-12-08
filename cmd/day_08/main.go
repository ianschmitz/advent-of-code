package main

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strings"
	"time"

	"aoc/internal/utils"
)

const (
	part1Expected = 40
	part2Expected = 0
)

func main() {
	runTest()
	run()
}

func runTest() {
	part1 := solvePart1("day_08_test.txt", 10)
	part2 := solvePart2("day_08_test.txt")
	if part1 != part1Expected {
		log.Fatalf("Part 1 - Expected: %d, receieved: %d", part1Expected, part1)
	}
	if part2 != part2Expected {
		log.Fatalf("Part 2 - Expected: %d, receieved: %d", part2Expected, part2)
	}
}

func run() {
	start := time.Now()
	part1 := solvePart1("day_08.txt", 1000)
	fmt.Println("Part 1 duration:", time.Since(start))
	start = time.Now()
	part2 := solvePart2("day_08.txt")
	fmt.Println("Part 2 duration:", time.Since(start))
	fmt.Println("Part 1 - Answer is:", part1)
	fmt.Println("Part 2 - Answer is:", part2)
}

func solvePart1(fileName string, pairsToConnect int) int {
	scanner := utils.GetInputFileLineScanner(fileName)

	points := []point3D{}
	for scanner.Scan() {
		line := scanner.Text()

		strCoords := strings.Split(line, ",")
		coords := point3D{
			X: utils.StringToInt(strCoords[0]),
			Y: utils.StringToInt(strCoords[1]),
			Z: utils.StringToInt(strCoords[2]),
		}

		points = append(points, coords)
	}

	pointPairs := getAllCombosWithDistance(points)
	slices.SortFunc(pointPairs, func(a distancePair, b distancePair) int {
		return int(a.Distance - b.Distance)
	})

	circuits := [][]point3D{}

	connectedCount := 0
	for _, pointPair := range pointPairs {
		if connectedCount >= pairsToConnect {
			break
		}

		foundCircuitIdx := slices.IndexFunc(circuits, func(circuit []point3D) bool {
			return slices.Contains(circuit, pointPair.P1) || slices.Contains(circuit, pointPair.P2)
		})

		if foundCircuitIdx > -1 {
			containsP1 := slices.Contains(circuits[foundCircuitIdx], pointPair.P1)
			containsP2 := slices.Contains(circuits[foundCircuitIdx], pointPair.P2)

			if containsP1 && !containsP2 {
				circuits[foundCircuitIdx] = append(circuits[foundCircuitIdx], pointPair.P2)
				connectedCount++
			} else if !containsP1 {
				circuits[foundCircuitIdx] = append(circuits[foundCircuitIdx], pointPair.P1)
				connectedCount++
			}
		} else {
			circuits = append(circuits, []point3D{pointPair.P1, pointPair.P2})
			connectedCount++
		}
	}

	slices.SortFunc(circuits, func(a, b []point3D) int {
		return len(b) - len(a)
	})

	for _, val := range circuits {
		fmt.Println(val)
	}

	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}

func solvePart2(fileName string) int {
	return 0
}

func calculateDistance(p1, p2 point3D) float64 {
	// Use the distance formula: d = sqrt((x2-x1)² + (y2-y1)² + (z2-z1)²)
	dx := float64(p2.X - p1.X)
	dy := float64(p2.Y - p1.Y)
	dz := float64(p2.Z - p1.Z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

type point3D struct {
	X, Y, Z int
}

type distancePair struct {
	P1, P2   point3D
	Distance float64
}

// getAllCombosWithDistance generates the Cartesian product of a single slice
// without creating a pair of the same element, or the reverse of the same pair
func getAllCombosWithDistance(slice []point3D) []distancePair {
	var result []distancePair

	// Outer loop iterates from the first element to the second-to-last
	for i := range slice {
		// Inner loop iterates from the element *after* the outer loop's current element to the end
		for j := i + 1; j < len(slice); j++ {
			// This logic ensures:
			// 1. We don't pair an element with itself (j starts from i + 1)
			// 2. We don't create reverse duplicates (e.g., if we have [1, 2], we won't get [2, 1] because for i=2, j starts from 3)

			pair := distancePair{
				P1:       slice[i],
				P2:       slice[j],
				Distance: calculateDistance(slice[i], slice[j]),
			}

			result = append(result, pair)
		}
	}

	return result
}
