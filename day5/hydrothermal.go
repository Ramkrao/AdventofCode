package main

import (
	"fmt"
	"strings"

	"github.com/Ramkrao/advent/utils"
)

func main() {

	lines := utils.ReadArrayFromFile("day5/input.txt")

	vents := make([][]int, 1000)
	for i := range vents {
		vents[i] = make([]int, 1000)
	}

	fmt.Println(lines)

	for _, line := range lines {
		values := parseInputLine(line)
		fmt.Println(values)
		if values[0] == values[2] || values[1] == values[3] {
			fmt.Println("processing input", values)
			vents = processHorzOrVertInput(vents, values)
		} else {
			vents = processDiagonalInput(vents, values)
		}
	}

	count := 0
	for i := range vents {
		for j := range vents[i] {
			if vents[i][j] > 1 {
				count++
			}
		}
	}
	for i := range vents {
		fmt.Println(vents[i])
	}
	fmt.Println("Total points", count)
}

func processHorzOrVertInput(vents [][]int, values []int) [][]int {
	values = swapHigherValues(values)
	for i := values[0]; i <= values[2]; i++ {
		for j := values[1]; j <= values[3]; j++ {
			fmt.Println("Updating coordinates", i, j)
			vents[i][j] = vents[i][j] + 1
		}
	}
	return vents
}

func processDiagonalInput(vents [][]int, values []int) [][]int {
	if values[0] > values[2] {
		temp := values[0]
		values[0] = values[2]
		values[2] = temp
		temp = values[1]
		values[1] = values[3]
		values[3] = temp
	}
	j := values[1]
	for i := values[0]; i <= values[2]; i++ {
		fmt.Println("Updating coordinates", i, j)
		vents[i][j] = vents[i][j] + 1
		if values[1] < values[3] {
			j++
		} else {
			j--
		}
	}
	return vents
}

func parseInputLine(line string) []int {
	return utils.StrArrToIntArr(strings.Split(strings.ReplaceAll(strings.ReplaceAll(line, "->", ","), " ", ""), ","))
}

func swapHigherValues(values []int) []int {
	if values[0] > values[2] {
		temp := values[0]
		values[0] = values[2]
		values[2] = temp
	}
	if values[1] > values[3] {
		temp := values[1]
		values[1] = values[3]
		values[3] = temp
	}
	return values
}
