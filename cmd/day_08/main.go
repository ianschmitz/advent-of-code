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
	part2Expected = 25272
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
	slices.SortFunc(pointPairs, func(a, b distancePair) int {
		return int(a.Distance - b.Distance)
	})
	pointPairs = pointPairs[:pairsToConnect]

	circuits := [][]point3D{}
	for _, point := range points {
		circuits = append(circuits, []point3D{point})
	}

	for _, pointPair := range pointPairs {
		p1Idx := slices.IndexFunc(circuits, func(circuit []point3D) bool {
			return slices.ContainsFunc(circuit, func(point point3D) bool {
				return point == pointPair.P1
			})
		})
		p2Idx := slices.IndexFunc(circuits, func(circuit []point3D) bool {
			return slices.ContainsFunc(circuit, func(point point3D) bool {
				return point == pointPair.P2
			})
		})

		circuits = mergeCircuits(circuits, p1Idx, p2Idx)
	}

	// Sort based on largest circuits
	slices.SortFunc(circuits, func(a, b []point3D) int {
		return len(b) - len(a)
	})

	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}

func solvePart2(fileName string) int {
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
	slices.SortFunc(pointPairs, func(a, b distancePair) int {
		return int(a.Distance - b.Distance)
	})

	circuits := [][]point3D{}
	for _, point := range points {
		circuits = append(circuits, []point3D{point})
	}

	for _, pointPair := range pointPairs {
		p1Idx := slices.IndexFunc(circuits, func(circuit []point3D) bool {
			return slices.ContainsFunc(circuit, func(point point3D) bool {
				return point == pointPair.P1
			})
		})
		p2Idx := slices.IndexFunc(circuits, func(circuit []point3D) bool {
			return slices.ContainsFunc(circuit, func(point point3D) bool {
				return point == pointPair.P2
			})
		})

		circuits = mergeCircuits(circuits, p1Idx, p2Idx)

		if len(circuits) == 1 {
			return pointPair.P1.X * pointPair.P2.X
		}
	}

	log.Fatal("Didn't form one circuit")
	return 0
}

func calculateDistance(p1, p2 point3D) float64 {
	// Use the distance formula: d = sqrt((x2-x1)² + (y2-y1)² + (z2-z1)²)
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	dz := p2.Z - p1.Z
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
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

	for i := range slice {
		// This logic ensures:
		// 1. We don't pair an element with itself
		// 2. We don't create reverse duplicates
		for j := i + 1; j < len(slice); j++ {
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

func mergeCircuits(circuits [][]point3D, idx1 int, idx2 int) [][]point3D {
	if idx1 == idx2 {
		return circuits
	}

	circuits[idx1] = append(circuits[idx1], circuits[idx2]...)
	return append(circuits[:idx2], circuits[idx2+1:]...)
}
