package internal

import (
	"adventOfCode/internal/adventhelper"
	"fmt"
	"strings"
)

func Part1(filename string) (*Challenge, *int) {
	derp := adventhelper.ReadFile(fmt.Sprintf(filename))
	Input := Challenge{RawData: derp}
	Input.init()
	derpInt := 0
	iter := 0
	startKey := "AAA"
	for {
		evalKey := Input.Maps[startKey]
		if evalKey.Value == "ZZZ" {
			break
		}
		leftRight := string(Input.LRInstruction[iter])
		if leftRight == "L" {
			startKey = evalKey.Left
		} else if leftRight == "R" {
			startKey = evalKey.Right
		}

		if iter >= len(Input.LRInstruction)-1 {
			iter = 0
		} else {
			iter++
		}
		derpInt++
	}

	return &Input, &derpInt
}

func Part2(Input *Challenge) *int {
	lowest := -1
	return &lowest
}

type Challenge struct {
	RawData       *[]string
	LRInstruction string
	Maps          map[string]Key
}

func (c *Challenge) init() {
	c.Maps = map[string]Key{}
	for idx, line := range *c.RawData {
		if idx == 0 {
			c.LRInstruction = line
			continue
		}
		if line == "" {
			continue
		}
		key := Key{Raw: line}
		key.init()
		c.Maps[key.Value] = key
	}
}

type Key struct {
	Raw   string
	Value string
	Left  string
	Right string
}

func (k *Key) init() {
	strSplit := strings.Split(k.Raw, " = ")
	k.Value = strSplit[0]
	strSplitRight := strSplit[1]
	strSplitRight = strings.Trim(strSplitRight, "()")
	strSplitRightSplit := strings.Split(strSplitRight, ", ")
	k.Left = strSplitRightSplit[0]
	k.Right = strSplitRightSplit[1]

}
