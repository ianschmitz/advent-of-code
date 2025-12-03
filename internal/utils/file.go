package utils

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

// GetInputFileLineScanner is less efficient than using a bufio scanner on a [os.File] object.
// I'm reading it straight into memory to simplify usage in solution files.
func GetInputFileLineScanner(fileName string) *bufio.Scanner {
	fileData := ReadInputFile(fileName)

	scanner := bufio.NewScanner(bytes.NewReader(fileData))
	scanner.Split(bufio.ScanLines)

	return scanner
}

func ReadInputFile(fileName string) []byte {
	file, err := os.ReadFile("input/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
