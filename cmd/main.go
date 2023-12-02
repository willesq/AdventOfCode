package main

import (
	_1 "adventOfCode2023/cmd/01"
	_2 "adventOfCode2023/cmd/02"
	"fmt"
)

func main() {
	fpPrefix := "./cmd"
	part1Filename := "p1_input.txt"
	part2Filename := "p2_input.txt"
	day01 := false // Challenge 1
	day02 := true
	if day01 {
		// Part 1
		sum1 := _1.Part1(fmt.Sprintf("%s/01/%s", fpPrefix, part1Filename))
		fmt.Print("Day 01 - Part 1: ", *sum1, "\n")
		// Part 2
		sum2 := _1.Part2(fmt.Sprintf("%s/01/%s", fpPrefix, part2Filename))
		fmt.Print("Day 01 - Part 2: ", *sum2, "\n")
	}
	if day02 {
		input, sum1 := _2.Part1(fmt.Sprintf("%s/02/%s", fpPrefix, part1Filename))
		fmt.Print("Day 02 - Part 1: ", *sum1, "\n")

		sum2 := _2.Part2(input)
		fmt.Print("Day 02 - Part 2: ", *sum2, "\n")
	}
}
