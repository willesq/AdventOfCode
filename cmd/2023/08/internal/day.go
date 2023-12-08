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
	startingsNodes := map[string]Key{}
	// Find All Starting Nodes
	for idx, key := range Input.Maps {
		if strings.HasSuffix(key.Value, "A") {
			startingsNodes[idx] = key
		}
	}
	derpInt := 0
	iter := 0
	for {
		totalEndInZ := 0
		for evalKey, evalValue := range startingsNodes {
			if strings.HasSuffix(evalValue.Value, "Z") {
				totalEndInZ++
			}
			leftRight := string(Input.LRInstruction[iter])
			if leftRight == "L" {
				startingsNodes[evalKey] = Input.Maps[evalValue.Left]
			} else if leftRight == "R" {
				startingsNodes[evalKey] = Input.Maps[evalValue.Right]
			}
		}
		if totalEndInZ == len(startingsNodes) {
			break
		}
		if iter >= len(Input.LRInstruction)-1 {
			iter = 0
		} else {
			iter++
		}
		derpInt++
	}
	return &derpInt
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
