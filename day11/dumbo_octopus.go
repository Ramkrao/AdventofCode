package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Ramkrao/advent/utils"
)

var glowPoints = []string{}
var totalGlowPoints = 0

func main() {
	lines := utils.ReadArrayFromFile("day11/input.txt")
	octopuses := make([][]int, 10)
	for i := range octopuses {
		octopuses[i] = utils.StrArrToIntArr(strings.Split(lines[i], ""))
	}
	fmt.Println(octopuses)

	for count := 1; count < 999; count++ {
		//reset the glowPoints
		glowPoints = []string{}
		// increase the energy level
		increaseEneryLevels(octopuses)
		// print the results
		fmt.Println("======================")
		for i := range octopuses {
			fmt.Println(octopuses[i])
		}
		// check for all glow
		allGlowed := true
		for i := range octopuses {
			for j := range octopuses[i] {
				if octopuses[i][j] != 0 {
					allGlowed = false
				}
			}
		}
		if allGlowed {
			fmt.Println("All octopues glowed at step ", count)
			break
		}
	}
	fmt.Println("Total number of glow points ", totalGlowPoints)
}

func increaseEneryLevels(octopuses [][]int) [][]int {
	for i := range octopuses {
		for j := range octopuses[i] {
			octopuses[i][j] += 1
		}
	}
	for i := range octopuses {
		for j := range octopuses[i] {
			if octopuses[i][j] > 9 {
				octopuses = flash(octopuses, i, j)
			}
		}
	}
	return octopuses
}

// iterative function to increment flash points
func flash(octopuses [][]int, x int, y int) [][]int {
	// add the current point to glowPoints array, so we can skip incrementing it
	glowPoints = append(glowPoints, strconv.Itoa(x)+":"+strconv.Itoa(y))
	// set the starting point to 0
	octopuses[x][y] = 0
	// increment the total number of glow points
	totalGlowPoints++
	// get all adjacent points
	adjPoints := utils.ComputeAdjacentPoints(octopuses, x, y)
	// increase the step for all adjacent points
	for _, point := range adjPoints {
		pos := utils.StrArrToIntArr(strings.Split(point, ":"))
		// check if the point isn't in glowPoints, then increment
		if !utils.ContainsStr(glowPoints, point) {
			octopuses[pos[0]][pos[1]] += 1
		}
	}

	// iterate through all the adjacent point to see if any of them are 9
	// if yes, call this func recursively
	for _, point := range adjPoints {
		pos := utils.StrArrToIntArr(strings.Split(point, ":"))
		if octopuses[pos[0]][pos[1]] > 9 {
			flash(octopuses, pos[0], pos[1])
		}
	}
	return octopuses
}
