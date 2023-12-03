package internal

import (
	"adventOfCode/internal/adventhelper"
	"fmt"
)

func Part1(filename string) (*Challenge, *int) {
	derp := adventhelper.ReadFile(fmt.Sprintf(filename))
	fmt.Sprint(derp)
	var input Challenge
	var der int
	return &input, &der
}

func Part2(input *Challenge) *int {
	derp := 0
	return &derp
}

type Challenge struct {
}
