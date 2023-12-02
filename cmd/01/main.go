package _1

import (
	"adventOfCode2023/internal/adventhelper"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func getCurrentFilePath() (string, error) {
	// Get the absolute path of the currently executing Go executable
	exeFilePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	// Get the directory containing the executable
	exeDir := filepath.Dir(exeFilePath)

	// Join the executable directory with the name of the currently running Go file
	// This assumes that the Go file has the same name as the executable (without the .exe extension on Windows)
	currentFileName := filepath.Base(exeFilePath)
	currentFilePath := filepath.Join(exeDir, currentFileName)

	return currentFilePath, nil
}

func currentDirectory() *string {
	currentFilePath, err := getCurrentFilePath()
	if err != nil {
		fmt.Println("Error:", err)
		return &currentFilePath
	}
	return &currentFilePath
}

func Part1(filename string) *int {
	derp := adventhelper.ReadFile(fmt.Sprintf(filename))
	return evalintstringslice(*derp)
}

func Part2(filename string) *int {
	derp := adventhelper.ReadFile(fmt.Sprintf(filename))
	derpUpdated := []string{}
	replacements := map[string]int{
		"nine":  9,
		"eight": 8,
		"seven": 7,
		"six":   6,
		"five":  5,
		"four":  4,
		"three": 3,
		"two":   2,
		"one":   1,
	}

	for _, line := range *derp {
		updatedLine := line

		var result strings.Builder
		i := 0
		for i < len(line) {
			found := false

			// Try to match spelled-out numbers from the current position
			for _, num := range keys(replacements) {
				if strings.HasPrefix(line[i:], num) {
					// Found a match, append the numeric value
					result.WriteString(fmt.Sprintf("%d", replacements[num]))
					// Move the position after the matched word
					i += len(num)
					found = true
					break
				}
			}

			// If no match is found, append the current character and move to the next
			if !found {
				result.WriteByte(line[i])
				i++
			}
		}
		updatedLine = result.String()
		derpUpdated = append(derpUpdated, updatedLine)
		fmt.Print(line, " -> ", updatedLine, "\n")
	}
	return evalintstringslice(derpUpdated)
}

func evalintstringslice(data []string) *int {
	pattern := `[0-9]+`
	re := regexp.MustCompile(pattern)
	intSlice := []int{}

	for _, line := range data {
		LineMatches := re.FindAllString(line, -1)
		if len(LineMatches) == 1 && len(LineMatches[0]) == 2 {
			intSlice = append(intSlice, *convertToInt(LineMatches[0]))
			continue
		}
		linenumes := ""
		for _, match := range LineMatches {
			linenumes += match
		}
		lineintstr := *getFirstChar(linenumes) + *getLastChar(linenumes)
		lineInt := convertToInt(lineintstr)
		intSlice = append(intSlice, *lineInt)
	}
	finalSum := 0

	for _, v := range intSlice {
		finalSum += v
	}
	return &finalSum
}

func getFirstChar(str string) *string {
	if str == "" {
		return nil // Return nil for an empty string
	}
	firstChar := string(str[0])
	return &firstChar
}

func getLastChar(str string) *string {
	if str == "" {
		return nil // Return nil for an empty string
	}
	lastChar := string(str[len(str)-1])
	return &lastChar
}

func convertToInt(str string) *int {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Print(err)
		return &num
	}
	return &num
}

func convertToStr(value int) *string {
	numstring := strconv.Itoa(value)
	return &numstring
}

func keys(m map[string]int) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
