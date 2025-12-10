package main

import (
	"fmt"
	"image"
	"log"
	"time"

	"aoc/internal/utils"
)

const (
	part1Expected = 50
	part2Expected = 24
)

func main() {
	runTest()
	run()
}

func runTest() {
	part1 := solvePart1("day_09_test.txt")
	part2 := solvePart2("day_09_test.txt")
	if part1 != part1Expected {
		log.Fatalf("Part 1 - Expected: %d, receieved: %d", part1Expected, part1)
	}
	if part2 != part2Expected {
		log.Fatalf("Part 2 - Expected: %d, receieved: %d", part2Expected, part2)
	}
}

func run() {
	start := time.Now()
	part1 := solvePart1("day_09.txt")
	fmt.Println("Part 1 duration:", time.Since(start))
	start = time.Now()
	part2 := solvePart2("day_09.txt")
	fmt.Println("Part 2 duration:", time.Since(start))
	fmt.Println("Part 1 - Answer is:", int(part1))
	fmt.Println("Part 2 - Answer is:", int(part2))
}

func solvePart1(fileName string) int {
	rects, _ := parseInput(fileName)

	biggestArea := 0
	for _, rect := range rects {
		rect.Max = rect.Max.Add(image.Point{1, 1})
		rectArea := rect.Dx() * rect.Dy()
		biggestArea = max(biggestArea, rectArea)
	}

	return biggestArea
}

func solvePart2(fileName string) int {
	rects, lines := parseInput(fileName)

	biggestArea := 0
loop:
	for _, rect := range rects {
		rect.Max = rect.Max.Add(image.Point{1, 1})
		rectArea := rect.Dx() * rect.Dy()

		for _, line := range lines {
			line.Max = line.Max.Add(image.Point{1, 1})
			if line.Overlaps(rect.Inset(1)) {
				continue loop
			}
		}

		biggestArea = max(biggestArea, rectArea)
	}

	return biggestArea
}

func parseInput(fileName string) ([]image.Rectangle, []image.Rectangle) {
	scanner := utils.GetInputFileLineScanner(fileName)

	points, rects, lines := []image.Point{}, []image.Rectangle{}, []image.Rectangle{}
	for scanner.Scan() {
		line := scanner.Text()

		var point image.Point
		_, err := fmt.Sscanf(line, "%d,%d", &point.X, &point.Y)
		if err != nil {
			log.Fatal("Error parsing point line", line, err)
		}

		// Create a new rectangle with this point and every past point
		for _, existingPoint := range points {
			rects = append(rects, image.Rectangle{point, existingPoint}.Canon())
		}

		points = append(points, point)

		// There's a previous point we can use to form a line
		if len(points) > 1 {
			lines = append(lines, image.Rectangle{points[len(points)-2], point}.Canon())
		}

	}

	// Connect the last point back to the starting point to complete the polygon lines
	lines = append(lines, image.Rectangle{points[len(points)-1], points[0]}.Canon())

	return rects, lines
}
