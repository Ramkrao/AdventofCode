package main

import (
	"fmt"
	"strings"

	"github.com/Ramkrao/advent/utils"
)

func main() {

	line := utils.ReadArrayFromFile("day6/input.txt")

	input := utils.StrArrToIntArr(strings.Split(line[0], ","))
	fmt.Println("Read ", line)

	fishes := make([]int, 9)
	for i := range input {
		fishes[input[i]] += 1
	}

	fmt.Println("Initial state: ", fishes)
	for days := 1; days <= 256; days++ {
		fishes = spawn(fishes)

		var total int
		for i := range fishes {
			total += fishes[i]
		}
		fmt.Println("After ", days, fishes, total)
	}
}

func spawn(fishes []int) []int {
	fishes_0 := fishes[0]
	for i := range fishes {
		if i == 8 {
			fishes[i] = fishes_0
		} else {
			fishes[i] = fishes[i+1]
		}
	}
	fishes[6] += fishes_0

	return fishes
}
