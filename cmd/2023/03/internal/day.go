package internal

import (
	"adventOfCode/internal/adventhelper"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Part1(filename string) (*Challenge, *int) {
	derp := adventhelper.ReadFile(fmt.Sprintf(filename))
	var Input Challenge
	for _, line := range *derp {
		tmp_line := Line{RawString: line}
		tmp_line.init()
		Input.Lines = append(Input.Lines, tmp_line)
	}
	Input.uniqueSymbols()
	Input.GetPartNumbers()

	totalSum := 0
	for _, v := range Input.PartNumbers {
		totalSum += v
	}

	return &Input, &totalSum
}

func Part2(input *Challenge) *int {
	input.GetGears()
	totalCount := 0

	for _, val := range input.Gears {
		valRatio := val[0] * val[1]
		totalCount += valRatio
	}

	return &totalCount
}

type Challenge struct {
	Lines         []Line
	UniqueSymbols []string
	PartNumbers   map[string]int
	Gears         map[string][]int
}

func (c *Challenge) uniqueSymbols() {
	for _, line := range c.Lines {
		for _, l := range line.SymbolLocations {
			if !slices.Contains(c.UniqueSymbols, l) {
				c.UniqueSymbols = append(c.UniqueSymbols, l)
			}
		}
	}
}

func (c *Challenge) GetPartNumbers() {
	if c.PartNumbers == nil {
		c.PartNumbers = map[string]int{}
	}
	for ridx, row := range c.Lines {
		// get current line symbols
		c.evalLines(row, ridx, row.SymbolLocations, &c.PartNumbers)
		if ridx != 0 {
			// Previous Line
			prevLineIndex := ridx - 1
			c.evalLines(c.Lines[prevLineIndex], prevLineIndex, row.SymbolLocations, &c.PartNumbers)
		}
		if ridx != len(c.Lines)-1 {
			// Next line
			nextLineIndex := ridx + 1
			c.evalLines(c.Lines[nextLineIndex], nextLineIndex, row.SymbolLocations, &c.PartNumbers)
		}

	}
}

func (c *Challenge) GetGears() {
	if c.Gears == nil {
		c.Gears = map[string][]int{}
	}
	for ridx, row := range c.Lines {
		// get current line symbols
		surroundingNumbers := map[string]int{}
		c.evalLines(row, ridx, row.SymbolLocations, &surroundingNumbers, true)
		if ridx != 0 {
			// Previous Line
			prevLineIndex := ridx - 1
			c.evalLines(c.Lines[prevLineIndex], prevLineIndex, row.SymbolLocations, &surroundingNumbers, true)
		}
		if ridx != len(c.Lines)-1 {
			// Next line
			nextLineIndex := ridx + 1
			c.evalLines(c.Lines[nextLineIndex], nextLineIndex, row.SymbolLocations, &surroundingNumbers, true)
		}
		if len(surroundingNumbers) > 0 {
			for symKey, symValue := range row.SymbolLocations {
				numberAroundSym := []int{}
				if symValue != "*" {
					continue
				}
				for numKey, numVal := range surroundingNumbers {
					keySplit := strings.Split(numKey, "-")
					keyValIndex, _ := strconv.Atoi(keySplit[1])
					keyValString := strconv.Itoa(numVal)
					for i := 0; i < len(keyValString); i++ {
						if keyValIndex+i == symKey || keyValIndex+i == symKey-1 {
							numberAroundSym = append(numberAroundSym, numVal)
							break
						}
					}
					if keyValIndex == symKey+1 {
						numberAroundSym = append(numberAroundSym, numVal)
					}
				}
				if len(numberAroundSym) == 2 {
					key_val := fmt.Sprintf("%d-%d", ridx, symKey)
					c.Gears[key_val] = numberAroundSym
				}
			}
		}
	}
}

func (c *Challenge) evalLines(row Line, rowidx int, symbols map[int]string, inventory *map[string]int, gear ...bool) {
	for symkey, symValue := range symbols {
		// iterate over the symbols in the lines
		for nkey, nvalue := range row.NumberLocations {
			// check current line
			if gear != nil && "*" != symValue {
				continue
			}

			extractNumber := false
			nvalueString := strconv.Itoa(nvalue)
			nvalueStr := len(nvalueString)
			if nkey == symkey+1 {
				// Number is to the right of key
				extractNumber = true
			}
			if !extractNumber {
				for i := 0; i <= nvalueStr; i++ {
					if nkey+i == symkey {
						extractNumber = true
						break
					}
				}
			}
			if extractNumber {
				key_val := fmt.Sprintf("%d-%d", rowidx, nkey)
				if _, ok := (*inventory)[key_val]; !ok {
					// Only add to map if key doesn't exist
					(*inventory)[key_val] = nvalue
				}
			}
		}
	}
}

type Line struct {
	RawString       string
	RowSlice        []string
	SymbolLocations map[int]string
	NumberLocations map[int]int
	numbers         []string
}

func (l *Line) init() {
	l.numbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	l.RowSlice = []string{}
	l.SymbolLocations = map[int]string{}
	l.NumberLocations = map[int]int{}
	prev := ""
	numberstr := ""
	numberidx := -1
	for idx, str := range l.RawString {
		str_char := string(str)
		l.RowSlice = append(l.RowSlice, str_char)
		cont := false
		if str_char == "." {
			// continue if a dot
			cont = true
		}
		if !slices.Contains(l.numbers, str_char) && !cont {
			// add symbols indices
			l.SymbolLocations[idx] = str_char
			cont = true
		} else if slices.Contains(l.numbers, str_char) {
			// add number indices
			if !slices.Contains(l.numbers, prev) && numberidx == -1 {
				numberidx = idx
			}
			numberstr += str_char
		}

		rsLen := len(l.RawString) - 1
		if idx == rsLen || cont {
			if numberidx != -1 && numberstr != "" {
				tmp_num, _ := strconv.Atoi(numberstr)
				l.NumberLocations[numberidx] = tmp_num
				numberidx = -1
				numberstr = ""
			}
		}
		prev = str_char
	}
}
