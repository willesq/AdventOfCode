package internal

import (
	"adventOfCode/internal/adventhelper"
	"fmt"
	"strconv"
	"strings"
)

func Part1(filename string) (*ChallengeInput, *int) {
	derp := adventhelper.ReadFile(fmt.Sprintf(filename))
	Input := ChallengeInput{PossibleContraints: Handful{
		Red:   12,
		Green: 13,
		Blue:  14,
	}}
	for _, line := range *derp {
		gr := GameRecord{Constraints: Input.PossibleContraints}
		gr.Init(line)
		Input.Games = append(Input.Games, gr)
		if gr.Valid {
			Input.PossibleGameIDs = append(Input.PossibleGameIDs, gr.GameNumber)
		}
	}
	totalSum := 0
	for _, n := range Input.PossibleGameIDs {
		totalSum += n
	}
	return &Input, &totalSum
}

func Part2(input *ChallengeInput) *int {
	v := 0
	for _, game := range input.Games {
		v += game.Power
	}
	return &v
}

type ChallengeInput struct {
	Games              []GameRecord
	PossibleGameIDs    []int
	IDSum              int
	PossibleContraints Handful
}

type GameRecord struct {
	GameNumber  int
	Hands       []Handful
	Constraints Handful
	possibleCnt int
	Valid       bool
	MinimumReq  Handful
	Power       int
}

func (r *GameRecord) Init(line string) {
	split := strings.Split(line, ": ")
	r.GameNumber, _ = strconv.Atoi(strings.TrimPrefix(split[0], "Game "))
	for _, hand := range strings.Split(split[1], "; ") {
		h := Handful{Contraints: &r.Constraints}
		h.Init(hand)
		r.Hands = append(r.Hands, h)
	}
	for _, h := range r.Hands {
		if h.Possible {
			r.possibleCnt++
		}
	}
	if r.possibleCnt == len(r.Hands) {
		r.Valid = true
	}
	r.evalMinimum()
}

func (r *GameRecord) evalMinimum() {
	for _, h := range r.Hands {
		if h.Red > r.MinimumReq.Red {
			r.MinimumReq.Red = h.Red
		}
		if h.Green > r.MinimumReq.Green {
			r.MinimumReq.Green = h.Green
		}
		if h.Blue > r.MinimumReq.Blue {
			r.MinimumReq.Blue = h.Blue
		}
	}
	r.Power = r.MinimumReq.Red * r.MinimumReq.Blue * r.MinimumReq.Green
}

type Handful struct {
	Red        int
	Green      int
	Blue       int
	Contraints *Handful
	Possible   bool
}

func (h *Handful) Init(line string) {
	split := strings.Split(line, ", ")
	for _, part := range split {
		if strings.Contains(part, "red") {
			h.Red, _ = strconv.Atoi(strings.TrimSuffix(part, " red"))
		} else if strings.Contains(part, "green") {
			h.Green, _ = strconv.Atoi(strings.TrimSuffix(part, " green"))
		} else if strings.Contains(part, "blue") {
			h.Blue, _ = strconv.Atoi(strings.TrimSuffix(part, " blue"))
		}
	}
	if h.Red <= h.Contraints.Red &&
		h.Blue <= h.Contraints.Blue &&
		h.Green <= h.Contraints.Green {
		h.Possible = true
	} else {
		h.Possible = false
	}
}
