package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Ramkrao/advent/utils"
)

func main() {

	lines := utils.ReadArrayFromFile("day7/input.txt")

	crab_positions := utils.StrArrToIntArr(strings.Split(lines[0], ","))
	sort.Ints(crab_positions)
	fmt.Println("Initial positions ", crab_positions)

	results := make([]int, 0)
	for i := 0; i <= crab_positions[len(crab_positions)-1]; i++ {
		results = append(results, determineFuel(crab_positions, i))
	}
	sort.Ints(results)
	fmt.Println(results[0])
}

func determineFuel(positions []int, index int) int {

	fuel := 0
	fmt.Println("Computing position ", index)
	for i := range positions {
		if positions[i] < index {
			fuel += cumulativeAdd(index - positions[i])
		} else if positions[i] > index {
			fuel += cumulativeAdd(positions[i] - index)
		}
	}
	fmt.Println("Required fuel", fuel)
	return fuel
}

func cumulativeAdd(val int) int {
	result := 0
	for i := 0; i <= val; i++ {
		result += i
	}
	return result
}
