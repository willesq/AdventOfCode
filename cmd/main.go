package main

import (
	_1 "adventOfCode2023/cmd/01"
	_2 "adventOfCode2023/cmd/02"
	"fmt"
)

func main() {
	fpPrefix := "./cmd"
	inputFilename := "input.txt"
	day01 := false // Challenge 1
	day02 := true  // Challenge 2
	if day01 {
		d1input := fmt.Sprintf("%s/01/%s", fpPrefix, inputFilename)
		// Part 1
		sum1 := _1.Part1(d1input)
		fmt.Print("Day 01 - Part 1: ", *sum1, "\n")
		// Part 2
		sum2 := _1.Part2(d1input)
		fmt.Print("Day 01 - Part 2: ", *sum2, "\n")
	}
	if day02 {
		input, sum1 := _2.Part1(fmt.Sprintf("%s/02/%s", fpPrefix, inputFilename))
		fmt.Print("Day 02 - Part 1: ", *sum1, "\n")
		sum2 := _2.Part2(input)
		fmt.Print("Day 02 - Part 2: ", *sum2, "\n")
	}
}
