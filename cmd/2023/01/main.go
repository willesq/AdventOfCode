package main

import (
	"adventOfCode/cmd/2023/01/internal"
	"fmt"
)

func main() {
	fpPrefix := "./cmd/2023/01"
	inputFilename := "input.txt"

	d1input := fmt.Sprintf("%s/%s", fpPrefix, inputFilename)
	// Part 1
	sum1 := internal.Part1(d1input)
	fmt.Print("2023: Day 01 - Part 1: ", *sum1, "\n")
	// Part 2
	sum2 := internal.Part2(d1input)
	fmt.Print("2023: Day 01 - Part 2: ", *sum2, "\n")

}
