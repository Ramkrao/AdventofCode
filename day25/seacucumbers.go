package main

import (
	"fmt"
	"strings"

	"github.com/Ramkrao/advent/utils"
)

func main() {
	lines := utils.ReadArrayFromFile("day25/input.txt")

	cucumbers := make([][]string, len(lines))
	for row, line := range lines {
		cucumbers[row] = strings.Split(line, "")
	}
	loop := 1
	is_moving := true
	for is_moving {
		cucumbers, is_moving = move(cucumbers)
		// fmt.Printf("----------Loop %d--------------\n", loop)
		// for row := range cucumbers {
		// 	fmt.Println(strings.Join(cucumbers[row], ""))
		// }
		loop++
	}
	fmt.Println("First step where no sea cucmbers moved", loop-1)
}

func move(cucumbers [][]string) ([][]string, bool) {
	is_moving := false
	// first make a copy
	copy := make([][]string, len(cucumbers))
	for row := range cucumbers {
		for range cucumbers[row] {
			copy[row] = append(copy[row], ".")
		}
	}
	// let's move east facing ones first
	for row := range cucumbers {
		for col := range cucumbers[row] {
			if cucumbers[row][col] == ">" {
				newcol := col + 1
				if newcol == len(cucumbers[row]) {
					newcol = 0
				}
				if cucumbers[row][newcol] == "." {
					copy[row][newcol] = ">"
					is_moving = true
				} else {
					copy[row][col] = ">"
				}
			}
		}
	}
	// let's move south facing ones next
	for row := range cucumbers {
		newrow := row + 1
		if newrow == len(cucumbers) {
			newrow = 0
		}
		for col := range cucumbers[row] {
			if cucumbers[row][col] == "v" {
				if (cucumbers[newrow][col] == "." || cucumbers[newrow][col] == ">") && copy[newrow][col] == "." {
					copy[newrow][col] = "v"
					is_moving = true
				} else {
					copy[row][col] = "v"
				}
			}
		}
	}
	return copy, is_moving
}
