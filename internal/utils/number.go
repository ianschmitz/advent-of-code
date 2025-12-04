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
