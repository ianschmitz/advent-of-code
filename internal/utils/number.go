package utils

import (
	"log"
	"strconv"
)

func StringToInt(stringNum string) int {
	num, err := strconv.Atoi(stringNum)
	if err != nil {
		log.Fatal("Error converting string to int:", err)
	}
	return num
}

func StringToFloat(stringNum string) float64 {
	num, err := strconv.ParseFloat(stringNum, 64)
	if err != nil {
		log.Fatal("Error converting string to float:", err)
	}
	return num
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
