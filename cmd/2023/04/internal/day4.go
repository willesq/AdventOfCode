package internal

import (
	"adventOfCode/internal/adventhelper"
	"fmt"
	"strconv"
	"strings"
)

func Part1(filename string) (*Challenge, *int) {
	derp := adventhelper.ReadFile(fmt.Sprintf(filename))
	Input := Challenge{}
	for _, line := range *derp {
		tmp_card := ScratchCard{RawString: line}
		tmp_card.init()
		Input.Cards = append(Input.Cards, tmp_card)
	}
	totalSum := 0
	for _, c := range Input.Cards {
		totalSum += c.CardScore
	}
	return &Input, &totalSum
}

func Part2(input *Challenge) *int {
	cardMap := map[int][]ScratchCard{}
	for cidx := 0; cidx < len(input.Cards); cidx++ {
		adjustment := cidx + 1
		cardMap[adjustment] = append(cardMap[adjustment], input.Cards[cidx])
	}
	cardPile := 0
	for mapKey := 1; mapKey <= len(input.Cards); mapKey++ {
		for _, c := range cardMap[mapKey] {
			cardWin := len(c.MatchingNumbers)
			for i := 1; i <= cardWin; i++ {
				addIdx := c.CardNumber + i
				if addIdx >= len(input.Cards) {
					continue
				}
				cardMap[addIdx] = append(cardMap[addIdx], input.Cards[addIdx-1])
			}
		}
		cardPile += len(cardMap[mapKey])
		fmt.Printf("Key: %d   \t arrayLen:%d     \t MatchingNums: %d   \t CardPile: %d\n", mapKey, len(cardMap[mapKey]), len(input.Cards[mapKey-1].MatchingNumbers), cardPile)
	}
	scratchCardPile := CountCards(&cardMap, 0, 0)
	return &scratchCardPile
}

type Challenge struct {
	Cards []ScratchCard
}

type ScratchCard struct {
	RawString       string
	CardNumber      int
	WinningNumbers  []int
	CardNumbers     []int
	MatchingNumbers []int
	CardScore       int
}

func (s *ScratchCard) init() {
	splitGame := strings.Split(s.RawString, ": ")
	s.CardNumber, _ = strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(splitGame[0], "Card ")))
	numSplit := strings.Split(splitGame[1], " | ")
	for _, v := range strings.Split(numSplit[0], " ") {
		vNum, _ := strconv.Atoi(strings.TrimSpace(v))
		// Numbers to match to, in order to win
		s.WinningNumbers = append(s.WinningNumbers, vNum)
	}
	for _, v := range strings.Split(numSplit[1], " ") {
		// Card numbers to perform match on
		if v == "" {
			continue
		}
		vNum, _ := strconv.Atoi(strings.TrimSpace(v))
		s.CardNumbers = append(s.CardNumbers, vNum)
	}
	for _, cN := range s.CardNumbers {
		for _, wN := range s.WinningNumbers {
			if cN == wN {
				s.MatchingNumbers = append(s.MatchingNumbers, cN)
				break
			}
		}
	}
	if len(s.MatchingNumbers) > 0 {
		for i := 0; i < len(s.MatchingNumbers); i++ {
			if s.CardScore == 0 {
				s.CardScore = 1
				continue
			}
			s.CardScore *= 2
		}
	}
}

func CountCards(cardArray *map[int][]ScratchCard, startval int, idx int) int {
	if idx == len(*cardArray) {
		return len((*cardArray)[idx])
	}
	value := len((*cardArray)[idx])
	return value + CountCards(cardArray, value, idx+1)
}
