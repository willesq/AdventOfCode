package adventhelper

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) *[]string {
	lines := []string{}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return &lines
	}
	defer file.Close() // Make sure to close the file when done

	// Create a scanner to read from the file
	scanner := bufio.NewScanner(file)

	// Iterate through the file line by line
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
	return &lines
}
