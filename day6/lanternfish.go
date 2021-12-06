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
		fishes[input[i]] = fishes[input[i]] + 1
	}

	fmt.Println("Initial state: ", fishes)
	for days := 1; days <= 256; days++ {
		fishes = spawn(fishes)

		var total int
		for i := range fishes {
			total = total + fishes[i]
		}
		fmt.Println("After ", days, fishes, total)
	}
}

func spawn(fishes []int) []int {
	temp := make([]int, len(fishes))

	temp[0] = fishes[1]
	temp[1] = fishes[2]
	temp[2] = fishes[3]
	temp[3] = fishes[4]
	temp[4] = fishes[5]
	temp[5] = fishes[6]
	temp[6] = fishes[7] + fishes[0]
	temp[7] = fishes[8]
	temp[8] = fishes[0]

	return temp
}
