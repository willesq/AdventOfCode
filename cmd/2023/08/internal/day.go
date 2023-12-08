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
	sNodes := map[string]Key{}
	// Find All Starting Nodes
	for idx, key := range Input.Maps {
		if strings.HasSuffix(key.Value, "A") {
			sNodes[idx] = key
		}
	}
	cycles := [][]int{}

	for _, val := range sNodes {
		cycle := []int{}
		tmpLFInstructions := Input.LRInstruction
		step := 0
		startKey := val.Value
		firstZ := ""
		for {
			evalKey := Input.Maps[startKey]
			for {
				if step == 0 || !strings.HasSuffix(evalKey.Value, "Z") {
					step++
					leftRight := string(tmpLFInstructions[0])
					if leftRight == "L" {
						evalKey = Input.Maps[evalKey.Left]
					} else if leftRight == "R" {
						evalKey = Input.Maps[evalKey.Right]
					}
					tmpLFInstructions = tmpLFInstructions[1:] + string(tmpLFInstructions[0])
				} else {
					break
				}
			}
			cycle = append(cycle, step)

			if firstZ == "" {
				firstZ = evalKey.Value
				step = 0
			} else if firstZ == evalKey.Value {
				break
			}
		}
		cycles = append(cycles, cycle)
	}

	derpInt := cycles[0][0]
	for _, c := range cycles[1:] {
		derpInt = (derpInt * c[0]) / GCD(derpInt, c[0])
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

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
