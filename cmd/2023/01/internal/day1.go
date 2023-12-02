package internal

import (
	"adventOfCode2023/internal/adventhelper"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Part1(filename string) *int {
	derp := adventhelper.ReadFile(fmt.Sprintf(filename))
	return evalintstringslice(*derp)
}

func Part2(filename string) *int {
	derp := adventhelper.ReadFile(fmt.Sprintf(filename))
	derpUpdated := []string{}

	for _, line := range *derp {
		updatedLine := *getIntString(line) + *convertStrWordToInt(line)
		derpUpdated = append(derpUpdated, updatedLine)
	}
	return evalintstringslice(derpUpdated)
}

func evalintstringslice(data []string) *int {
	intSlice := []int{}
	for _, line := range data {
		linenumes := getIntString(line)
		if *linenumes == "" {
			linenumes = convertStrWordToInt(line)
		}
		lineintstr := *getFirstChar(*linenumes) + *getLastChar(*linenumes)
		lineInt := convertToInt(lineintstr)
		intSlice = append(intSlice, *lineInt)
	}
	finalSum := 0

	for _, v := range intSlice {
		finalSum += v
	}
	return &finalSum
}

func getIntString(line string) *string {
	pattern := `[0-9]+`
	re := regexp.MustCompile(pattern)
	LineMatches := re.FindAllString(line, -1)
	linenumes := ""
	for _, match := range LineMatches {
		linenumes += match
	}
	return &linenumes
}

func convertStrWordToInt(str string) *string {
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
	var result strings.Builder
	i := 0
	for i < len(str) {
		found := false

		// Try to match spelled-out numbers from the current position
		for _, num := range keys(replacements) {
			if strings.HasPrefix(str[i:], num) {
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
			//result.WriteByte(str[i])
			i++
		}
	}
	res := result.String()
	return &res
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

func keys(m map[string]int) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
