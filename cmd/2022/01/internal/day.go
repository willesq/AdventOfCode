package internal

import (
	"adventOfCode/internal/adventhelper"
	"fmt"
	"strconv"
)

type Elf struct {
	TotalCalories int
	Calories      []int
}

type Elves struct {
	Group       []Elf
	MaxCalories int
	Top1        int
	Top2        int
	Top3        int
}

func (e *Elves) Deserialize(data *[]string) {
	var tmpElf Elf
	for _, val := range *data {
		if val == "" {
			e.Group = append(e.Group, tmpElf)
			tmpElf = Elf{}
		} else {
			valInt, _ := strconv.Atoi(val)
			tmpElf.Calories = append(tmpElf.Calories, valInt)
			tmpElf.TotalCalories += valInt
		}
	}
	for _, elf := range e.Group {
		if elf.TotalCalories > e.MaxCalories {
			e.MaxCalories = elf.TotalCalories
		}
	}
}

func Part1(filename string) (*Elves, *int) {
	derp := adventhelper.ReadFile(fmt.Sprintf(filename))
	var elves Elves
	elves.Deserialize(derp)
	return &elves, &elves.MaxCalories
}

func Part2(elves *Elves) *int {
	for _, elf := range elves.Group {
		if elf.TotalCalories > elves.Top1 {
			elves.Top3 = elves.Top2
			elves.Top2 = elves.Top1
			elves.Top1 = elf.TotalCalories
			continue
		} else if elf.TotalCalories > elves.Top2 {
			elves.Top3 = elves.Top2
			elves.Top2 = elf.TotalCalories
			continue
		} else if elf.TotalCalories > elves.Top3 {
			elves.Top3 = elf.TotalCalories
		}
	}
	top3Sum := elves.Top1 + elves.Top2 + elves.Top3
	return &top3Sum
}
