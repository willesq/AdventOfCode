package main

import (
	"adventOfCode/cmd/2022/01/internal"
	"fmt"
)

func main() {
	fpPrefix := "./cmd/2022/01"
	inputFilename := "input.txt"

	elves, sum1 := internal.Part1(fmt.Sprintf("%s/%s", fpPrefix, inputFilename))
	fmt.Print("2022: Day 01 - Part 1: ", *sum1, "\n")
	sum2 := internal.Part2(elves)
	fmt.Print("2022: Day 01 - Part 2: ", *sum2, "\n")
}
