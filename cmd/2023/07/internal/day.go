package internal

import (
	"adventOfCode/internal/adventhelper"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type HandStrength int

const (
	HighCard HandStrength = iota
	OnePair
	TwoPair
	ThreeOfKind
	FullHouse
	FourOfKind
	FiveOfKind
)

var Cards = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

//var weight = make(map[string]int)
//weight := map[string]int{}

func Part1(filename string) (*Challenge, *int) {
	derp := adventhelper.ReadFile(fmt.Sprintf(filename))
	Input := Challenge{RawData: derp}
	Input.init()
	Input.RankCards()

	for key, cm := range Input.Rankings {
		Input.Rankings[key] = MergeSortHands(cm)
	}
	iteration := 1
	totalWinnings := 0
	keyRanges := []HandStrength{HighCard, OnePair, TwoPair, ThreeOfKind, FullHouse, FourOfKind, FiveOfKind}
	for _, kr := range keyRanges {
		fmt.Printf("Key: %v\n", kr)
		for _, val := range Input.Rankings[kr] {
			winning := val.Bid * iteration
			totalWinnings += winning
			iteration++
		}
	}
	return &Input, &totalWinnings
}

func Part2(Input *Challenge) *int {
	lowest := -1
	return &lowest
}

type Challenge struct {
	RawData  *[]string
	Hands    []*Hand
	Rankings map[HandStrength][]*Hand
}

func (c *Challenge) init() {
	for _, line := range *c.RawData {
		lineSplit := strings.Split(line, " ")
		bid, _ := strconv.Atoi(lineSplit[1])
		hand := Hand{Bid: bid}
		for _, char := range lineSplit[0] {
			hand.Cards = append(hand.Cards, string(char))
		}
		hand.findStrength()
		c.Hands = append(c.Hands, &hand)
	}
}

func (c *Challenge) RankCards() {
	c.Rankings = map[HandStrength][]*Hand{}
	for _, hand := range c.Hands {
		if _, exists := c.Rankings[hand.Strength]; exists {
			c.Rankings[hand.Strength] = append(c.Rankings[hand.Strength], hand)
		} else {
			c.Rankings[hand.Strength] = []*Hand{hand}
		}
	}
}

type Hand struct {
	Cards    []string
	Bid      int
	Strength HandStrength
	CardMap  map[string]int
}

func (h *Hand) findStrength() {
	// make eval map
	h.CardMap = map[string]int{}
	for _, card := range h.Cards {
		if _, exists := h.CardMap[card]; exists {
			h.CardMap[card]++
		} else {
			h.CardMap[card] = 1
		}
	}
	Three := false
	Pair1 := false
	Pair2 := false
	distinctCnt := 0
	for _, val := range h.CardMap {
		if val == 5 {
			h.Strength = FiveOfKind
			return
		}
	}
	for _, val := range h.CardMap {
		if val == 4 {
			h.Strength = FourOfKind
			return
		}
	}
	for _, val := range h.CardMap {
		customeTypeVal := HandStrength(val)
		if ThreeOfKind == customeTypeVal {
			Three = true
			if Three && Pair1 {
				h.Strength = FullHouse
				return
			}
		}
		if TwoPair == customeTypeVal {
			if !Pair1 {
				Pair1 = true
			} else {
				Pair2 = true
			}
			if Three && Pair1 {
				h.Strength = FullHouse
				return
			}
			if Pair1 && Pair2 {
				h.Strength = TwoPair
				return
			}
		}
		if val == 1 {
			distinctCnt++
		}
	}
	if Three && !Pair1 {
		h.Strength = ThreeOfKind
		return
	}
	if Pair1 && !Pair2 {
		h.Strength = OnePair
		return
	}
	if distinctCnt == 5 {
		h.Strength = HighCard
		return
	}
}

func MergeSortHands(hands []*Hand) []*Hand {
	//result := []*Hand{}
	if len(hands) <= 1 {
		return hands
	}

	left := []*Hand{}
	right := []*Hand{}

	for i := 0; i < len(hands); i++ {
		if i < (len(hands) / 2) {
			left = append(left, hands[i])
		} else {
			right = append(right, hands[i])
		}
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		left = MergeSortHands(left)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		right = MergeSortHands(right)
	}()
	wg.Wait()
	//left = MergeSortHands(left)
	//right = MergeSortHands(right)
	res := MergeHands(left, right)
	return res
}

func MergeHands(left, right []*Hand) []*Hand {
	result := []*Hand{}

	for {
		if len(left) != 0 && len(right) != 0 {
			lHand := left[0]
			rHand := right[0]
			cardsEqual := 0
			for i := 0; i < 5; i++ {
				if lHand.Cards[i] != rHand.Cards[i] {
					leftCardWeight := Cards[lHand.Cards[i]]
					rightCardWeight := Cards[rHand.Cards[i]]
					if leftCardWeight < rightCardWeight {
						result = append(result, lHand)
						if len(left) == 1 {
							left = []*Hand{}
						} else {
							left = left[1:]
						}
					} else {
						result = append(result, rHand)
						if len(right) == 1 {
							right = []*Hand{}
						} else {
							right = right[1:]
						}
					}
					break
				} else {
					cardsEqual++
					continue
				}
			}
			if cardsEqual == 5 {
				break
			}
		} else {
			break
		}
	}

	for _, item := range left {
		result = append(result, item)
	}
	for _, item := range right {
		result = append(result, item)
	}

	return result
}
